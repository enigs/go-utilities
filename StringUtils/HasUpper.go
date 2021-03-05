package StringUtils

import (
	"unicode"
)

// Check string for upper case instance
func HasUpper(str string) bool {
	// Split string to rune
	for _, row := range str {
		if unicode.IsUpper(row) && unicode.IsLetter(row) {
			return true
		}
	}

	// Return false if no upper case string detected
	return false
}
