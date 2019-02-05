package proxy

import (
	"fmt"
	"net/http"
	"net/url"
	"sync/atomic"
)

type Proxy interface {
	http.Handler
	Address() string
}

func New(backend string, atomictoken *atomic.Value) (Proxy, error) {
	var transport http.RoundTripper
	target, err := url.Parse(backend)
	if err != nil {
		return nil, err
	}

	transport = http.DefaultTransport

	return &proxy{
		Backend:     newProxyBackend(target, transport),
		AtomicToken: atomictoken,
	}, nil
}

type proxy struct {
	Backend     *ProxyBackend
	AtomicToken *atomic.Value
}

func (prx *proxy) Address() string {
	return prx.Backend.URL().String()
}

func (prx *proxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	token := prx.AtomicToken.Load().(string)
	fmt.Println(fmt.Sprintf("Bearer %s", token))
	req.Host = prx.Backend.URL().Host
	req.URL.Scheme = prx.Backend.URL().Scheme
	req.Header.Set("Host", prx.Address())
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	prx.Backend.ServeHTTP(rw, req)
}
