package calc

import (
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		firstNum  int
		secondNum int
		expected  int
	}{
		{1, 2, 3},      // Test case 1
		{-1, 1, 0},     // Test case 2
		{0, 0, 0},      // Test case 3
		{100, -50, 50}, // Test case 4
		//{1, 2, 4},      // Test case 5  ** SHOULD FAIL **
	}

	for _, test := range tests {
		result, err := Sum(test.firstNum, test.secondNum)
		if err != nil {
			t.Errorf("Sum(%d, %d) returned an error: %v", test.firstNum, test.secondNum, err)
		}
		if result != test.expected {
			t.Errorf("Sum(%d, %d) = %d; expected %d", test.firstNum, test.secondNum, result, test.expected)
		}
	}
}
