package qboauth

import (
	b64 "encoding/base64"
	"fmt"
)

func basicAuthString(clientID, clientSecret string) string {
	authString := fmt.Sprintf("%s:%s", clientID, clientSecret)
	b64AuthString := b64.StdEncoding.EncodeToString([]byte(authString))
	return fmt.Sprintf("Basic %s", b64AuthString)
}
