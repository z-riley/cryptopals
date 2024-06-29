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

func TestEncryptAESECB(t *testing.T) {
	expected := "The FitnessGram Pacer Test is a "
	key := []byte("YELLOW SUBMARINE")
	iv := []byte(strings.Repeat("\x00", aes.BlockSize))
	encrypted := EncryptAESCBC([]byte(expected), key, iv)
	actual := DecryptAESCBC([]byte(encrypted), key, iv)
	if expected != actual {
		t.Errorf("Expected: %v\nGot: %v\n", expected, actual)
	}
}

func TestEncryptionOracle(t *testing.T) {
	// Generate test ciphertext
	plainText := []byte("The FitnessGram Pacer Test is a multistage aerobic capacity test that progressively gets more difficult as it continues. The 20 meter pacer test will begin in 30 seconds. Line up at the start. The running speed starts slowly but gets faster each minute after you hear this signal bodeboop. A sing lap should be completed every time you hear this sound. ding Remember to run in a straight line and run as long as possible. The second time you fail to complete a lap before the sound, your test is over. The test will begin on the word start. On your mark. Get ready!... Start!!")
	key := []byte("YELLOW SUBMARINE")
	randomEncrypted := EncryptRandomECBOrCBC(plainText, key)
	fmt.Println(randomEncrypted)

	// Detect AES mode for each block
	EncryptionOracle(nil)
}
