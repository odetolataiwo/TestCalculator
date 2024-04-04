package unit

import (
	"errors"
	"math"
	"testing"

	"TestCalculator/calculator"
	"TestCalculator/testing/assertions"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		operand1 float64
		operator string
		operand2 float64
		expected float64
		err      error
	}{
		{1, "+", 1, 2, nil},
		{5, "-", 3, 2, nil},
		{2, "*", 3, 6, nil},
		{6, "/", 2, 3, nil},
		{6, "/", 0, 0, errors.New("division by zero is not allowed")},
		{1, "%", 1, 0, errors.New("invalid operator")},

		// Testing Floating Point Precision
		{1.1, "+", 2.2, 3.3, nil},

		// Testing Large Numbers
		{999999999, "+", 1, 1000000000, nil},
		{1e300, "+", 1e300, 2e300, nil},
		{-1e300, "-", 1e300, -2e300, nil},

		// Infinity and NaN
		{math.Inf(1), "+", 1, math.Inf(1), nil},
		{1, "+", math.NaN(), math.NaN(), nil},

		// Exponents
		{1e2, "*", 3, 300, nil},
		{5e-3, "/", 2, 2.5e-3, nil},

		// Mixed signs
		{-5, "*", 2, -10, nil},
		{3, "/", -4, -0.75, nil},

		// Zero as operand
		{0, "+", 1, 1, nil},
		{4, "-", 0, 4, nil},
		{0, "*", 5, 0, nil},
		{10, "/", 0, 0, errors.New("division by zero is not allowed")},

		// Division by zero with both operands 0 panics. This scenario should also be handled in the code.
		//{0, "/", 0, math.NaN(), nil},

		// Testing Negative Numbers
		{5, "-", 8, -3, nil},
		{-5, "+", 8, 3, nil},
	}

	// Define a tolerance for floating-point comparison
	tolerance := 1e-10

	for _, test := range tests {
		result, err := calculator.Calculate(test.operand1, test.operand2, test.operator)
		if err != nil && err.Error() != test.err.Error() {
			assertions.Equal(t, test.err, err)
		}

		// Check if the difference between expected and actual values is within tolerance
		if math.Abs(result-test.expected) > tolerance {
			assertions.Equal(t, test.expected, result)
		}
	}
}
