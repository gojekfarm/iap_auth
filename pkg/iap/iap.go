package iap

import (
	"io/ioutil"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gojektech/iap_auth/pkg/jws"
	"github.com/gojektech/iap_auth/pkg/pkey"
	"github.com/gojektech/iap_auth/pkg/token"
	"golang.org/x/oauth2/google"
)

// IAP represents the information needed to access IAP-protected app
type IAP struct {
	ServiceAccount string       `valid:"required"`
	ClientID       string       `valid:"dns,required"`
	JWS            jws.JWS      `valid:"-"`
	HTTPClient     *http.Client `valid:"-"`
}

func New(hc *http.Client, sa, id string) (*IAP, error) {
	newiap := &IAP{
		ServiceAccount: sa,
		ClientID:       id,
	}
	result, err := govalidator.ValidateStruct(newiap)
	if result == false || err != nil {
		return &IAP{}, err
	}
	serviceaccount, err := ioutil.ReadFile(newiap.ServiceAccount)
	if err != nil {
		return nil, err
	}
	conf, err := google.JWTConfigFromJSON(serviceaccount)
	if err != nil {
		return nil, err
	}
	rsaKey, err := pkey.Parse(conf.PrivateKey)
	if err != nil {
		return nil, err
	}

	newiap.JWS = jws.JWS{IssuerEmail: conf.Email, Audience: token.TokenURI, PrivateKey: rsaKey, ClientID: id}
	newiap.HTTPClient = hc
	return newiap, nil
}

func (c *IAP) Token() (string, error) {
	assertionMsg, err := c.JWS.Assertion()
	if err != nil {
		return "", err
	}

	tokenClient := token.TokenClient{HTTPClient: c.HTTPClient}
	return tokenClient.Refresh(assertionMsg)
}
