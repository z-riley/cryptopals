package challenges

import (
	"crypto/aes"
	"fmt"
	"strings"

	"github.com/zac460/cryptopals/helpers"
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

// 10. Implement CBC mode
func DecryptAESCBC(ciphertext, key, iv []byte) string {
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("DecryptAESCBC() input length is not a multiple of the block size")
	}

	cb, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	vec := iv
	var output string
	for i := 0; i < len(ciphertext)/aes.BlockSize; i++ {
		lb := i * aes.BlockSize
		ub := (i + 1) * aes.BlockSize
		// 1. Decrypt ciphertext
		buf := make([]byte, 16)
		blockCiphertext := ciphertext[lb:ub]
		cb.Decrypt(buf, []byte(blockCiphertext))
		// 2. XOR decrypted data with IV to get plaintext
		plaintext := helpers.XOR(buf, vec)
		// 3. Set vec to the previous ciphertext
		vec = []byte(blockCiphertext)
		// 4. Store plaintext in output slice
		output += string(plaintext)
	}

	return output
}

// not my code
// Decrypt cipherText with key using Cipher Block Chaining (CBC) mode.
func decryptAESCBC(cipherText []byte, key []byte, iv []byte) ([]byte, error) {
	blockSize := 16
	if len(key) < blockSize {
		return nil, fmt.Errorf("key size must be %d", blockSize)
	}

	if len(iv) < blockSize {
		return nil, fmt.Errorf("iv size must be %d", blockSize)
	}

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("could not initialize AES: %w", err)
	}

	plainText := make([]byte, len(cipherText))
	buffer := make([]byte, blockSize)
	plainTextBuffer := make([]byte, blockSize)
	lastCipherText := make([]byte, blockSize)
	copy(lastCipherText, iv)
	for i := 0; i < (len(plainText) / blockSize); i++ {
		start := i * blockSize
		end := (i + 1) * blockSize

		cipher.Decrypt(buffer, cipherText[start:end])
		for j := 0; j < blockSize; j++ {
			plainTextBuffer[j] = lastCipherText[j] ^ buffer[j]
		}

		copy(plainText[start:end], plainTextBuffer)
		copy(lastCipherText, cipherText[start:end])
	}

	return plainText, nil
}
