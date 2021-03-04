package StringUtils

import (
	"regexp"
)

// Check if string is valid name
func IsName(str string) bool {
	name := regexp.MustCompile("^[\\p{L} \\.-]+$")
	return name.MatchString(str)
}
