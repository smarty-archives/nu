package nu

import (
	"runtime"
	"testing"
)

func assertEqual(expected, actual string, t *testing.T) {
	if actual != expected {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("\nOn: %s:%d\nExpected: [%s]\nActual:   [%s]", file, line, expected, actual)
	}
}
func assertPanic(callback func(), t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The callback provided should have paniced, but didn't.")
		}
	}()

	callback()
}
