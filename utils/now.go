package utils

import "time"

func Now() *time.Time {
	t := time.Now().Add(-time.Hour * 3)
	return &t
}
