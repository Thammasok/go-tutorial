package tests

import (
	"testing"

	L "../libs"
)

func TestCalculate(t *testing.T) {
	if L.Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		if output := L.Calculate(test.input); output != test.expected {
			t.Error("Test Failed: inputted: {}, expected: {}, recieved: {}", test.input, test.expected, output)
		}
	}
}
