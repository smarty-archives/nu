package nu

import (
	"errors"
	"fmt"
)

// Error is a convenient combination of errors.New and fmt.Sprint.
func Error(args ...interface{}) error {
	return errors.New(fmt.Sprint(args...))
}

// Errorln is a convenient combination of errors.New and fmt.Sprintln.
func Errorln(args ...interface{}) error {
	return errors.New(fmt.Sprintln(args...))
}

// Errorf is a convenient combination of errors.New and fmt.Sprintf.
func Errorf(format string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(format, args...))
}
