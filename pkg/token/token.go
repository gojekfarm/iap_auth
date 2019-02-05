package token

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type TokenClient struct {
	HTTPClient *http.Client
}

const (
	TokenURI      = "https://www.googleapis.com/oauth2/v4/token"
	JWTBearerType = "urn:ietf:params:oauth:grant-type:jwt-bearer"
)

func (t *TokenClient) Refresh(jwsassertion string) (string, error) {

	params := url.Values{}
	params.Set("grant_type", JWTBearerType)
	params.Set("assertion", jwsassertion)

	resp, err := t.HTTPClient.PostForm(TokenURI, params)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var tokenRes struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		IDToken     string `json:"id_token"`
		ExpiresIn   int64  `json:"expires_in"`
	}

	if err := json.Unmarshal(body, &tokenRes); err != nil {
		return "", err
	}

	return tokenRes.IDToken, nil
}
