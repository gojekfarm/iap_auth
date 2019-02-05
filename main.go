package main

import (
	"context"
	"fmt"

	"github.com/gojektech/iap_auth/config"
	"github.com/gojektech/iap_auth/pkg/iap"
	"golang.org/x/oauth2"
)

func main() {
	c, err := config.Load()
	fmt.Printf("Main....%v, %v\n", c.IapHost, err)
	hc := oauth2.NewClient(context.Background(), nil)
	iap, err := iap.New(hc, c.ServiceAccountCredentials, c.ClientID)
	if err != nil {
		fmt.Println(err)
		return
	}
	token, err := iap.Token()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(token)
}
