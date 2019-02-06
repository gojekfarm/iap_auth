package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gojektech/iap_auth/config"
	"github.com/gojektech/iap_auth/pkg/iap"
	"github.com/gojektech/iap_auth/pkg/proxy"
	"golang.org/x/oauth2"
)

func main() {
	cfg, err := config.Load()
	fmt.Println(cfg.RefreshTimeSeconds)
	tickPeriod, err := time.ParseDuration(cfg.RefreshTimeSeconds)
	if err != nil {
		fmt.Println(err)
		return
	}
	ticker := time.NewTicker(tickPeriod)
	var mu sync.Mutex

	var atomictoken atomic.Value

	var tokenfn = func() string {
		hc := oauth2.NewClient(context.Background(), nil)
		iap, err := iap.New(hc, cfg.ServiceAccountCredentials, cfg.ClientID)
		if err != nil {
			return "INVALID"
		}
		token, err := iap.Token()
		if err != nil {
			return "INVALID"
		}
		return token
	}
	atomictoken.Store(tokenfn())
	go func() {
		for range ticker.C {
			mu.Lock()
			atomictoken.Store(tokenfn())
			mu.Unlock()
		}
	}()

	p, err := proxy.New(cfg.IapHost, &atomictoken)
	if err != nil {
		fmt.Println(err)
		return
	}
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: p,
	}

	server.ListenAndServe()
}
