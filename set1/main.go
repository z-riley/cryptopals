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

/* Challenge 5

Implement repeating-key XOR
Here is the opening stanza of an important work of the English language:

Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal
Encrypt it, under the key "ICE", using repeating-key XOR.

In repeating-key XOR, you'll sequentially apply each byte of the key; the first byte of plaintext will be XOR'd against I, the next C, the next E, then I again for the 4th byte, and so on.

It should come out to:

0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272
a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f
Encrypt a bunch of stuff using your repeating-key XOR function. Encrypt your mail. Encrypt your password file. Your .sig file. Get a feel for it. I promise, we aren't wasting your time with this.
*/

// 5. Repeating-key XOR
func RepeatingKeyXOR() {

}
