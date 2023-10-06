package helper

import (
	"testing"
)

func TestIsValidName(t *testing.T) {
	// Test cases for IsValidName function
	tests := []struct {
		input    string
		expected bool
	}{
		{"valid_name", true},                     // Valid username
		{"invalid name", false},                  // Contains a space (invalid character)
		{"low", false},                           // Too short
		{"very_long_username_1234567890", false}, // Too long
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := IsValidName(test.input)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v for input %s", test.expected, result, test.input)
			}
		})
	}
}

func TestIsValidEmail(t *testing.T) {
	// Test cases for IsValidEmail function
	tests := []struct {
		input    string
		expected bool
	}{
		{"valid.email@example.com", true},           // Valid email
		{"invalid-email", false},                    // Missing '@' character
		{"invalid@example", false},                  // Missing top-level domain (TLD)
		{"invalid@.com", false},                     // Missing domain name
		{"valid+email@example.com", true},           // Valid email with '+'
		{"valid.email@subdomain.example.com", true}, // Valid email with subdomain
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := IsValidEmail(test.input)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v for input %s", test.expected, result, test.input)
			}
		})
	}
}

func TestIsValidDOB(t *testing.T) {
	// Test cases for IsValidDOB function
	tests := []struct {
		input    string
		expected bool
	}{
		{"2022-01-01", true},  // Valid date of birth
		{"2022-13-01", false}, // Invalid month (13)
		{"2022-01-32", false}, // Invalid day (32)
		{"22-01-01", false},   // Invalid year (short)
		{"2022/01/01", false}, // Invalid separator
		{"2022-01", false},    // Incomplete date
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := IsValidDOB(test.input)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v for input %s", test.expected, result, test.input)
			}
		})
	}
}
