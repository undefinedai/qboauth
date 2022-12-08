package qboauth

import (
	"golang.org/x/oauth2"
)

func (c Client) GetAuthCodeURL(state string) (string, error) {
	conf, err := c.getOAuthConfig()
	if err != nil {
		return "", err
	}
	url := conf.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return url, nil
}
