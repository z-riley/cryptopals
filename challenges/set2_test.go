package challenges

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"strings"
	"testing"
)

func TestPadPKCS7(t *testing.T) {
	expected := "YELLOW SUBMARINE\x04\x04\x04\x04"
	actual := PadPKCS7("YELLOW SUBMARINE", 20)
	if expected != actual {
		t.Errorf("Expected: %v\nGot: %v\n", expected, actual)
	}
}

func TestDecryptAESCBC(t *testing.T) {
	key := []byte("YELLOW SUBMARINE")
	iv := []byte(strings.Repeat("\x00", aes.BlockSize))
	input, err := base64.StdEncoding.DecodeString(ch10input)
	if err != nil {
		t.Fatal(err)
	}
	actual := DecryptAESCBC(input, key, iv)
	fmt.Println(actual)
}
