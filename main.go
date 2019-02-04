package main

import (
	"fmt"

	"github.com/gojektech/iap_auth/config"
)

func main() {
	c, err := config.Load()
	fmt.Printf("Main....%v, %v\n", c.IapHost, err)
}
