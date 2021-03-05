package StringUtils

import (
	"regexp"
	"unicode"
)

// Check string for lower case instance
func IsValidPassword(str string) (bool, bool, bool, bool) {
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecialChar := false

	// Split string to rune
	for _, row := range str {
		if unicode.IsUpper(row) && unicode.IsLetter(row) {
			hasUpper = true
		}

		if unicode.IsLower(row) && unicode.IsLetter(row) {
			hasLower = true
		}

		if unicode.IsDigit(row) {
			hasNumber = true
		}

		if unicode.IsDigit(row) {
			hasNumber = true
		}

		// Set regex
		re := regexp.MustCompile("^[a-zA-Z0-9]*$")
		if !re.MatchString(string(row)) {
			hasSpecialChar = true
		}
	}

	// Return password validation
	return hasUpper, hasLower, hasNumber, hasSpecialChar
}
