package nu

import "testing"

func TestURL(t *testing.T) {
	assertEqual("https://asdf.com", URL("https", "asdf.com", "").String(), t)
	assertEqual("https://asdf.com/path", URL("https", "asdf.com", "/path").String(), t)
	assertEqual("https://asdf.com/path?a=1&b=2", URL("https", "asdf.com", "/path", "a", "1", "b", "2").String(), t)
}
