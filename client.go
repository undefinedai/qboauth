package qboauth

import (
	"net/http"
	"time"
)

type Client struct {
	env         Environment
	httpClient  *http.Client
	oAuthKey    string
	oAuthSecret string
	redirectURL string
	scopes      []string
}

func NewClient(clientID, clientSecret, redirectURL string, scopes []Scope, env Environment) Client {
	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    1,
			IdleConnTimeout: 30 * time.Second,
		},
	}
	return Client{
		env:         env,
		httpClient:  httpClient,
		oAuthKey:    clientID,
		oAuthSecret: clientSecret,
		redirectURL: redirectURL,
		scopes:      scopesToStrings(scopes),
	}
}
