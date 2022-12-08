package qboauth

import (
	"testing"
)

func Test_isDocumentFor(t *testing.T) {
	t.Run("it returns true if the doc is a new pointer in Sandbox mode", func(t *testing.T) {
		doc = new(Document)
		correctEnv := doc.isFor(Sandbox)
		if !correctEnv {
			t.Errorf("expected correct environment, got %t", correctEnv)
		}
	})
	t.Run("it returns false if the doc is a new pointer in Production mode", func(t *testing.T) {
		doc = new(Document)
		correctEnv := doc.isFor(Production)
		if correctEnv {
			t.Errorf("expected wrong environment, got %t", correctEnv)
		}
	})
	t.Run("it returns false if the doc is for Sandbox but given Production mode", func(t *testing.T) {
		doc := &Document{environment: Sandbox}
		correctEnv := doc.isFor(Production)
		if correctEnv {
			t.Errorf("expected wrong environment, got %t", correctEnv)
		}
	})
	t.Run("it returns false if the doc is for Production but given Sandbox mode", func(t *testing.T) {
		doc := &Document{environment: Production}
		correctEnv := doc.isFor(Sandbox)
		if correctEnv {
			t.Errorf("expected wrong environment, got %t", correctEnv)
		}
	})
	t.Run("it returns true if the doc is for Sandbox when given Sandbox mode", func(t *testing.T) {
		doc := &Document{environment: Sandbox}
		correctEnv := doc.isFor(Sandbox)
		if !correctEnv {
			t.Errorf("expected correct environment, got %t", correctEnv)
		}
	})
	t.Run("it returns true if the doc is for Production when given Production mode", func(t *testing.T) {
		doc := &Document{environment: Production}
		correctEnv := doc.isFor(Production)
		if !correctEnv {
			t.Errorf("expected correct environment, got %t", correctEnv)
		}
	})
}
