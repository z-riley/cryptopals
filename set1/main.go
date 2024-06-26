package set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"slices"
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
			buf := make([]byte, len(b))
			for i := range b {
				buf[i] = b[i] ^ byte(char)
			}
			out = append(out, keyVal{key: char, val: string(buf)})
		}
	}
	return out
}

// 5. Repeating-key XOR
func RepeatingKeyXOR(s string, key string) string {
	contKey := strings.Repeat(key, len(s))[0:len(s)]
	func(string) {}(contKey)
	var out []byte
	for i := range s {
		out = append(out, s[i]^contKey[i])
	}
	return hex.EncodeToString(out)

}

// 6. Break repeating-key XOR
func BreakRepeatingKeyXOR(b64 string) {
	const maxKeysize = 40

	bytes, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		panic(err)
	}

	if len(bytes) < maxKeysize*3 {
		panic("BreakRepeatingKeyXOR() input string too short")
	}

	// Let KEYSIZE be the guessed length of the key; try values
	// from 2 to (say) 40.
	// For each KEYSIZE, take the first KEYSIZE worth of bytes,
	// and the second KEYSIZE worth of bytes, and find the edit
	// distance between them. Normalize this result by dividing
	// by KEYSIZE.
	type sizeToDist struct {
		keysize int
		dist    float64
	}
	var allKeysizeDists []sizeToDist
	for keysize := 2; keysize <= maxKeysize; keysize++ {
		b := bytes
		k := keysize

		// The KEYSIZE with the smallest normalized edit distance
		// is probably the key. Take 4 KEYSIZE blocks and average
		// the distances.
		// TODO: implement a neater way of doing this - go func?
		d := helpers.HammingDistance(b[0:k], b[k:2*k])
		hammingDist1 := float64(d)

		d = helpers.HammingDistance(b[2*k:3*k], b[3*k:4*k])
		hammingDist2 := float64(d)

		d = helpers.HammingDistance(b[4*k:5*k], b[5*k:6*k])
		hammingDist3 := float64(d)

		avDistance := (hammingDist1 + hammingDist2 + hammingDist3) / 3
		normalisedDist := avDistance / float64(k)
		allKeysizeDists = append(allKeysizeDists, sizeToDist{k, normalisedDist})
	}

	// The KEYSIZE with the smallest normalized edit distance
	// is probably the key. You could proceed perhaps with the
	// smallest 2-3 KEYSIZE values.
	slices.SortFunc(allKeysizeDists, func(a, b sizeToDist) int {
		switch {
		case a.dist < b.dist:
			return -1
		case a.dist > b.dist:
			return 1
		default:
			return 0
		}
	})
	bestKeysizeDists := allKeysizeDists[0:3]

	// Now that you probably know the KEYSIZE: break the ciphertext
	// into blocks of KEYSIZE length.
	keysize := bestKeysizeDists[0].keysize // using best single keysize for now

	var blocks [][]byte
	for i := 0; i < len(bytes); i += keysize {
		if i+keysize > len(bytes)-1 {
			// Partial block remaining
			blocks = append(blocks, bytes[i:])
		} else {
			blocks = append(blocks, bytes[i:i+keysize])
		}
	}

	// Now transpose the blocks: make a block that is the first
	// byte of every block, and a block that is the second byte of
	// every block, and so on.
	tBlocks := make([]string, keysize)
	for _, block := range blocks {
		for i, char := range block {
			tBlocks[i] += string(char)
		}
	}

	// Solve each block as if it was single-character XOR. You
	// already have code to do this.
	// Just solving first block for now.............
	block := tBlocks[0]
	// block, err := base64.StdEncoding.DecodeString(tBlocks[0])
	// if err != nil {
	// 	panic(err)
	// }

	decryptedBlock := []keyVal{}
	for _, char := range helpers.PrintableASCII {
		buf := make([]byte, keysize)
		for i := range keysize {
			buf[i] = block[i] ^ byte(char)
		}
		decryptedBlock = append(decryptedBlock, keyVal{key: char, val: string(buf)})
	}

	var results []string
	for _, res := range decryptedBlock {
		if helpers.IsASCII(res.val) && helpers.IsPrintableASCII(res.val) {
			const thresh = 1 // determined empirically
			if helpers.IsLikelyEnglishThresh(res.val, thresh) {
				results = append(results, fmt.Sprintf("%c - %s\n", res.key, res.val))
			}
		}
	}
	fmt.Println(results)

	// For each block, the single-byte XOR key that produces the
	// best looking histogram is the repeating-key XOR key byte for
	// that block. Put them together and you have the key.

}

func UNUSED(...any) {}
