package qboauth

import (
	"golang.org/x/oauth2"
)

func (c Client) getOAuthConfig() (*oauth2.Config, error) {
	doc, err := c.GetDiscoveryDocument()
	if err != nil {
		return nil, err
	}
	conf := &oauth2.Config{
		ClientID:     c.oAuthKey,
		ClientSecret: c.oAuthSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  doc.AuthorizationEndpoint,
			TokenURL: doc.TokenEndpoint,
		},
		RedirectURL: c.redirectURL,
		Scopes:      c.scopes,
	}

	return conf, nil
}
