package nu

import "testing"

func TestURL(t *testing.T) {
	assertEqual("https://asdf.com", URL("https", "asdf.com", "").String(), t)
	assertEqual("https://asdf.com/path", URL("https", "asdf.com", "/path").String(), t)
	assertEqual("https://asdf.com/path?a=1&b=2", URL("https", "asdf.com", "/path", "a", "1", "b", "2").String(), t)
	assertEqual(URLParsed("https://asdf.com/path?a=1&b=2").String(), URL("https", "asdf.com", "/path", "a", "1", "b", "2").String(), t)
	assertPanic(func() { URL("https", "asdf.com", "/path", "a", "1", "b") }, t) // incorrect number of query string parameters
	assertPanic(func() { URLParsed("/%%28") }, t) // malformed URL
}
