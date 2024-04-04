// Package runners defines the interface for running tests.
package runners

import (
	"testing"
)

// UnitTestRunner executes unit tests.
type UnitTestRunner struct{}

// Run executes unit tests.
func (utr *UnitTestRunner) Run(t *testing.T, tests []func(*testing.T)) {
	for _, test := range tests {
		test(t)
	}
}
