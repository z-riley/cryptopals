package challenges

import (
	"crypto/aes"
	crand "crypto/rand"
	"fmt"
	"math/rand"
	"strings"

	"github.com/zac460/cryptopals/helpers"
)

// 9. Implement PKCS#7 padding
func PadPKCS7(s string, blockSize int) string {
	numPadding := blockSize - len(s)%blockSize
	padded := s + strings.Repeat(string(rune(numPadding)), numPadding)
	return padded
}

// 10. Implement CBC mode
func DecryptAESCBC(ciphertext, key, iv []byte) string {
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("DecryptAESCBC() input length is not a multiple of the block size")
	}

	cb, err := aes.NewCipher(key)
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
		vec = blockCiphertext
		// 4. Store plaintext in output slice
		output += string(plaintext)
	}

	return output
}

func EncryptAESCBC(plaintext, key, iv []byte) string {
	if len(plaintext)%aes.BlockSize != 0 {
		panic("DecryptAESCBC() input length is not a multiple of the block size")
	}

	cb, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	vec := iv
	var output string
	for i := 0; i < len(plaintext)/aes.BlockSize; i++ {
		lb := i * aes.BlockSize
		ub := (i + 1) * aes.BlockSize
		// 1. XOR plaintext with IV
		toEncrypt := helpers.XOR(plaintext[lb:ub], vec)
		// 2. Encrypt the result
		buf := make([]byte, 16)
		cb.Encrypt(buf, toEncrypt)
		// 3. Set vec to the result
		vec = buf
		// 4. Store ciphertext in output slice
		output += string(buf)
	}

	return output
}

func EncryptRandomECBOrCBC(plaintext, key []byte) []byte {
	const ECBMode, CBCMode = 0, 1
	if len(plaintext)%aes.BlockSize != 0 {
		panic("DecryptAESCBC() input length is not a multiple of the block size")
	}

	// Add bytes before and after as per instructions
	beforeText := make([]byte, rand.Intn(5)+5)
	_, err := crand.Read(beforeText)
	if err != nil {
		panic(err)
	}
	afterText := make([]byte, rand.Intn(5)+5)
	_, err = crand.Read(afterText)
	if err != nil {
		panic(err)
	}
	pt := []byte(
		PadPKCS7(string(beforeText)+
			string(plaintext)+
			string(afterText),
			aes.BlockSize),
	)

	// Encrypt each block in ECB or CBC mode (chosen randomly)
	var out []byte
	for i := 0; i < len(pt)/aes.BlockSize; i++ {
		lb := i * aes.BlockSize
		ub := (i + 1) * aes.BlockSize
		plaintextBlock := pt[lb:ub]
		switch rand.Intn(2) {
		case ECBMode:
			fmt.Println("Encrypted ECB")
			out = append(out, []byte(helpers.EncryptAESECB(string(plaintextBlock), string(key)))...)
		case CBCMode:
			fmt.Println("Encrypted CBC")
			out = append(out, EncryptAESCBC(plaintextBlock, key, helpers.RandomAESKey())...)
		}
	}

	return out
}

// 11. ECB/CBC detection oracle
func EncryptionOracle(cipherText []byte) {
	// Assuming known block size (FIXME:)
	blockSize := aes.BlockSize

	for i := 0; i < len(cipherText)/blockSize; i++ {
		lb := i * blockSize
		ub := (i + 1) * blockSize
		block := cipherText[lb:ub]
		fmt.Println("Block", i, ":", DetectAESECB(string(block)))
	}

}
