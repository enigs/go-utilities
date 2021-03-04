package StringUtils

import (
	"regexp"
)

// Check if string is alpha numeric
func IsAlphaNum(str string) bool {
	isAesEncoded := regexp.MustCompile(`^[A-Za-z0-9ñÑ]+$`).MatchString

	if isAesEncoded(str) {
		return true
	}

	return false
}