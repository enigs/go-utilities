package SliceUtils

// Check if a string is within the provided slice
func InSliceSensitive(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}

	return false
}

