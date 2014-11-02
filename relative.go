package main

import (
	"strconv"
	"time"
)

const (
	Minute       = 60
	Hour         = 60 * Minute
	Day          = 24 * Hour
	Week         = 7 * Day
	Year         = 52 * Week
	MinuteSuffix = "m"
	HourSuffix   = "h"
	DaySuffix    = "d"
	WeekSuffix   = "w"
	YearSuffix   = "y"
	Now          = "now"
)

func relative(timestamp time.Time) string {
	delta := time.Now().Unix() - timestamp.Unix()
	if delta > Year {
		return strconv.Itoa(int(delta/Year)) + YearSuffix
	}
	if delta > Week {
		return strconv.Itoa(int(delta/Week)) + WeekSuffix
	}
	if delta > Day {
		return strconv.Itoa(int(delta/Day)) + DaySuffix
	}
	if delta > Hour {
		return strconv.Itoa(int(delta/Hour)) + HourSuffix
	}
	if delta > Minute {
		return strconv.Itoa(int(delta/Minute)) + MinuteSuffix
	}
	return Now
}
