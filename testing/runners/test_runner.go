package runners

import (
	"testing"
)

// TestRunner defines common interfaces for test runners.
type TestRunner interface {
	Run(t *testing.T, tests []func(*testing.T))
}
