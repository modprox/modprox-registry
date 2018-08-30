package service

import (
	"github.com/modprox/libmodprox/clients/registry"
	"github.com/modprox/libmodprox/clients/zips"
	"github.com/modprox/libmodprox/loggy"
	"github.com/modprox/modprox-proxy/internal/modules/background"
	"github.com/modprox/modprox-proxy/internal/modules/store"
)

type Proxy struct {
	config         Configuration
	store          store.Store
	registryClient registry.Client
	zipsClient     zips.Client
	reloader       background.Reloader
	log            loggy.Logger
}

func NewProxy(config Configuration) *Proxy {
	p := &Proxy{
		config: config,
		log:    loggy.New("proxy-service"),
	}

	for _, f := range []initer{
		initStore,
		initRegistryClient,
		initZipsClient,
		initReloader,
		initWebserver,
	} {
		if err := f(p); err != nil {
			p.log.Errorf("failed to initialize proxy: %v", err)
			panic(err)
		}
	}

	return p
}

func (p *Proxy) Run() {
	select {
	// intentionally left blank
	}
}
