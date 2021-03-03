package ByteUtils

import (
	"encoding/base64"
)

// Generates url safe base64 encoded string where byte = n
func GenerateRandomURLSafeString(n int) (string, error) {
	b, err := GenerateRandomBytes(n)

	return base64.URLEncoding.EncodeToString(b), err
}

