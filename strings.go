package nu

import "strings"

// StringBuilder returns a *strings.Builder with start already written.
// Analogous to bytes.NewBufferString("...").
// See https://golang.org/pkg/bytes/#NewBufferString
func StringBuilder(start string) *strings.Builder {
	builder := new(strings.Builder)
	builder.WriteString(start)
	return builder
}
