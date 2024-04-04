package runners

import (
	"testing"
)

// IntegrationTestRunner executes integration tests.
type IntegrationTestRunner struct{}

// Run executes integration tests.
func (itr *IntegrationTestRunner) Run(t *testing.T, tests []func(*testing.T)) {
	for _, test := range tests {
		test(t)
	}
}
