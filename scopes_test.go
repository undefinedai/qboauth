package qboauth

import (
	"strings"
	"testing"
)

func Test_scopes(t *testing.T) {
	t.Run("it adds openid if an openid child scope is requested but openid is not", func(t *testing.T) {
		scopes := []Scope{Accounting, Email}
		got := scopesToStrings(scopes)
		hasOpenId := false
		for _, v := range got {
			if v == "openid" {
				hasOpenId = true
			}
		}
		if !hasOpenId {
			t.Errorf("expected openid to be added to scopes, got %s", strings.Join(got, ","))
		}
	})
	t.Run("it doesn't duplicate openid if an openid child scope is requested and openid is also passed", func(t *testing.T) {
		scopes := []Scope{Accounting, Email, OpenID}
		got := scopesToStrings(scopes)
		hasOpenId := false
		for _, v := range got {
			if v == "openid" {
				hasOpenId = true
			}
		}
		if !hasOpenId {
			t.Errorf("expected openid to be added to scopes, got %s", strings.Join(got, ","))
		}
		if len(got) != 3 {
			t.Errorf("expected only 3 scopes, got %d: %s", len(got), strings.Join(got, ","))
		}
	})
	t.Run("it doesn't add openid if it is not needed", func(t *testing.T) {
		scopes := []Scope{Accounting, Payments}
		got := scopesToStrings(scopes)
		hasOpenId := false
		for _, v := range got {
			if v == "openid" {
				hasOpenId = true
			}
		}
		if hasOpenId {
			t.Errorf("expected openid NOT to be added to scopes, got %s", strings.Join(got, ","))
		}
		if len(got) != 2 {
			t.Errorf("expected only 2 scopes, got %d: %s", len(got), strings.Join(got, ","))
		}
	})
}
