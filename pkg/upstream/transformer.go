package upstream

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/shoenig/toolkit"

	"github.com/modprox/mp/pkg/coordinates"
	"github.com/modprox/mp/pkg/loggy"
)

// A Resolver is able to turn the globally unique identifier of
// a Go module (which includes a Source and a Version) and applies
// a set of Transform operations until a Request is created
// that can later be used to fetch the module from some source,
// which is typically a VCS host (e.g. github).
type Resolver interface {
	// Resolve applies any underlying Transform operations
	// and returns the resulting Request, or an error if
	// one of the Transform operations does not work.
	Resolve(coordinates.Module) (*Request, error)
}

// A Transform is one operation that is applied to a Request,
// which creates a new Request with zero or more parameters
// of the input Request having been modified. A Transform can
// be used to handle things like static domain name redirection,
// indirect domain name redirect (i.e. accommodate go-get meta URIs),
// domain-based path rewriting, etc.
//
// As time goes on, more and more Transform implementations will be
// added, to support additional use cases for enterprise environments
// which tend to have special needs.
type Transform interface {
	Modify(*Request) *Request
}

type resolver struct {
	transforms []Transform
}

// NewResolver creates a Resolver which will apply the given set
// of Transform operations in the order in which they appear.
func NewResolver(transforms ...Transform) Resolver {
	return &resolver{
		transforms: transforms,
	}
}

func (r *resolver) Resolve(mod coordinates.Module) (*Request, error) {
	request, err := NewRequest(mod)
	if err != nil {
		return nil, err
	}

	for _, transform := range r.transforms {
		request = transform.Modify(request)
	}

	return request, nil
}

// NewRequest creates a default Request from the given module. This
// initial Request is likely useless, as it only becomes useful after
// a set of Transform operations are applied to it, which then compute
// correct URI for the module it represents.
func NewRequest(mod coordinates.Module) (*Request, error) {
	domain, namespace, err := splitSource(mod.Source)
	if err != nil {
		return nil, err
	}
	return &Request{
		Transport: "https",
		Domain:    domain,
		Namespace: namespace,
		Version:   mod.Version,
	}, nil
}

func splitSource(s string) (string, Namespace, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", nil, errors.New("source is empty string")
	}

	split := strings.Split(s, "/")

	if len(split) == 1 {
		// e.g. go.opencensus.io is a whole domain used
		// to represent one package using go-get meta
		return s, nil, nil
	}

	// we have a domain and a path

	domain := split[0]
	namespace := Namespace(split[1:])
	return domain, namespace, nil
}

// A StaticRedirectTransform is used to directly convert one domain
// to another. For example, if your organization internally keeps packages
// organized like
//   ${GOPATH}/company/...
// but the internal VCS is only addressable in a way like
//   code.internal.company.net/...
// then the StaticRedirectTransform can be used to automatically acquire
// modules prefixed with name "company/" from the internal VCS of the
// different domain name.
type StaticRedirectTransform struct {
	original     string
	substitution string
	log          loggy.Logger
}

// NewStaticRedirectTransform creates a Transform which will convert
// domains of the original name to become the substitution name.
//
// Currently only exact matches on the domain are supported.
func NewStaticRedirectTransform(original, substitution string) Transform {
	return &StaticRedirectTransform{
		original:     original,
		substitution: substitution,
		log:          loggy.New("redirect-transform"),
	}
}

func (t *StaticRedirectTransform) Modify(r *Request) *Request {
	newDomain := r.Domain
	if newDomain == t.original {
		newDomain = t.substitution
	}

	modified := &Request{
		Transport: r.Transport,
		Domain:    newDomain,
		Namespace: r.Namespace,
		Version:   r.Version,
	}

	t.log.Tracef("original: %s", r)
	t.log.Tracef("modified: %s", modified)
	return modified
}

// The GoGetTransform triggers an http request to the domain
// to simply do a "?go-get=1" lookup for the real domain of where
// the module is being hosted.
//
// Additional domains can be specified via configuration.
// The known go-get redirectors in the wild include:
// - golang.org
// - google.golang.org
// - cloud.google.com
// - gopkg.in
// - contrib.go.opencensus.io
// - go.uber.org
type GoGetTransform struct {
	domains    map[string]bool // only implement redirect metadata
	httpClient *http.Client
	log        loggy.Logger
}

// NewGoGetTransform creates a GoGetTransform where any module URIs
// found in the given list of domains will be first redirected to wherever
// the go-get meta HTML tag in the domain indicates.
//
// Read more about this functionality here:
//   https://golang.org/cmd/go/#hdr-Remote_import_paths
func NewGoGetTransform(domains []string) Transform {
	match := make(map[string]bool)
	for _, domain := range domains {
		match[domain] = true
	}

	match["golang.org"] = true
	match["cloud.google.com"] = true
	match["google.golang.org"] = true
	match["gopkg.in"] = true
	match["contrib.go.opencensus.io"] = true
	match["go.opencensus.io"] = true
	match["go.uber.org"] = true

	return &GoGetTransform{
		domains: match,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		log: loggy.New("go-get-transform"),
	}
}

func (t *GoGetTransform) Modify(r *Request) *Request {
	if !t.domains[r.Domain] {
		t.log.Tracef("domain %s is not set for go-get redirects", r.Domain)
		return r
	}
	t.log.Infof("doing go-get redirect lookup for domain %s", r.Domain)

	meta, err := t.doGoGetRequest(r)
	if err != nil {
		t.log.Warnf("unable to lookup go get redirect to %s (assuming none): %v", meta, err)
		return r
	}

	t.log.Infof("redirect to: %s", meta)
	modified := &Request{
		Transport: meta.transport,
		Domain:    meta.domain,
		Namespace: strings.Split(meta.path, "/"),
		Version:   r.Version,
		// Path: set by the domain rewriter
	}

	t.log.Tracef("original: %s", r)
	t.log.Tracef("modified: %s", modified)

	return modified
}

type goGetMeta struct {
	transport string
	domain    string
	path      string
}

func (t *GoGetTransform) doGoGetRequest(r *Request) (goGetMeta, error) {
	var meta goGetMeta
	uri := fmt.Sprintf("%s://%s/%s?go-get=1", r.Transport, r.Domain, strings.Join(r.Namespace, "/"))
	request, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return meta, err
	}

	response, err := t.httpClient.Do(request)
	if err != nil {
		return meta, err
	}
	defer toolkit.Drain(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return meta, err
	}

	return parseGoGetMetadata(string(body))
}

var (
	sourceRe = regexp.MustCompile(`(http[s]?)://([\w-.]+)/([\w-./]+)`)
)

// gives us transport, domain, path
func parseGoGetMetadata(content string) (goGetMeta, error) {
	var meta goGetMeta
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.Contains(line, `name="go-source"`) {
			groups := sourceRe.FindStringSubmatch(line)
			if len(groups) != 4 {
				return meta, errors.Errorf("malformed go-source meta tag: %q", line)
			}
			return goGetMeta{
				transport: groups[1],
				domain:    groups[2],
				path:      groups[3],
			}, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return meta, err
	}
	return meta, errors.New("no go-source meta tag in response")
}

// A DomainPathTransform is used to generate or rewrite the URL path
// of the module archive that is to be fetched per the domain of desired
// module of the Request. Default path rewriting rules are provided for
// repositories ultimately hosted in github or gitlab. Additional path
// transformations should be defined for internally hosed VCSs.
//
// e.g. github:
//   https://github.com/ELEM1/ELEM2/archive/VERSION.zip
// e.g. gitlab:
//   https://gitlab.com/ELEM1/ELEM2/-/archive/VERSOIN/ELEM2-v2.0.1.zip
type DomainPathTransform struct {
	pathFmt string
}

func (t *DomainPathTransform) Modify(r *Request) *Request {
	version := addressableVersion(r.Version) // this seems a little conflated
	newPath := formatPath(t.pathFmt, version, r.Namespace)
	return &Request{
		Transport: r.Transport,
		Domain:    r.Domain,
		Namespace: r.Namespace,
		Version:   r.Version,
		Path:      newPath,
	}
}

// e.g. v2.0.0 => v2.0.0
// e.g. v0.0.0-20180111040409-fbec762f837d => fbec762f837d
// e.g. v2.3.3+incompatible => v2.3.3
func addressableVersion(version string) string {
	// dashes indicate <version>-<timestamp>-<hash> format,
	// where the hash is what is addressable in vcs
	splitOnDash := strings.Split(version, "-")
	if len(splitOnDash) == 3 {
		return splitOnDash[2] // return the hash if it exists
	}

	// plus indicates <version>+<comment> where the version
	// is what is addressable in vcs
	splitOnPlus := strings.Split(version, "+")
	if len(splitOnPlus) > 1 {
		return splitOnPlus[0]
	}

	// the version is just the version
	return version
}

// e.g. ELEM1/ELEM2/archive/VERSION.zip => shoenig/petrify/archive/v4.0.1.zip
// e.g. ELEM1/ELEM2/-/archive/VERSION/ELEM2-VERSION.zip => crypo/cryptsetup/-/archive/v2.0.1/cryptsetup-v2.0.1.zip

func formatPath(pathFmt, version string, namespace Namespace) string {
	var path = pathFmt
	for i, elem := range namespace {
		elemIdx := fmt.Sprintf("ELEM%d", i+1)
		path = strings.Replace(path, elemIdx, elem, -1)
	}
	path = strings.Replace(path, "VERSION", version, -1)
	return path
}

func NewDomainPathTransform(pathFmt string) Transform {
	return &DomainPathTransform{
		pathFmt: pathFmt,
	}
}

// DefaultPathTransforms provides a set of default Transform types which
// create the Request.Path for a known set of VCSs systems in the open source
// world (i.e. github and gitlab).
// Additional Transforms should be specified via NewSetPathTransform, which
// accepts a map of domain to Transform, for internally hosed code.
var DefaultPathTransforms = map[string]Transform{
	"github.com": NewDomainPathTransform("ELEM1/ELEM2/archive/VERSION.zip"),
	"gitlab.com": NewDomainPathTransform("ELEM1/ELEM2/-/archive/VERSION/ELEM2-VERSION.zip"),
	"":           NewDomainPathTransform(""), // unknown
}

// A SetPathTransform is a collection of transforms which set the Path of
// a Request given a domain. Think of it as a map from a domain to a
// DomainPathTransform, which can be used in the general case rather than
// specifying an explicit list of DomainPathTransform.
type SetPathTransform struct {
	domainPathTransforms map[string]Transform
	log                  loggy.Logger
}

func NewSetPathTransform(customDomainPathTransforms map[string]Transform) Transform {
	combined := combinedDomainPathTransforms(customDomainPathTransforms)
	return &SetPathTransform{
		domainPathTransforms: combined,
		log:                  loggy.New("set-path-transform"),
	}
}

func combinedDomainPathTransforms(
	customDomainPathTransforms map[string]Transform,
) map[string]Transform {
	m := make(map[string]Transform, len(DefaultPathTransforms)+len(customDomainPathTransforms))
	for domain, transform := range DefaultPathTransforms {
		m[domain] = transform
	}
	for domain, transform := range customDomainPathTransforms {
		m[domain] = transform
	}
	return m
}

func (t *SetPathTransform) Modify(r *Request) *Request {
	domainPathTransform := t.domainPathTransforms[r.Domain]
	modified := domainPathTransform.Modify(r)
	t.log.Tracef("original: %s", r)
	t.log.Tracef("modified: %s", modified)
	return modified
}

func NewDomainHeaderTransform(domain string, headers map[string]string) Transform {
	return &DomainHeaderTransform{
		domain:  domain,
		headers: headers,
		log:     loggy.New("domain-header-transform"),
	}
}

// A DomainHeaderTransform is used to set the header for a request.
// Typically one of these will be used to set the authentication key
// for https requests to an internal VCS system.
type DomainHeaderTransform struct {
	domain  string
	headers map[string]string
	log     loggy.Logger
}

func (t *DomainHeaderTransform) Modify(r *Request) *Request {
	if r.Domain != t.domain {
		return r
	}

	newHeaders := make(map[string]string, len(r.Headers))
	for k, v := range r.Headers {
		newHeaders[k] = v
	}

	for key, value := range t.headers {
		t.log.Tracef("setting a value for request header %q", key)
		newHeaders[key] = value
	}

	return &Request{
		Transport: r.Transport,
		Domain:    r.Domain,
		Namespace: r.Namespace,
		Version:   r.Version,
		Path:      r.Path,
		Headers:   newHeaders,
	}
}