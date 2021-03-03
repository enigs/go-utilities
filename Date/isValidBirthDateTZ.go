package Date

import (
	"fmt"
	"time"

	"github.com/bearbin/go-age"
	"github.com/enigs/go-utilities/StringUtils"
)

// Checks birth date if valid using Asia/Manila timezone
func IsValidBirthDateTZ(day, month, year, minAge int64, tz string) bool {

	// Load asia/manila location
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return  false
	}

	// Set format
	format := fmt.Sprintf("%s-%s-%sT00:00:01+08:00", StringUtils.PadInt64ToString(year, 4), StringUtils.PadInt64ToString(month, 2), StringUtils.PadInt64ToString(day, 2))

	// Load birth date into date time GMT +8
	var bd time.Time
	bd, err = time.Parse("2006-01-02T15:04:05Z07:00", format)
	if err != nil {
		return false
	}

	// Set birthday
	birthday := time.Date(bd.Year(), bd.Month(), bd.Day(), 0, 0, 0, 0, loc)

	// If min age is set as 0 then it is automatically a valid date
	if minAge == 0 {
		return true
	}

	// Check age if current age is valid
	if age.Age(birthday) < int(minAge) {
		return false
	}

	return true
}