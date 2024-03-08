package util

import (
	"slices"
	"time"
)

func IsWeekend(weekday time.Weekday) bool {
	return slices.Contains([]time.Weekday{time.Saturday, time.Sunday}, weekday)
}

func IsWorkday(weekday time.Weekday) bool {
	return !IsWeekend(weekday)
}

func GetStartOfDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
