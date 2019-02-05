package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ProxyBackend struct {
	url   *url.URL
	proxy *httputil.ReverseProxy
}

func (this *ProxyBackend) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	this.proxy.ServeHTTP(rw, req)
}

func (this *ProxyBackend) URL() *url.URL {
	return this.url
}

func newProxyBackend(backendURL *url.URL, transport http.RoundTripper) *ProxyBackend {
	proxy := httputil.NewSingleHostReverseProxy(backendURL)
	proxy.Transport = transport
	return &ProxyBackend{
		url:   backendURL,
		proxy: proxy,
	}
}
