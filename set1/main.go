package set1

import (
	"encoding/base64"
	"encoding/hex"
)

// 1. Convert hex to base64
func HexToBase64(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return []byte(base64.StdEncoding.EncodeToString(b))
}
