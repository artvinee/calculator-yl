package main

import (
	"calculator-yl/core"
	"testing"
)

func TestCalculateExpression(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
		shouldFail bool
	}{
		// Валидные выражения
		{"2 + 2", 4, false},
		{"((2 * 9) - 3) - 5 / 2", 12.5, false},
		{"(5 + 3) * 2", 16, false},
		{"-5 + 3", -2, false},
		{"((2 + 3) * (4 - 1))", 15, false},
		{"-(-1)", 1, false},
		{"-(-(-1))", -1, false},
		{"2 * (-3)", -6, false},
		{"(-2) * (-3)", 6, false},
		{"(-2) + (-3)", -5, false},
		{"10 / (-2)", -5, false},
		{"-5 * (-2 + 3)", -5, false},
		{"(-4) * (-5) + (-3)", 17, false},
		{"-(2 + 3)", -5, false},
		{"(-2.5) * (-4)", 10, false},
		{"-(3 + (-2))", -1, false},

		// Невалидные выражения
		{"10 / 0", 0, true},
		{"", 0, true},
		{"   ", 0, true},
		{"2 + a", 0, true},
		{"2 @ 3", 0, true},
		{"(2 + 3", 0, true},
		{"2 + 3)", 0, true},
		{"2 ++ 3", 0, true},
		{"2 3", 0, true},
		{"2 +", 0, true},
		{"()", 0, true},
		{"2.5.6 + 1", 0, true},
		{"+ 2 + 3", 0, true},
		{"2 + 3 +", 0, true},
		{"2 + () + 1", 0, true},
		{"2.3.4 + 1", 0, true},
		{"2 */ 3", 0, true},
		{"2plus3", 0, true},
		{"((()))", 0, true},
		{"(-)", 0, true},
		{"2-", 0, true},
		{"(-", 0, true},
		{"2 - - 3", 0, true},
		{"(-2", 0, true},
	}

	for _, test := range tests {
		result, err := core.CalculateExpression(test.expression)
		if test.shouldFail {
			if err == nil {
				t.Errorf("Expected an error for expression: %s", test.expression)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for expression: %s, error: %v", test.expression, err)
			}
			if result != test.expected {
				t.Errorf("For expression: %s, expected: %v, got: %v", test.expression, test.expected, result)
			}
		}
	}
}
