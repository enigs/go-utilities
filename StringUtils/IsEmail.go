package StringUtils

import (
	"regexp"
)

// Check if string is email
func  IsEmail(str string) bool {
	re := regexp.MustCompile("^[\\p{L}\\d.!#$%&'*+/=?^_`{|}~-]+@[\\p{L}\\d](?:[\\p{L}\\d-_]{0,61}[\\p{L}\\d])?(?:\\.[\\p{L}\\d](?:[\\p{L}\\d-_]{0,61}[\\p{L}\\d])?)*$")

	return re.MatchString(str)
}
