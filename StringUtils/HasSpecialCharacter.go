package StringUtils

import (
	"regexp"
)

// Check string for special character instance
func HasSpecialCharacter(str string) bool {
	// Set alphanumeric
	re := regexp.MustCompile("^[a-zA-Z0-9]*$")

	// Split string to rune
	for _, row := range str {
		if !re.MatchString(string(row)) {
			return true
		}
	}

	// Return false if no special character detected
	return false
}

