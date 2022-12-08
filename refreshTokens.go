package qboauth

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (c Client) Refresh(refreshToken string) (Tokens, error) {
	conf, err := c.getOAuthConfig()
	if err != nil {
		return Tokens{}, err
	}

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)
	encodedData := data.Encode()

	req, err := http.NewRequest("POST", conf.Endpoint.TokenURL, strings.NewReader(encodedData))
	if err != nil {
		return Tokens{}, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", basicAuthString(c.oAuthKey, c.oAuthSecret))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.httpClient.Do(req)

	if err != nil {
		return Tokens{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Tokens{}, err
	}

	var tokens Tokens
	err = json.Unmarshal(body, &tokens)
	return tokens, err
}
