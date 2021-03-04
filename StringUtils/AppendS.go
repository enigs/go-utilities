package StringUtils

import (
	"fmt"
	"strings"
)

// Check append 's or ' on names
func AppendS(firstName, lastName string) string {
	ln := lastName[len(lastName)-1:]

	if strings.ToLower(ln) == "s" {
		lastName = fmt.Sprintf("%s'", lastName)
	} else {
		lastName = fmt.Sprintf("%s's", lastName)
	}

	return fmt.Sprintf("%s %s", firstName, lastName)
}
