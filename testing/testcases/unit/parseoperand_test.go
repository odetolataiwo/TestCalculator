package unit

import (
	"errors"
	"testing"

	"TestCalculator/calculator"
	"TestCalculator/testing/assertions"
)

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
		result, err := calculator.ParseOperand(test.input)
		if err != nil && err.Error() != test.err.Error() {
			assertions.Equal(t, test.err, err)
		}

		if result != test.expected {
			assertions.Equal(t, test.expected, result)
		}
	}
}
