package qboauth

import "time"

func isOutdated(d *Document) bool {
	oneDayAgo := time.Now().Add(-24 * time.Hour)
	return d.refreshedAt.Before(oneDayAgo)
}
