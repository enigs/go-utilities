package StringUtils

import (
	"github.com/guregu/null"

	"encoding/json"
)

// Converts json item into null string
func JsonToString(item interface{}) (null.String, error) {
	// Convert item to bytes
	itemByte, err := json.Marshal(item)
	if err != nil {
		return null.NewString("", false), err
	}

	if string(itemByte) == "null" {
		return null.NewString("", false), nil
	} else {
		return null.NewString(string(itemByte), true), nil
	}
}
