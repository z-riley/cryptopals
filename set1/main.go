package set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/zac460/cryptopals/helpers"
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

// 3. Single-byte XOR cipher
func SingleByteXORCipher(s string) []string {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	// Try with each ASCII character
	out := make([]string, len(helpers.PrintableASCII))
	for _, char := range helpers.PrintableASCII {
		buf := make([]byte, len(b))
		for i := range b {
			buf[i] = b[i] ^ byte(char)
		}
		out = append(out, fmt.Sprintf("%c - %s\n", char, buf))
	}
	return out
}

type keyVal struct {
	key rune
	val string
}

// 4. Detect single-character XOR
func SingleCharacterXOR(s string) []keyVal {
	out := []keyVal{}
	for _, line := range strings.Split(s, "\n") {
		b, err := hex.DecodeString(line)
		if err != nil {
			panic(err)
		}

		for _, char := range helpers.PrintableASCII {
			// for _, char := range []rune{'5'} {
			buf := make([]byte, len(b))
			for i := range b {
				buf[i] = b[i] ^ byte(char)
			}
			out = append(out, keyVal{key: char, val: string(buf)})
		}
	}
	return out
}
