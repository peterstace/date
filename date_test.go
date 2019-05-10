package date

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	Sydney, err := time.LoadLocation("Australia/Sydney")
	if err != nil {
		panic(err)
	}

	for i, test := range []struct {
		want string
		got  interface{}
	}{
		{
			"1970-01-01",
			MustFromString("1970-01-01"),
		},
		{
			"2018-03-30",
			FromTime(time.Date(2018, time.March, 30, 12, 01, 45, 0, time.UTC)),
		},
		{
			"2005-04-21",
			FromTime(MustFromString("2005-04-21").Time()),
		},
		{
			"2001-04-11 00:00:00 +1000 AEST",
			MustFromString("2001-04-11").TimeIn(Sydney),
		},
		{
			"2000-02-25",
			MustFromString("1999-12-25").AddMonths(2),
		},
		{
			"1999-10-25",
			MustFromString("1999-12-25").AddMonths(-2),
		},
		{
			"1999-07-01", // May only has 31 days, so normalises to 1st of July.
			MustFromString("1999-05-31").AddMonths(1),
		},
		{
			"2000-12-25",
			MustFromString("1999-12-25").AddYears(1),
		},
		{
			"2017-03-01", // Feb 2016 has 31 days, but Feb 2017 has 28 days, so normalises.
			MustFromString("2016-02-29").AddYears(1),
		},
		{
			"1999-12-25",
			MustFromString("1999-12-26").AddDays(-1),
		},
		{
			"1999-12-25",
			MustFromString("1999-12-24").AddDays(1),
		},
		{
			"2022-11-01",
			MustFromString("2022-11-15").StartOfMonth(),
		},
		{
			"2022-12-01",
			MustFromString("2022-12-15").StartOfMonth(),
		},
		{
			"2022-01-01",
			MustFromString("2022-01-15").StartOfMonth(),
		},
		{
			"2022-11-30",
			MustFromString("2022-11-01").EndOfMonth(),
		},
		{
			"2022-12-31",
			MustFromString("2022-12-15").EndOfMonth(),
		},
		{
			"2022-01-31",
			MustFromString("2022-01-31").EndOfMonth(),
		},
		{
			"30",
			MustFromString("2022-11-01").DaysInMonth(),
		},
		{
			"31",
			MustFromString("2022-12-15").DaysInMonth(),
		},
		{
			"31",
			MustFromString("2022-01-31").DaysInMonth(),
		},
		{
			"2022-01-01",
			MustFromString("2022-01-15").StartOfQuarter(),
		},
		{
			"2022-01-01",
			MustFromString("2022-02-15").StartOfQuarter(),
		},
		{
			"2023-01-01",
			MustFromString("2022-12-15").StartOfNextQuarter(),
		},
		{
			"2022-07-01",
			MustFromString("2022-06-15").StartOfNextQuarter(),
		},
		{
			"2022-07-01",
			MustFromString("2022-04-22").StartOfNextQuarter(),
		},

		{"4", MustFromString("2015-03-04").Day()},
		{"March", MustFromString("2015-03-04").Month()},
		{"2015", MustFromString("2015-03-04").Year()},
		{"63", MustFromString("2015-03-04").YearDay()},

		{
			time.Now().Format("2006-01-02"),
			Today(),
		},
		{
			time.Now().Add(-24 * time.Hour).Format("2006-01-02"),
			Yesterday(),
		},
		{
			time.Now().Add(24 * time.Hour).Format("2006-01-02"),
			Tomorrow(),
		},

		{
			time.Now().In(Sydney).Format("2006-01-02"),
			TodayIn(Sydney),
		},
		{
			time.Now().In(Sydney).Add(-24 * time.Hour).Format("2006-01-02"),
			YesterdayIn(Sydney),
		},
		{
			time.Now().In(Sydney).Add(24 * time.Hour).Format("2006-01-02"),
			TomorrowIn(Sydney),
		},

		{"Wednesday", MustFromString("1989-06-14").Weekday()},
		{"Thursday", MustFromString("2014-12-25").Weekday()},
		{"Saturday", MustFromString("2018-08-18").Weekday()},

		{
			"2018-05-26",
			New(2018, time.May, 26),
		},

		{
			"2018-05-05",
			Max(MustFromString("2018-05-05"), MustFromString("2018-01-01")),
		},
		{
			"2018-05-05",
			Max(MustFromString("2018-01-01"), MustFromString("2018-05-05")),
		},
		{
			"2018-01-01",
			Min(MustFromString("2018-05-05"), MustFromString("2018-01-01")),
		},
		{
			"2018-01-01",
			Min(MustFromString("2018-01-01"), MustFromString("2018-05-05")),
		},
	} {
		if gotStr := fmt.Sprintf("%v", test.got); gotStr != test.want {
			t.Errorf("i=%d got=%v want=%v", i, gotStr, test.want)
		}
	}
}

func TestFromStringErr(t *testing.T) {
	if _, err := FromString("not a date"); err == nil {
		t.Error("expected error")
	}
}

func TestJSON(t *testing.T) {
	type J struct {
		D Date `json:"d"`
	}
	j := J{MustFromString("2015-05-21")}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(j)
	if err != nil {
		t.Fatalf("Could not encode: %v", err)
	}
	if want := `{"d":"2015-05-21"}`; strings.TrimSpace(buf.String()) != want {
		t.Fatalf("Did not encode correctly, want=%q got=%q",
			want, strings.TrimSpace(buf.String()))
	}

	j = J{}
	err = json.NewDecoder(&buf).Decode(&j)
	if err != nil {
		t.Fatalf("Could not decode: %v", err)
	}
	if want := MustFromString("2015-05-21"); j.D != want {
		t.Fatalf("Did not decode correctly, want=%q got=%q", want, j.D)
	}
}

func TestSQLScan(t *testing.T) {
	var d Date
	if err := d.Scan(time.Date(2013, time.July, 13, 0, 0, 0, 0, time.UTC)); err != nil {
		t.Fatal(err)
	}
	if d != MustFromString("2013-07-13") {
		t.Fatalf("Got=%v", d)
	}
}

func TestSQLValue(t *testing.T) {
	v, err := MustFromString("2013-07-13").Value()
	if err != nil {
		t.Fatal(err)
	}
	if v != "2013-07-13" {
		t.Fatalf("Got=%v", v)
	}
}
