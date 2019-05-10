package date

import (
	"database/sql/driver"
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

// FromString creates a date from its ISO8601 (YYYY-MM-DD) representation.
func FromString(str string) (Date, error) {
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		return 0, err
	}
	return FromTime(t), nil
}

// New creates a day given a year, month, and day.
func New(year int, month time.Month, day int) Date {
	return FromTime(time.Date(year, month, day, 0, 0, 0, 0, time.UTC))
}

// MustFromString creates a date from its ISO8601 (YYYY-MM-DD) representation.
// It panics if str is not in the right format.
func MustFromString(str string) Date {
	d, err := FromString(str)
	if err != nil {
		panic(err)
	}
	return d
}

// Today gives today's date.
func Today() Date {
	return FromTime(time.Now())
}

// Yesterday gives yesterday's date.
func Yesterday() Date {
	return Today() - 1
}

// Tomorrow gives tomorrow's date.
func Tomorrow() Date {
	return Today() + 1
}

// TodayIn gives today's date in given timezone.
func TodayIn(loc *time.Location) Date {
	return FromTime(time.Now().In(loc))
}

// Yesterday gives yesterday's date in given timezone.
func YesterdayIn(loc *time.Location) Date {
	return TodayIn(loc) - 1
}

// Tomorrow gives tomorrow's date in given timezone.
func TomorrowIn(loc *time.Location) Date {
	return TodayIn(loc) + 1
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
	return FromTime(d.Time().AddDate(years, 0, 0))
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

// EndOfMonth gives the date that is the last day of the current month.
func (d Date) EndOfMonth() Date {
	return d.StartOfMonth().AddMonths(1).AddDays(-1)
}

// DaysInMonth gives the number of days in the current month
func (d Date) DaysInMonth() int {
	return d.EndOfMonth().Day()
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
	for {
		if m == time.December {
			m = time.January
			y++
		} else {
			m++
		}
		if startOfQuarterMonth(m) {
			break
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

// Weekday gives the day of the week that the date falls on.
func (d Date) Weekday() time.Weekday {
	return d.Time().Weekday()
}

// YearDay gives how many days into the year the date is (1-365).
func (d Date) YearDay() int {
	return d.Time().YearDay()
}

// MarshalJSON marshals the date into a JSON string in ISO8601 format.
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

// UnmarshalJSON unmarshals a JSON string in the ISO8601 format into a date.
func (d *Date) UnmarshalJSON(p []byte) error {
	if len(p) < 2 || p[0] != '"' || p[len(p)-1] != '"' {
		return fmt.Errorf("could not unmarshal JSON into Date: value is not a string")
	}
	var err error
	*d, err = FromString(string(p)[1 : len(p)-1])
	return err
}

// Scan implements the sql.Scanner interface, allowing the sql package to read
// sql dates into Date.
func (d *Date) Scan(src interface{}) error {
	t, ok := src.(time.Time)
	if !ok {
		return fmt.Errorf("can not scan as Date: %T", src)
	}
	*d = FromTime(t)
	return nil
}

// Value implements the driver.Valuer interface, allowing sql drivers to send
// Dates to sql databases.
func (d Date) Value() (driver.Value, error) {
	return d.String(), nil
}

// Max finds the maximum date (furthest in the direction of the future) out of
// the two given dates.
func Max(d1, d2 Date) Date {
	if d1 > d2 {
		return d1
	}
	return d2
}

// Min finds the minimum date (furthest in the direction of the past) out of
// the two given dates.
func Min(d1, d2 Date) Date {
	if d1 < d2 {
		return d1
	}
	return d2
}
