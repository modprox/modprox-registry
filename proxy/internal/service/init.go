package service

import (
	"net/http"
	"os"
	"time"

	"github.com/modprox/mp/pkg/metrics/stats"

	"github.com/modprox/mp/proxy/internal/problems"

	"github.com/pkg/errors"

	"github.com/modprox/mp/pkg/clients/payloads"
	"github.com/modprox/mp/pkg/clients/registry"
	"github.com/modprox/mp/pkg/clients/zips"
	"github.com/modprox/mp/pkg/netservice"
	"github.com/modprox/mp/pkg/upstream"
	"github.com/modprox/mp/pkg/webutil"
	"github.com/modprox/mp/proxy/internal/modules/background"
	"github.com/modprox/mp/proxy/internal/modules/store"
	"github.com/modprox/mp/proxy/internal/status/heartbeat"
	"github.com/modprox/mp/proxy/internal/status/startup"
	"github.com/modprox/mp/proxy/internal/web"
)

type initer func(*Proxy) error

func initSender(p *Proxy) error {
	cfg := p.config.Statsd.Agent
	if cfg.Port == 0 || cfg.Address == "" {
		p.emitter = stats.Discard()
		p.log.Warnf("stats emitter is set to discard client - no metrics will be reported")
		return nil
	}

	emitter, err := stats.New(stats.Proxy, p.config.Statsd)
	if err != nil {
		return err
	}
	p.emitter = emitter
	return nil
}

func initTrackers(p *Proxy) error {
	dlTracker := problems.New("downloads")
	p.dlProblems = dlTracker
	return nil
}

func initIndex(p *Proxy) error {
	var err error
	indexPath := p.config.ModuleStorage.IndexPath
	p.index, err = store.NewIndex(store.IndexOptions{
		Directory: indexPath,
	})
	return err
}

func initStore(p *Proxy) error {
	storePath := p.config.ModuleStorage.DataPath
	if storePath == "" {
		return errors.New("module_storage.path is required")
	}

	tmpPath := p.config.ModuleStorage.TmpPath
	p.store = store.NewStore(store.Options{
		Directory:    storePath,
		TmpDirectory: tmpPath,
	}, p.emitter)

	return nil
}

func initRegistryClient(p *Proxy) error {
	clientTimeout := p.config.Registry.RequestTimeoutS
	if clientTimeout <= 0 {
		return errors.Errorf(
			"registry.request_timeout_s must be > 0, got %d",
			clientTimeout,
		)
	}

	instances := p.config.Registry.Instances
	if len(instances) <= 2 {
		p.log.Warnf(
			"at least 2 registry instances recommended, got %d",
			len(instances),
		)
	}

	p.registryClient = registry.NewClient(registry.Options{
		Timeout:   time.Duration(clientTimeout) * time.Second,
		Instances: p.config.Registry.Instances,
		APIKey:    p.config.Registry.APIKey,
	})

	return nil
}

func initZipsClient(p *Proxy) error {
	httpClient := zips.NewHTTPClient(
		zips.HTTPOptions{
			Timeout: 1 * time.Minute,
		},
	)
	p.zipsClient = zips.NewClient(httpClient)
	return nil
}

func initRegistryReloadWorker(p *Proxy) error {
	reloadFreqS := time.Duration(p.config.Registry.PollFrequencyS) * time.Second
	registryRequester := background.NewRegistryAPI(p.registryClient, p.index)

	p.reloader = background.NewReloadWorker(
		background.Options{
			Frequency: reloadFreqS,
		},
		p.emitter,
		p.dlProblems,
		p.index,
		p.store,
		upstream.NewResolver(
			initTransforms(p)...,
		),
		registryRequester,
		p.zipsClient,
	)
	p.reloader.Start()
	return nil
}

func initTransforms(p *Proxy) []upstream.Transform {
	transforms := make([]upstream.Transform, 0, 1)
	transforms = append(transforms, initGoGetTransform(p))
	transforms = append(transforms, initStaticRedirectTransforms(p)...)
	transforms = append(transforms, initSetPathTransform(p))
	transforms = append(transforms, initHeaderTransforms(p)...)
	transforms = append(transforms, initTransportTransforms(p)...)
	return transforms
}

func initGoGetTransform(p *Proxy) upstream.Transform {
	if p.config.Transforms.AutomaticRedirect {
		return upstream.NewAutomaticGoGetTransform()
	}

	goGetDomains := make([]string, 0, len(p.config.Transforms.DomainGoGet))
	for _, domain := range p.config.Transforms.DomainGoGet {
		goGetDomains = append(goGetDomains, domain.Domain)
	}
	return upstream.NewGoGetTransform(goGetDomains)
}

func initStaticRedirectTransforms(p *Proxy) []upstream.Transform {
	transforms := make([]upstream.Transform, 0, len(p.config.Transforms.DomainRedirects))
	for _, domainRedirect := range p.config.Transforms.DomainRedirects {
		transforms = append(transforms, upstream.NewStaticRedirectTransform(
			domainRedirect.Original,
			domainRedirect.Substitution,
		))
	}
	return transforms
}

func initSetPathTransform(p *Proxy) upstream.Transform {
	transforms := make(map[string]upstream.Transform)
	for _, t := range p.config.Transforms.DomainPath {
		transforms[t.Domain] = upstream.NewDomainPathTransform(t.Path)
	}
	return upstream.NewSetPathTransform(transforms)
}

func initHeaderTransforms(p *Proxy) []upstream.Transform {
	transforms := make([]upstream.Transform, 0, len(p.config.Transforms.DomainHeaders))
	for _, t := range p.config.Transforms.DomainHeaders {
		transforms = append(transforms, upstream.NewDomainHeaderTransform(
			t.Domain, t.Headers,
		))
	}
	return transforms
}

func initTransportTransforms(p *Proxy) []upstream.Transform {
	transforms := make([]upstream.Transform, 0, len(p.config.Transforms.DomainTransport))
	for _, t := range p.config.Transforms.DomainTransport {
		transforms = append(transforms, upstream.NewDomainTransportTransform(
			t.Domain, t.Transport,
		))
	}
	return transforms
}

func initHeartbeatSender(p *Proxy) error {
	sender := heartbeat.NewSender(
		netservice.Instance{
			Address: netservice.Hostname(),
			Port:    p.config.APIServer.Port,
		},
		p.registryClient,
		p.emitter,
	)

	looper := heartbeat.NewLooper(
		10*time.Second,
		p.index,
		p.emitter,
		sender,
	)

	go looper.Loop()

	return nil
}

func initStartupConfigSender(p *Proxy) error {
	sender := startup.NewSender(
		p.registryClient,
		30*time.Second,
		p.emitter,
	)
	go func() {
		_ = sender.Send(
			payloads.Configuration{
				Self: netservice.Instance{
					Address: netservice.Hostname(),
					Port:    p.config.APIServer.Port,
				},
				Storage:    p.config.ModuleStorage,
				Registry:   p.config.Registry,
				Transforms: p.config.Transforms,
			},
		)
	}()
	return nil
}

func initWebServer(p *Proxy) error {
	var middles []webutil.Middleware

	mux := web.NewRouter(
		middles,
		p.index,
		p.store,
		p.emitter,
		p.dlProblems,
	)

	server, err := p.config.APIServer.Server(mux)
	if err != nil {
		return err
	}

	go func(h http.Handler) {
		var err error
		if p.config.APIServer.TLS.Enabled {
			err = server.ListenAndServeTLS(
				p.config.APIServer.TLS.Certificate,
				p.config.APIServer.TLS.Key,
			)
		} else {
			err = server.ListenAndServe()
		}

		p.log.Errorf("server stopped serving: %v", err)
		os.Exit(1)
	}(mux)

	return nil
}
