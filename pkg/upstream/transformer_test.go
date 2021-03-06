package upstream

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"gophers.dev/pkgs/loggy"

	"oss.indeed.com/go/modprox/pkg/coordinates"
)

func ns(path string) Namespace {
	return strings.Split(path, "/")
}

func Test_Resolver_UseProxy_static_no(t *testing.T) {
	resolver := NewResolver(
		NewStaticRedirectTransform("a.com", "b.com"),
	)

	useProxy, err := resolver.UseProxy(coordinates.Module{
		Source:  "a.com", // static redirect to b.com
		Version: "v1.1.1",
	})
	require.NoError(t, err)
	require.False(t, useProxy)
}

func Test_Resolver_UseProxy_static_yes(t *testing.T) {
	resolver := NewResolver(
		NewStaticRedirectTransform("a.com", "b.com"),
	)

	useProxy, err := resolver.UseProxy(coordinates.Module{
		Source:  "z.com", // no static redirect
		Version: "v1.1.1",
	})
	require.NoError(t, err)
	require.True(t, useProxy)
}

func Test_Resolver_UseProxy_domain_header_no(t *testing.T) {
	resolver := NewResolver(
		NewDomainHeaderTransform("f.com", map[string]string{
			"X-My-Header": "abc123",
		}),
	)

	useProxy, err := resolver.UseProxy(coordinates.Module{
		Source:  "b.com", // no headers set
		Version: "v1.1.1",
	})
	require.NoError(t, err)
	require.True(t, useProxy)
}

func Test_Resolver_UseProxy_domain_transport_yes(t *testing.T) {
	resolver := NewResolver(
		NewDomainTransportTransform("c.com", "http"),
	)

	useProxy, err := resolver.UseProxy(coordinates.Module{
		Source:  "d.com", // no transport set
		Version: "v1.1.1.",
	})
	require.NoError(t, err)
	require.True(t, useProxy)
}

func Test_Resolver_UseProxy_domain_transport_no(t *testing.T) {
	resolver := NewResolver(
		NewDomainTransportTransform("c.com", "http"),
	)

	useProxy, err := resolver.UseProxy(coordinates.Module{
		Source:  "c.com", // custom transport set
		Version: "v1.1.1.",
	})
	require.NoError(t, err)
	require.False(t, useProxy)
}

func Test_Resolver_UseProxy_domain_header_yes(t *testing.T) {
	resolver := NewResolver(
		NewDomainHeaderTransform("f.com", map[string]string{
			"X-My-Header": "abc123",
		}),
	)

	useProxy, err := resolver.UseProxy(coordinates.Module{
		Source:  "f.com", // has headers set
		Version: "v1.1.1",
	})
	require.NoError(t, err)
	require.False(t, useProxy)
}

func Test_NewAutomaticGoGetTransform200(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadFile("../../hack/html/fuse.html")
		require.NoError(t, err)

		w.WriteHeader(200)
		_, err = w.Write(data)
		require.NoError(t, err)
	}))
	defer ts.Close()

	transform := &GoGetTransform{
		httpClient: ts.Client(),
		log:        loggy.New("log"),
	}

	uri, err := url.ParseRequestURI(ts.URL)
	require.NoError(t, err)

	request := &Request{
		Transport: uri.Scheme,
		Domain:    uri.Host,
		Namespace: ns("fuse"),
		Version:   "latest",
	}

	newRequest, err := transform.Modify(request)
	require.NoError(t, err)
	require.Equal(t, &Request{
		Transport: "https",
		Domain:    "github.com",
		Namespace: ns("bazil/bazil"),
		Version:   "latest",
	}, newRequest)
}

func Test_NewAutomaticGoGetTransform404(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		_, err := w.Write([]byte{})
		require.NoError(t, err)
	}))
	defer ts.Close()

	transform := &GoGetTransform{
		httpClient: ts.Client(),
		log:        loggy.New("log"),
	}

	uri, err := url.ParseRequestURI(ts.URL)
	require.NoError(t, err)

	request := &Request{
		Transport: uri.Scheme,
		Domain:    uri.Host,
		Namespace: ns("fuse"),
		Version:   "latest",
	}

	newRequest, err := transform.Modify(request)
	require.NoError(t, err)
	require.Equal(t, &Request{
		Transport: uri.Scheme,
		Domain:    uri.Host,
		Namespace: ns("fuse"),
		Version:   "latest",
	}, newRequest)
}

func Test_NewRequest(t *testing.T) {
	mod := coordinates.Module{
		Source:  "github.com/example/toolkit",
		Version: "v1.0.1",
	}

	request, err := NewRequest(mod)
	require.NoError(t, err)
	require.Equal(t, &Request{
		Transport: "https",
		Domain:    "github.com",
		Namespace: ns("example/toolkit"),
		Version:   "v1.0.1",
	}, request)
}

func Test_NewRequest_no_path_is_ok(t *testing.T) {
	// An example pulled from real life, go.opencensus.io is
	// itself pointed at a repository using go-get meta.
	// It has no path.
	mod := coordinates.Module{
		Source:  "go.opencensus.io",
		Version: "v0.15.0",
	}

	request, err := NewRequest(mod)
	require.NoError(t, err)
	require.Equal(t, &Request{
		Transport: "https",
		Domain:    "go.opencensus.io",
		Namespace: nil,
		Version:   "v0.15.0",
	}, request)
}

func Test_StaticRedirectTransform(t *testing.T) {
	request := &Request{
		Transport: "https",
		Domain:    "mycompany",
		Namespace: ns("a/b/c"),
		Version:   "v1.0.1",
	}

	rt := NewStaticRedirectTransform("mycompany", "code.mycompany.net")

	transformed, err := rt.Modify(request)
	require.NoError(t, err)
	require.Equal(t, &Request{
		Transport: "https",
		Domain:    "code.mycompany.net",
		Namespace: ns("a/b/c"),
		Version:   "v1.0.1",
	}, transformed)
}

func Test_formatPath(t *testing.T) {
	try := func(pathFmt string, ns Namespace, version, exp string) {
		result := formatPath(pathFmt, version, ns)
		require.Equal(t, exp, result)
	}

	// github
	try(
		"ELEM1/ELEM2/archive/VERSION.zip",
		ns("example/toolkit"),
		"v1.0.1",
		"example/toolkit/archive/v1.0.1.zip",
	)

	// gitlab
	try(
		"ELEM1/ELEM2/-/archive/VERSION/ELEM2-VERSION.zip",
		ns("crypo/cryptsetup"),
		"v2.0.1",
		"crypo/cryptsetup/-/archive/v2.0.1/cryptsetup-v2.0.1.zip",
	)
}

func Test_addressableVersion(t *testing.T) {
	try := func(input string, exp string) {
		output := addressableVersion(input)
		require.Equal(t, exp, output)
	}

	try("v2.0.0", "v2.0.0")
	try("v0.0.0-20180111040409-fbec762f837d", "fbec762f837d")
	try("v2.3.3+incompatible", "v2.3.3")
}

func Test_DomainTransportTransform(t *testing.T) {
	request := &Request{
		Transport: "https",
		Domain:    "foo.com",
		Namespace: ns("a/b"),
		Version:   "v1.0.0",
	}

	dtt := NewDomainTransportTransform("foo.com", "http")

	transformed, err := dtt.Modify(request)
	require.NoError(t, err)
	require.Equal(t, &Request{
		Transport: "http", // changed
		Domain:    "foo.com",
		Namespace: ns("a/b"),
		Version:   "v1.0.0",
	}, transformed)
}
