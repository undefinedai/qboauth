package qboauth

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c Client) Revoke(token string) error {
	doc, err := c.GetDiscoveryDocument()
	if err != nil {
		return err
	}

	postBody := struct {
		Token string `json:"token"`
	}{Token: token}

	data, err := json.Marshal(postBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", doc.RevocationEndpoint, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", basicAuthString(c.oAuthKey, c.oAuthSecret))
	req.Header.Add("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)

	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return errors.New(string(b))
	}
	return nil
}
