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

	ticker := time.NewTicker(cfg.RefreshTimeSeconds * time.Second)
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
		for t := range ticker.C {
			fmt.Println("Tick at", t)
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
