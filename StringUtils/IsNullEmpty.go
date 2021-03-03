package StringUtils

import (
	"strings"

	"github.com/guregu/null"
)

// Check if null string is empty
func IsNullEmpty(str null.String) bool {
	if !str.Valid || len(strings.TrimSpace(str.String)) == 0 || strings.TrimSpace(str.String) == "" {
		return true
	}

	return false
}