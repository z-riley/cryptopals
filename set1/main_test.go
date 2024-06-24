package set1

import (
	"fmt"
	"reflect"
	"testing"
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
	actual := FixedXOR(
		"1c0111001f010100061a024b53535009181c",
		"686974207468652062756c6c277320657965",
	)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %s\nGot: %s", expected, actual)
	}
}

func TestSingleByteXORCipher(t *testing.T) {
	result := SingleByteXORCipher("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	fmt.Println(result)
	// Correct answer found by examining the result.
	// Answer: string was XOR'd against 'X'
}
