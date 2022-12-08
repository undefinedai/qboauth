package qboauth

import (
	"regexp"
	"testing"
	"time"
)

func Test_Get(t *testing.T) {
	t.Run("it gets a new sandbox document", func(t *testing.T) {
		c := mockClient{env: Sandbox}
		d, err := get(c, c.env)
		if err != nil {
			t.Error(err)
		}
		doc = &d
		matched, err := regexp.MatchString(`sandbox.*`, doc.UserinfoEndpoint)
		if err != nil {
			t.Error(err)
		}
		if !matched {
			t.Errorf("expected a sandbox url for user info endpoint, got %s", doc.UserinfoEndpoint)
		}
		doc = nil
	})
	t.Run("it gets a new production document", func(t *testing.T) {
		c := mockClient{env: Production}
		d, err := get(c, c.env)
		if err != nil {
			t.Error(err)
		}
		doc = &d
		matched, err := regexp.MatchString(`sandbox.*`, doc.UserinfoEndpoint)
		if err != nil {
			t.Error(err)
		}
		if matched {
			t.Errorf("expected a production url for user info endpoint, got %s", doc.UserinfoEndpoint)
		}
		doc = nil
	})
	t.Run("it returns an in-memory document if one exists", func(t *testing.T) {
		c := mockClient{env: Production}
		d, err := get(c, c.env)
		if err != nil {
			t.Error(err)
		}
		doc = &d
		time.Sleep(1 * time.Second)
		newD, err := get(c, c.env)
		if err != nil {
			t.Error(err)
		}
		if newD.refreshedAt.After(doc.refreshedAt) {
			diff := newD.refreshedAt.Sub(doc.refreshedAt).Milliseconds()
			t.Errorf("expected the document from in memory but a new request was made %d ms apart", diff)
		}
		doc = nil
	})
	t.Run("it gets a new production document if the previous one was sandbox", func(t *testing.T) {
		c := mockClient{env: Sandbox}
		d, err := get(c, c.env)
		if err != nil {
			t.Error(err)
		}
		doc = &d
		if doc.environment != Sandbox {
			t.Error("expected sandbox document")
		}
		c = mockClient{env: Production}
		newD, err := get(c, c.env)
		if err != nil {
			t.Error(err)
		}
		doc = &newD
		if newD.environment != Production {
			t.Error("expected a production document")
		}
		doc = nil
	})
	t.Run("it gets a new sandbox document if the previous one was production", func(t *testing.T) {
		c := mockClient{env: Production}
		d, err := get(c, c.env)
		if err != nil {
			t.Error(err)
		}
		doc = &d
		if doc.environment != Production {
			t.Error("expected production document")
		}
		c = mockClient{env: Sandbox}
		newD, err := get(c, c.env)
		if err != nil {
			t.Error(err)
		}
		doc = &newD
		if newD.environment != Sandbox {
			t.Error("expected a sandbox document")
		}
		doc = nil
	})
}
