package helper

import (
	"regexp"
	"strconv"
	"time"
)

// isValidName checks if a user name is valid based on certain rules.
func IsValidName(name string) bool {
	// Define the allowed character set (alphanumeric, underscore, hyphen)
	validChars := "^[a-zA-Z0-9_-]+$"

	// Define minimum and maximum user name lengths
	minLength := 4
	maxLength := 20

	// Check length requirement
	if len(name) < minLength || len(name) > maxLength {
		return false
	}

	// Check character set using regular expression
	match, _ := regexp.MatchString(validChars, name)

	return match
}

func IsValidEmail(email string) bool {
	// Basic email format validation using regular expression
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
	return regexp.MustCompile(emailRegex).MatchString(email)
}

func IsValidDOB(dob string) bool {
	// Basic date of birth format validation (YYYY-MM-DD)
	dobRegex := `^\d{4}-\d{2}-\d{2}$`
	if !regexp.MustCompile(dobRegex).MatchString(dob) {
		return false
	}

	// Extract year, month, and day from the DOB string
	year := dob[0:4]
	month := dob[5:7]
	day := dob[8:10]

	// Convert year, month, and day to integers
	yearInt := atoi(year)
	monthInt := atoi(month)
	dayInt := atoi(day)

	// Check for valid year (e.g., 1900 to the current year)
	currentYear := time.Now().Year()
	if yearInt < 1900 || yearInt > currentYear {
		return false
	}

	// Check for valid month (1 to 12)
	if monthInt < 1 || monthInt > 12 {
		return false
	}

	// Check for valid day based on month
	switch monthInt {
	case 2: // February
		// Check for leap year (29 days) or non-leap year (28 days)
		if (yearInt%4 == 0 && yearInt%100 != 0) || yearInt%400 == 0 {
			return dayInt >= 1 && dayInt <= 29
		} else {
			return dayInt >= 1 && dayInt <= 28
		}
	case 4, 6, 9, 11: // April, June, September, November (30 days)
		return dayInt >= 1 && dayInt <= 30
	default: // Months with 31 days
		return dayInt >= 1 && dayInt <= 31
	}
}

// atoi is a helper function to convert a string to an integer.
func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0 // Return 0 for invalid or non-numeric strings
	}
	return i
}
