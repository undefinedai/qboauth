package qboauth

import (
	"encoding/json"
	"time"
)

func parseJsonIntoStruct(body []byte, env Environment) (Document, error) {
	var d Document
	err := json.Unmarshal(body, &d)
	if err == nil {
		d.environment = env
		d.refreshedAt = time.Now()
	}
	return d, err
}
