package main

import (
	"testing"
)

// Test cases for isBinary function
func TestIsBinary(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"1010", true},    // Valid binary number
		{"1234", false},   // Contains digits other than 0 and 1
		{"", false},       // Empty string is not a valid binary number
		{"10a01", false},  // Contains a non-digit character
	}

	for _, test := range tests {
		result := isBinary(test.input)
		if result != test.expected {
			t.Errorf("isBinary(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
