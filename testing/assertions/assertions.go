// Package assertions provides custom assertion functions for testing.
package assertions

import (
	"testing"
)

// Equal asserts that two values are equal.
func Equal(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected: %v, got: %v", expected, actual)
	}
}

// NotEqual asserts that two values are not equal.
func NotEqual(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if expected == actual {
		t.Errorf("expected: %v to not equal: %v", expected, actual)
	}
}

// Nil asserts that a value is nil.
func Nil(t *testing.T, actual interface{}) {
	t.Helper()
	if actual != nil {
		t.Errorf("expected nil, got: %v", actual)
	}
}

// NotNil asserts that a value is not nil.
func NotNil(t *testing.T, actual interface{}) {
	t.Helper()
	if actual == nil {
		t.Errorf("expected non-nil value")
	}
}

// ErrorOccurred asserts that an error occurred.
func ErrorOccurred(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("expected an error, but got nil")
	}
}

// InvalidInputErrorOccurred asserts that an error occurred.
func InvalidInputErrorOccurred(t *testing.T, input string) {
	t.Helper()
	t.Errorf("Invalid expression: %s", input)
}

// NoError asserts that no error occurred.
func NoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}
}
