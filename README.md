# date

[![Build Status](https://travis-ci.org/peterstace/date.svg?branch=master)](https://travis-ci.org/peterstace/date)
[![Documentation](https://godoc.org/github.com/peterstace/date?status.svg)](http://godoc.org/github.com/peterstace/date)
[![Coverage Status](https://coveralls.io/repos/github/peterstace/date/badge.svg?branch=master)](https://coveralls.io/github/peterstace/date?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/peterstace/date)](https://goreportcard.com/report/github.com/peterstace/date)

Provides a data type `Date` that represents a day-precision point in time.

`Date` is conceptually similar to Unix time (also known as POSIX time or Unix
Epoch time). But rather than operating at second-level precision, it operates
at day-level precision. Under the hood, `Date` is an integer type that
represents the number of days elapsed since the epoch date 1970-01-01.

`date.Date` leverages much of its functionality from the `time` package.

## Features

- Since `Date` is an integer type, the normal integer comparison and
  equality/inequality operators all work out of the box.

- Conversions to and from ISO8601 (YYYY-MM-DD) string format.

- Conversion to and from `time.Time` values.

- Day, month, and year arithmetic (can add or subtract years, months, days from
  a `Date`).

- Common date related operations, such as finding the start of a month, start
  of a quarter etc.

- JSON marshal/unmarshal support.

- SQL marshal/unmarshal support.
