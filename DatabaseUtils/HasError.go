package DatabaseUtils

// Check if gorm has valid error
func HasError(err error) bool {
	if err == nil || err.Error() == "record not found" {
		return false
	}

	return true
}
