package challenges

import (
	"strings"
)

// 9. Implement PKCS#7 padding
func PadPKCS7(s string, blockSize int) string {
	if len(s) > blockSize {
		panic("PadPKCS7() string cannot be larger than the block size")
	}
	numPadding := blockSize - len(s)
	padded := s + strings.Repeat(string(rune(numPadding)), numPadding)
	return padded
}
