package helpers

import (
	"crypto/aes"
)

// XOR performs an XOR operation on two strings
func XOR(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("FixedXOR() strings must be same length")
	}
	buf := make([]byte, len(a))
	for i := range a {
		buf[i] = a[i] ^ b[i]
	}
	return buf
}

// EncryptAESECB encrypts a string using AES and ECB mode
func EncryptAESECB(s string, key string) string {
	b := []byte(s)
	cb, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	buf := make([]byte, len(b))
	for i := 0; i < len(b)/aes.BlockSize; i++ {
		lb := i * aes.BlockSize
		ub := (i + 1) * aes.BlockSize
		cb.Encrypt(buf[lb:ub], b[lb:ub])
	}
	return string(buf)
}

// DecryptAESECB decrypts a string using AES and ECB mode
func DecryptAESECB(s string, key string) string {
	b := []byte(s)
	cb, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	buf := make([]byte, len(b))
	for i := 0; i < len(b)/aes.BlockSize; i++ {
		lb := i * aes.BlockSize
		ub := (i + 1) * aes.BlockSize
		cb.Decrypt(buf[lb:ub], b[lb:ub])
	}
	return string(buf)
}
