package qboauth

import (
	"io"
	"net/http"
)

func (c Client) getFromUrl() (Document, error) {
	url := urlFor(c.env)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Document{}, err
	}
	req.Header.Add("Accept", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Document{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Document{}, err
	}
	return parseJsonIntoStruct(body, c.env)
}
