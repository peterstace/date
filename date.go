package main

import (
	"fmt"
	"time"
)

// Date represents a calendar day (and thus has day precision). It's stored as
// the number of days since the epoch date 1970-01-01. It can be negative, to
// represent dates before the epoch. Because it's a date, it doesn't exist
// within any particular timezone.
type Date int

// FromTime creates a Date by truncating away the time of day portion of a
// time.Time.
func FromTime(t time.Time) Date {
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	return Date(t.Unix() / (60 * 60 * 24))
}

// String returns the ISO8601 representation (YYYY-MM-DD).
func (d Date) String() string {
	t := d.Time()
	return fmt.Sprintf("%04d-%02d-%02d", t.Year(), t.Month(), t.Day())
}

// Time returns a time.Time at midnight at the start of the date in the UTC
// timezone.
func (d Date) Time() time.Time {
	return d.TimeIn(time.UTC)
}

// TimeIn returns a time.Time at midnight at the start of the date in the
// Location specified.
func (d Date) TimeIn(loc *time.Location) time.Time {
	t := time.Date(1970, time.January, 1, 0, 0, 0, 0, loc)
	return t.AddDate(0, 0, int(d))
}

// AddMonths adds the number of specified months to create a new date. Dates are
// normalized in the same way as `AddDate` in the `time` package.
func (d Date) AddMonths(months int) Date {
	return FromTime(d.Time().AddDate(0, months, 0))
}

// AddYears adds the number of specified years to create a new date. Dates are
// normalized in the same way as `AddDate` in the `time` package.
func (d Date) AddYears(years int) Date {
	return FromTime(d.Time().AddDate(0, 0, years))
}

// AddDays adds the number of specified days to create a new date. Since Dates
// are just integers representing the number of days since 1970-01-01, the
// usual `+` operator can be used instead.
func (d Date) AddDays(days int) Date {
	return d + Date(days)
}

// StartOfMonth gives the date that is the 1st day of the current month.
func (d Date) StartOfMonth() Date {
	t := d.Time()
	return FromTime(time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC))
}

// StartOfQuarter gives the date that is the 1st day of the current quarter
// (starting in Jan, Apr, Jul, or Oct).
func (d Date) StartOfQuarter() Date {
	t := d.Time()
	m := t.Month()
	for !startOfQuarterMonth(m) {
		m--
	}
	return FromTime(time.Date(t.Year(), m, 1, 0, 0, 0, 0, time.UTC))
}

// StartOfNextQuarter gives the date that is the 1st day of the next quarter
// (starting in Jan, Apr, Jul, or Oct).
func (d Date) StartOfNextQuarter() Date {
	t := d.Time()
	m := t.Month()
	y := t.Year()
	for !startOfQuarterMonth(m) {
		if m == time.December {
			m = time.January
			y++
		} else {
			m++
		}
	}
	return FromTime(time.Date(y, m, 1, 0, 0, 0, 0, time.UTC))
}

func startOfQuarterMonth(m time.Month) bool {
	return (m-1)%3 == 0
}

// Day gives the day of the month (1-31).
func (d Date) Day() int {
	return d.Time().Day()
}

// Month gives the month the date is in.
func (d Date) Month() time.Month {
	return d.Time().Month()
}

// Year gives the year the date is in.
func (d Date) Year() int {
	return d.Time().Year()
}

// YearDay gives how many days into the year the date is (1-365).
func (d Date) YearDay() int {
	return d.Time().YearDay()
}
