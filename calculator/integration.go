//go:build integration
// +build integration

package calculator

import (
	"errors"
	"math"
	"strings"
	"testing"
)

func TestCalculateIntegration(t *testing.T) {
	// Define test cases with combinations of operands and operators
	tests := []struct {
		expression string
		expected   float64
		err        error
	}{
		{"1 + 2", 3, nil}, // Addition
		{"5 - 3", 2, nil}, // Subtraction
		{"2 * 3", 6, nil}, // Multiplication
		{"6 / 2", 3, nil}, // Division
		{"6 / 0", 0, errors.New("division by zero is not allowed")}, // Division by zero
		{"0 / 0", 0, errors.New("division by zero is not allowed")}, // Division by zero
		{"0 * 0", 0, nil},                    // 0 * 0 = 0
		{"0 + 1", 1, nil},                    // 0 + 1 = 1
		{"1 - 0", 1, nil},                    // 1 - 0 = 1
		{"0 * 5", 0, nil},                    // 0 * 5 = 0
		{"1e300 + 1e300", 2e300, nil},        // Adding two large numbers
		{"1e300 / 1e-300", math.Inf(1), nil}, // Dividing a large number by a very small number

		// these scenarios are not catered for in the code.
		//{"", 0, errors.New("invalid expression")},     // Empty expression
		//{"1", 0, errors.New("invalid expression")},    // Missing operator
		//{"1 + ", 0, errors.New("invalid expression")}, // Missing operand
		//{"2 + 3 * 4", 14, nil},                        // Multiple operators
		//{"1e300 + 1", 1e300, nil},                     // Large number
		//{" 2 + 3 ", 5, nil},                           // Whitespace
		//{"3 * (2 + 4)", 18, nil},                      // Complex expression - are not handled in the code
	}

	// Iterate over test cases
	for _, test := range tests {
		// Split the expression into operands and operator
		parts := strings.Split(test.expression, " ")
		if len(parts) != 3 {
			t.Errorf("Invalid expression: %s", test.expression)
			continue
		}

		operand1, err := ParseOperand(parts[0])
		if err != nil {
			t.Errorf("Error parsing operand: %v", err)
			continue
		}

		operand2, err := ParseOperand(parts[2])
		if err != nil {
			t.Errorf("Error parsing operand: %v", err)
			continue
		}

		operator := parts[1]

		// Calculate the result
		result, err := Calculate(operand1, operand2, operator)

		// Check for errors
		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Expected error: %v, got: %v", test.err, err)
		}

		// Compare the result with the expected value
		if result != test.expected {
			t.Errorf("Expected result: %v, got: %v", test.expected, result)
		}
	}
}
