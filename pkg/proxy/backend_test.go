package proxy

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReverseProxyToGivenUrl(t *testing.T) {
	expectedresponse := []byte("backend response")
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(expectedresponse)
		w.WriteHeader(200)
	}))
	defer backend.Close()

	backendurl, _ := url.Parse(backend.URL)
	frontend := httptest.NewServer(newProxyBackend(backendurl, http.DefaultTransport))
	defer frontend.Close()
	frontendClient := frontend.Client()

	getReq, _ := http.NewRequest("GET", frontend.URL, nil)
	getReq.Host = "some-name"
	getReq.Header.Set("Connection", "close")

	getReq.Close = true

	res, err := frontendClient.Do(getReq)
	assert.Nil(t, err)
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, expectedresponse, bodyBytes)
}
