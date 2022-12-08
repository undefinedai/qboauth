package qboauth

import (
	"testing"
)

func Test_BasicAuthString(t *testing.T) {
	t.Run("it generates a basic auth string", func(t *testing.T) {
		got := basicAuthString("someClientId", "someClientSecret")
		expected := "Basic c29tZUNsaWVudElkOnNvbWVDbGllbnRTZWNyZXQ="
		if got != expected {
			t.Errorf("expected %s, got %s", expected, got)
		}
	})
}
