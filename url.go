package nu

import "net/url"

// URLParsed parses the value as a url and panics in the case of an error.
func URLParsed(value string) *url.URL {
	parsed, err := url.Parse(value)
	if err != nil {
		panic(err)
	}
	return parsed
}

// URL builds a *url.URL using the provided scheme, host, path and queryPairs. Examples:
// - URL("https", "example.com", "cats", "a", "1", "b", "2") would return a *url.URL
// equivalent to "https://example.com/cats?a=1&b=2"
func URL(scheme, host, path string, queryPairs ...string) *url.URL {
	return &url.URL{
		Scheme:   scheme,
		Host:     host,
		Path:     path,
		RawQuery: URLValues(queryPairs...).Encode(),
	}
}

// URLValues populates and returns a url.Values using the provided string key/value pairs.
// Much like strings.NewReplacer, it panics if len(pairs) is not even.
// See https://golang.org/pkg/strings/#NewReplacer
func URLValues(keyValuePairs ...string) url.Values {
	if len(keyValuePairs)%2 != 0 {
		panic("Must provide an even number of key/value strings.")
	}
	values := make(url.Values, len(keyValuePairs)/2)
	for x := 0; x < len(keyValuePairs); x += 2 {
		key := keyValuePairs[x]
		value := keyValuePairs[x+1]
		values.Set(key, value)
	}
	return values
}
