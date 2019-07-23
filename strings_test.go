package nu

import "testing"

func TestStringBuilder(t *testing.T) {
	assertEqual("", StringBuilder("").String(), t)
	assertEqual("Hello, World!", StringBuilder("Hello, World!").String(), t)
}
