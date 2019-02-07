package proxy_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/gojektech/iap_auth/pkg/logger"
	"github.com/gojektech/iap_auth/pkg/proxy"
	"github.com/stretchr/testify/assert"
)

func TestShouldReverseProxyToGivenUrlWithAuthorizationHeaders(t *testing.T) {
	logger.SetupLogger("debug")
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.Header.Get("Authorization")))
		w.WriteHeader(200)
	}))
	defer backend.Close()

	var atomictoken atomic.Value
	atomictoken.Store("blahblah")
	p, _ := proxy.New(backend.URL, &atomictoken)
	frontend := httptest.NewServer(p)
	defer frontend.Close()
	frontendClient := frontend.Client()

	getReq, _ := http.NewRequest("GET", frontend.URL, nil)
	getReq.Host = "some-name"
	getReq.Header.Set("Connection", "close")

	getReq.Close = true

	res, err := frontendClient.Do(getReq)
	assert.Nil(t, err)
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, []byte("Bearer blahblah"), bodyBytes)
}
