package StringUtils

import (
	"regexp"
)

// Checks if string is username
func  IsUsername(str string) bool {
	username := regexp.MustCompile("^[\\p{L}\\d_\\.-]+$")
	return username.MatchString(str)
}

