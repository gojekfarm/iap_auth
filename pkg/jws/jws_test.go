package jws_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"strings"
	"testing"

	"github.com/gojektech/iap_auth/pkg/jws"
	"github.com/gojektech/iap_auth/pkg/pkey"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateAndAssertionValidFor1Hour(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 2048)
	rsakey, _ := pkey.Parse(pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}))

	j := jws.JWS{IssuerEmail: "foo@bar.com", Audience: "some_audience", PrivateKey: rsakey}
	assertion, err := j.Assertion()

	assert.Nil(t, err)
	assert.NotNil(t, assertion)

	var jwsprops map[string]interface{}
	rawjson, _ := base64.StdEncoding.DecodeString(fmt.Sprintf("%s=", strings.Split(assertion, ".")[1]))
	json.Unmarshal(rawjson, &jwsprops)
	assert.Equal(t, 3600, int(jwsprops["exp"].(float64)-jwsprops["iat"].(float64)))
}

func TestShouldAddAudienceAndTargetAudienceAndIssuer(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 2048)
	rsakey, _ := pkey.Parse(pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}))

	j := jws.JWS{IssuerEmail: "foo@bar.com", Audience: "some_audience", PrivateKey: rsakey, ClientID: "some-client"}
	assertion, _ := j.Assertion()

	var jwsprops map[string]interface{}
	rawjson, _ := base64.StdEncoding.DecodeString(fmt.Sprintf("%s==", strings.Split(assertion, ".")[1]))
	json.Unmarshal(rawjson, &jwsprops)
	assert.Equal(t, "some_audience", jwsprops["aud"])
	assert.Equal(t, "foo@bar.com", jwsprops["iss"])
	assert.Equal(t, "some-client", jwsprops["target_audience"])
}
