package Date

import (
	"fmt"
	"time"

	"digita_api/application/Library/StringUtils"
)

// Checks if date is valid
func IsValidDate(day, month, year int64) bool {
	// Set format
	format := fmt.Sprintf("%s-%s-%s",
		StringUtils.PadInt64ToString(year, 4),
		StringUtils.PadInt64ToString(month, 2),
		StringUtils.PadInt64ToString(day, 2))

	// Load date into date time GMT +8
	dte, err := time.Parse("2006-01-02", format)
	if err != nil {
		return false
	}

	if int(day) != dte.Day() ||
		int(dte.Month()) != int(month) ||
		dte.Year() != int(year) {
		return false
	}

	return true
}