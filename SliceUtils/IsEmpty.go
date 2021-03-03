package SliceUtils

import (
	"strings"
)

// Check if a slice is empty
func IsEmpty(args ...string) bool {
	if args == nil {
		return  true
	}

	flag := true

	for _,v := range args {
		if v = strings.TrimSpace(strings.Trim(v, " ")); len(v) > 0 {
			flag = false
		}

		if  v = strings.TrimSpace(strings.Trim(v, " ")); v != "" {
			flag = false
		}
	}

	return flag
}
