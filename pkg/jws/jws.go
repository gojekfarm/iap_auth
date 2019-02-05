package jws

import (
	"crypto/rsa"
	"time"

	googlejws "golang.org/x/oauth2/jws"
)

type JWS struct {
	IssuerEmail string
	Audience    string
	PrivateKey  *rsa.PrivateKey
	ClientID    string
}

func jwsHeader() *googlejws.Header {
	return &googlejws.Header{
		Algorithm: "RS256",
		Typ:       "JWT",
	}
}

func (j *JWS) Assertion() (string, error) {
	iat := time.Now()
	exp := iat.Add(time.Hour)
	jwt := &googlejws.ClaimSet{
		Iss: j.IssuerEmail,
		Aud: j.Audience,
		Iat: iat.Unix(),
		Exp: exp.Unix(),
		PrivateClaims: map[string]interface{}{
			"target_audience": j.ClientID,
		},
	}
	return googlejws.Encode(jwsHeader(), jwt, j.PrivateKey)
}
