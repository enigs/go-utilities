package StringUtils

import (
	"regexp"
)

// Check if string is contact number
func IsContactNum(str string) bool {
	isContactNumber := regexp.MustCompile(`^[()0-9\-\+]+$`).MatchString

	if isContactNumber(str) {
		return true
	}

	return false
}

