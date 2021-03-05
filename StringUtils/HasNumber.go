package StringUtils

import (
	"unicode"
)

// Check string for number instance
func HasNumber(str string) bool {
	// Split string to rune
	for _, row := range str {
		if unicode.IsDigit(row) {
			return true
		}
	}

	// Return false if no number detected
	return false
}
