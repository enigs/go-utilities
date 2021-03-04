package StringUtils

import (
	"strings"
)

// Normalize name
func NormalizeName(str string) string {
	// Check if string is empty
	if IsEmpty(str) {
		return ""
	}

	// trim space
	str = strings.TrimSpace(str)

	// Create name slice
	var ns []string
	str = strings.ToLower(str)
	str = strings.Title(str)

	// Break names based on spaces
	strSlice := strings.Split(str, " ")

	// Iterate through names
	for _, v := range strSlice {
		switch true {
		case v == ".":
			ns = append(ns, "")
		case v == "Jr.":
			ns = append(ns, "Jr")
		case v == "Sr.":
			ns = append(ns, "Sr")
		case v == "Ii":
			ns = append(ns, "II")
		case v == "Iii":
			ns = append(ns, "III")
		case v == "Iv":
			ns = append(ns, "IV")
		case v == "Vi":
			ns = append(ns, "VI")
		case v == "Vii":
			ns = append(ns, "VII")
		case v == "Viii":
			ns = append(ns, "VIII")
		case v == "Ix":
			ns = append(ns, "IX")
		case v == "Xi":
			ns = append(ns, "XI")
		case v == "Xii":
			ns = append(ns, "XII")
		case v == "Xiii":
			ns = append(ns, "XIII")
		case v == "Xiv":
			ns = append(ns, "XIV")
		case v == "Xv":
			ns = append(ns, "XV")
		case v == "Xvi":
			ns = append(ns, "XVI")
		case v == "Xvii":
			ns = append(ns, "XVII")
		case v == "Xviii":
			ns = append(ns, "XVIII")
		case v == "Xix":
			ns = append(ns, "XIX")
		case v == "Xx":
			ns = append(ns, "XX")
		default:
			ns = append(ns, v)
		}
	}

	// Return name
	return strings.Join(ns, " ")
}
