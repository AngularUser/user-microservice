package helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/vineela-devarashetty/user-microservice/model"
	"testing"
)

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

func TestValidateUser_ValidUser(t *testing.T) {
	user := model.User{
		UserID: "123",
		Name:   "John Doe",
		Email:  "john@example.com",
		DOB:    "2000-01-01",
	}

	if err := ValidateUser(user); err != nil {
		t.Errorf("Expected no validation errors, but got %v", err)
	}
}

func TestValidateUser_InvalidUser(t *testing.T) {
	user := model.User{
		UserID: "",
		Name:   "John",          // Empty name
		Email:  "invalid-email", // Invalid email format
		DOB:    "",
	}

	if err := ValidateUser(user); err == nil {
		t.Error("Expected validation errors, but got none")
	} else {
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			t.Error("Expected a ValidationErrors type error")
		}

	}
}
