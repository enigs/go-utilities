package StringUtils

import (
	"strings"
)

// split email
func SplitEmail(str string) (string, string){
	// Split string with @ sign
	es1 := strings.Split(str, "@")

	// Check if email slice is nil
	if es1 == nil || len(es1) == 0 {
		return "", ""
	}

	// Check if email slice consists of only 1 slice
	if len(es1) == 1 {
		return es1[0], ""
	}

	// Split second part dot
	es2 := strings.Split(es1[1], ".")

	// Check if email slice is nil
	if es2 == nil || len(es2) == 0 {
		return "", ""
	}

	return es1[0], es2[0]
}