package pkey

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

func Parse(bytes []byte) (*rsa.PrivateKey, error) {
	var key *rsa.PrivateKey
	var err error
	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, fmt.Errorf("invalid private key data")
	}

	if block.Type == "RSA PRIVATE KEY" {
		key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
	} else if block.Type == "PRIVATE KEY" {
		keyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		var ok bool
		key, ok = keyInterface.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("not RSA private key")
		}
	} else {
		return nil, fmt.Errorf("invalid private key type: %s", block.Type)
	}

	key.Precompute()

	if err := key.Validate(); err != nil {
		return nil, err
	}
	return key, nil
}
