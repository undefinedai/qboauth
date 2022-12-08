package qboauth

import (
	"regexp"
	"testing"
)

func Test_urlFor(t *testing.T) {
	t.Run("it gets production", func(t *testing.T) {
		got := urlFor(Production)
		matched, err := regexp.Match(`sandbox.*`, []byte(got))
		if err != nil {
			t.Error(err)
		}
		if matched {
			t.Errorf("expected Production url, got %s", got)
		}
	})
	t.Run("it gets sandbox", func(t *testing.T) {
		got := urlFor(Sandbox)
		matched, err := regexp.MatchString(`sandbox.*`, got)
		if err != nil {
			t.Error(err)
		}
		if !matched {
			t.Errorf("expected Sandbox url, got %s", got)
		}
	})
	t.Run("it defaults to sandbox if given a bad environment", func(t *testing.T) {
		got := urlFor(99999999)
		matched, err := regexp.MatchString(`sandbox.*`, got)
		if err != nil {
			t.Error(err)
		}
		if !matched {
			t.Errorf("expected Sandbox url, got %s", got)
		}
	})
}
