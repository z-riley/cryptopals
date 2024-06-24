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

// 2. Fixed XOR
func FixedXOR(s string, xor string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	bxor, err := hex.DecodeString(xor)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, len(b))
	for i := range b {
		buf[i] = b[i] ^ bxor[i]
	}
	return []byte(hex.EncodeToString(buf))
}
