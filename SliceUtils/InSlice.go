package SliceUtils

import (
	"strings"
)

// Check if a string is within the provided slice
func InSlice(slice []string, val string) bool {
	for _, item := range slice {
		if strings.ToLower(item) == strings.ToLower(val) {
			return true
		}
	}

	return false
}
