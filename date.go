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

func NewDate(t time.Time) Date {
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	return Date(t.Unix() / (60 * 60 * 24))
}

func (d Date) String() string {
	t := d.Time()
	return fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
}

func (d Date) Time() time.Time {
	t := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	return t.AddDate(0, 0, int(d))
}

func (d Date) AddMonths(months int) Date {
	return NewDate(d.Time().AddDate(0, months, 0))
}

func (d Date) AddYear(years int) Date {
	return NewDate(d.Time().AddDate(0, 0, years))
}

func (d Date) StartOfMonth() Date {
	t := d.Time()
	return NewDate(time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC))
}

func (d Date) StartOfQuarter() Date {
	t := d.Time()
	m := t.Month()
	for !startOfQuarterMonth(m) {
		m--
	}
	return NewDate(time.Date(t.Year(), m, 1, 0, 0, 0, 0, time.UTC))
}

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
	return NewDate(time.Date(y, m, 1, 0, 0, 0, 0, time.UTC))
}

func startOfQuarterMonth(m time.Month) bool {
	switch m {
	case time.January, time.April, time.July, time.October:
		return true
	default:
		return false
	}
}

func (d Date) Day() int {
	return d.Time().Day()
}

func (d Date) Month() time.Month {
	return d.Time().Month()
}

func (d Date) Year() int {
	return d.Time().Year()
}

func (d Date) YearDay() int {
	return d.Time().YearDay()
}
