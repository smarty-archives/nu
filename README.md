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
Duration builds a duration using the provided values as hours, minutes, seconds,
milliseconds, microseconds, and nanoseconds. Any extra values are ignored.

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
UTCDate builds a time.Time in the UTC timezone using the provided values as
year, month, day, hour, minute, second, and nanoseconds. Any extra values are
ignored.

#### func  WaitGroup

```go
func WaitGroup(count int) *sync.WaitGroup
```
WaitGroup returns a *sync.WaitGroup with count as the initial counter value.
