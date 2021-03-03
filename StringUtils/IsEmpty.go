package StringUtils

import (
	"strings"
)

// Check if string is empty
func IsEmpty(str string) bool {
	if len(strings.TrimSpace(str)) == 0 || strings.TrimSpace(str) == "" {
		return true
	}

	return false
}