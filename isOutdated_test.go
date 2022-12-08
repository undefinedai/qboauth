package qboauth

import (
	"testing"
	"time"
)

func Test_isOutdated(t *testing.T) {
	t.Run("it returns true if the document is more than 24 hours old", func(t *testing.T) {
		d := Document{refreshedAt: time.Now().Add(-48 * time.Hour)}
		shouldBeRefreshed := isOutdated(&d)
		if !shouldBeRefreshed {
			t.Errorf("expected it should be refreshed, got %t", shouldBeRefreshed)
		}
	})
	t.Run("it returns false if the document is less than 24 hours old", func(t *testing.T) {
		d := Document{refreshedAt: time.Now().Add(-23 * time.Hour)}
		shouldNotBeRefreshed := isOutdated(&d)
		if shouldNotBeRefreshed {
			t.Errorf("expected it should not refresh, got %t", shouldNotBeRefreshed)
		}
	})
}
