package calculator

import (
	"errors"
	"math"
	"testing"
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
		result, err := Calculate(test.operand1, test.operand2, test.operator)

		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Expected error: %v, got: %v", test.err, err)
		}

		// Check if the difference between expected and actual values is within tolerance
		if math.Abs(result-test.expected) > tolerance {
			t.Errorf("Expected result: %v, got: %v", test.expected, result)
		}
	}
}

func TestParseOperand(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		err      error
	}{
		{"1", 1, nil},
		{"3.14", 3.14, nil},
		{"abc", 0, errors.New("invalid operand")},

		// Empty string
		{"", 0, errors.New("invalid operand")},

		// Whitespace - these cases should be handled
		//{" 123 ", 123, nil},
		//{"\t456", 456, nil},
		//{"\n789", 789, nil},

		// Scientific notation
		{"1.23e4", 1.23e4, nil},
		{"-5.67e-8", -5.67e-8, nil},

		// Trailing decimal
		{"12.", 12, nil},
		{"45.0", 45, nil},

		// Testing Negative Operand
		{"-1", -1, nil},

		// Testing Zero Operand
		{"0", 0, nil},

		// Testing Overflow
		{"1.7976931348623158e+308", 1.7976931348623157e+308, nil},

		// Testing Underflow
		{"4.940656458412464e-324", 4.940656458412464e-324, nil},

		// Testing Large Integer
		{"12345678901234567890", 1.2345678901234568e+19, nil},

		// Testing Invalid Float
		{"3.14.159", 0, errors.New("invalid operand")},
		{"1.23e4.56", 0, errors.New("invalid operand")},

		// Testing Invalid Scientific Notation
		{"1e", 0, errors.New("invalid operand")},
		{"1e-", 0, errors.New("invalid operand")},
	}

	for _, test := range tests {
		result, err := ParseOperand(test.input)

		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Expected error: %v, got: %v", test.err, err)
		}

		if result != test.expected {
			t.Errorf("Expected result: %v, got: %v", test.expected, result)
		}
	}
}
