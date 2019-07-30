[![Build Status](https://travis-ci.org/smartystreets/nu.svg?branch=master)](https://travis-ci.org/smartystreets/nu)
[![codecov](https://codecov.io/gh/smartystreets/nu/branch/master/graph/badge.svg)](https://codecov.io/gh/smartystreets/nu)

# nu
--
    import "github.com/smartystreets/nu"

package nu implements convenient constructor functions for commonly used
standard library types.

## Usage

#### func  Duration

```go
func Duration(values ...int) time.Duration
```
Duration builds a duration using up to the first 6 provided values as hours,
minutes, seconds, milliseconds, microseconds, and nanoseconds. Any extra values
are ignored. Examples: - Duration(1, 2, 3) would return the duration equal to
"1h2m3s". - Duration(1, 2, 3, 4, 5, 6) would return the duration equal to
"1h2m3.004005006s".

#### func  Error

```go
func Error(args ...interface{}) error
```
Error is a convenient combination of errors.New and fmt.Sprint.

#### func  Errorf

```go
func Errorf(format string, args ...interface{}) error
```
Errorf is a convenient combination of errors.New and fmt.Sprintf.

#### func  Errorln

```go
func Errorln(args ...interface{}) error
```
Errorln is a convenient combination of errors.New and fmt.Sprintln.

#### func  StringBuilder

```go
func StringBuilder(start string) *strings.Builder
```
StringBuilder returns a *strings.Builder with start already written. Analogous
to bytes.NewBufferString("..."). See
https://golang.org/pkg/bytes/#NewBufferString

#### func  URL

```go
func URL(scheme, host, path string, queryPairs ...string) *url.URL
```
URL builds a *url.URL using the provided scheme, host, path and queryPairs.
Examples: - URL("https", "example.com", "cats", "a", "1", "b", "2") would return
a *url.URL equivalent to "https://example.com/cats?a=1&b=2"

#### func  URLParsed

```go
func URLParsed(value string) *url.URL
```
URLParsed parses the value as a url and panics in the case of an error.

#### func  URLValues

```go
func URLValues(keyValuePairs ...string) url.Values
```
URLValues populates and returns a url.Values using the provided string key/value
pairs. Much like strings.NewReplacer, it panics if len(pairs) is not even. See
https://golang.org/pkg/strings/#NewReplacer

#### func  UTCDate

```go
func UTCDate(values ...int) time.Time
```
UTCDate builds a time.Time in the UTC timezone using up to the first 6 provided
values as year, month, day, hour, minute, second, and nanoseconds. Any extra
values are ignored. Examples: - UTCDate(2000, 1, 2) would return a time.Time
equal to "2000-01-02 00:00:00 +0000 UTC" - UTCDate(2000, 1, 2, 15, 16, 17, 18)
would return a time.Time equal to "2000-01-02 15:16:17.000000018 +0000 UTC"

#### func  WaitGroup

```go
func WaitGroup(count int) *sync.WaitGroup
```
WaitGroup returns a *sync.WaitGroup with count as the initial counter value.
