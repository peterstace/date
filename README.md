# Date

[![Build Status](https://travis-ci.org/peterstace/date.svg?branch=master)](https://travis-ci.org/peterstace/date)
[![Documentation](https://godoc.org/github.com/peterstace/date?status.svg)](http://godoc.org/github.com/peterstace/date)
[![Coverage Status](https://coveralls.io/repos/github/peterstace/date/badge.svg?branch=master)](https://coveralls.io/github/peterstace/date?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/peterstace/date)](https://goreportcard.com/report/github.com/peterstace/date)

Provides a data type `date.Date` that represents a date.

Under the hood, a `date.Date` is an integer type that represents the number of
days since the epoch date 1st Jan 1970. This means that `date.Date`s are easy
to understand and cheap to pass around.

`date.Date` leverages much of its functionality from the `time` package.

## Examples

TODO
