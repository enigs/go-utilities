package StringUtils

import (
	"fmt"
)

// Pad string with zeroes
func PadInt64ToString(amount int64, value int64) string {
	format := fmt.Sprintf("0%dd", value)
	format = "%" + format

	return fmt.Sprintf(format, amount)
}