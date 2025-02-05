package challenges

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/zac460/cryptopals/helpers"
)

func TestHexToBase64(t *testing.T) {
	expected := []byte("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t")
	actual := HexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %s\nGot: %s", expected, actual)
	}

}

func TestFixedXOR(t *testing.T) {
	expected := []byte("746865206b696420646f6e277420706c6179")
	actual := FixedXORHex(
		"1c0111001f010100061a024b53535009181c",
		"686974207468652062756c6c277320657965",
	)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %s\nGot: %s", expected, actual)
	}
}

func TestSingleByteXORCipher(t *testing.T) {
	result := SingleByteXORCipher("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	// Answer found by examining the result:
	// The string was XOR'd against 'X'
	fmt.Println(result)
}

func TestSingleCharacterXOR(t *testing.T) {
	decryptedStrings := SingleCharacterXOR(ch4Input)
	var results []string
	for _, res := range decryptedStrings {
		if helpers.IsASCII(res.val) && helpers.IsPrintableASCII(res.val) {
			const thresh = 61 // determined empirically
			if helpers.IsLikelyEnglishThresh(res.val, thresh) {
				results = append(results, fmt.Sprintf("%c - %s\n", res.key, res.val))
			}
		}
	}

	// Answer found by examining the results:
	// The string is "Now that the party is jumping\n"
	fmt.Println(results)
}

func TestRepeatingKeyXOR(t *testing.T) {
	expected := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272" +
		"a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	actual := RepeatingKeyXOR(`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`, "ICE")

	if expected != actual {
		t.Errorf("Expected: %s\nGot: %s\n", expected, actual)
	}
}

func TestBreakRepeatingKeyXOR(t *testing.T) {
	BreakRepeatingKeyXOR(ch6Input)
}

func TestDecryptAESECB(t *testing.T) {
	expectedPhrase := "Play that funky music"
	decoded, err := base64.StdEncoding.DecodeString(ch7input)
	if err != nil {
		t.Fatal(err)
	}
	actual := DecryptAESECB(string(decoded), "YELLOW SUBMARINE")
	if !strings.Contains(actual, expectedPhrase) {
		t.Errorf("Result did not to contain \"%s\":\n%s", expectedPhrase, actual)
	}

}

func TestDetectAESECB(t *testing.T) {
	expected := map[string]int{"08649af70dc06f4f": 4, "d5d2d69c744cd283": 4}
	// Feed it one line at a time
	lines := strings.Split(ch8input, "\n")
	for lineNum, s := range lines {
		actual := DetectAESECB(s)
		if len(actual) > 0 {
			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("\nExpected: %v\nGot: %v\n", expected, actual)
			}
			fmt.Printf("Line %d encrypted with AES ECB:\n%s\n", lineNum, s)
			return
		}
	}
}
