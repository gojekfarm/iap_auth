package pkey_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"testing"

	"github.com/gojektech/iap_auth/pkg/pkey"
	"github.com/stretchr/testify/assert"
)

func TestShouldHandleInvalidKey(t *testing.T) {
	parsedkey, err := pkey.Parse([]byte(""))
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf("invalid private key data"), err)
	assert.Nil(t, parsedkey)
}

func TestShouldParseRSAPKS1PKey(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("Failed to generate test data %v", err))
	}

	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	rsaPKCS1PKey := pem.EncodeToMemory(block)
	parsedkey, err := pkey.Parse(rsaPKCS1PKey)
	assert.Nil(t, err)
	assert.NotNil(t, parsedkey)
}

func TestShouldParseRSAPKCS8PKey(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("Failed to generate test data %v", err))
	}

	keyInterface, _ := x509.MarshalPKCS8PrivateKey(key)
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: keyInterface,
	}

	rsaPKCS1PKey := pem.EncodeToMemory(block)
	parsedkey, err := pkey.Parse(rsaPKCS1PKey)
	assert.Nil(t, err)
	assert.NotNil(t, parsedkey)
}

func TestShouldReturnErrorOnParseOtherKinds(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("Failed to generate test data %v", err))
	}

	keyInterface, _ := x509.MarshalPKCS8PrivateKey(key)
	block := &pem.Block{
		Type:  "OTHER KIND KEY",
		Bytes: keyInterface,
	}

	rsaPKCS1PKey := pem.EncodeToMemory(block)
	parsedkey, err := pkey.Parse(rsaPKCS1PKey)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf("invalid private key type: OTHER KIND KEY"), err)
	assert.Nil(t, parsedkey)
}
