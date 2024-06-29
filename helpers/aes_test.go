package helpers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestXOR(t *testing.T) {
	expected := []byte(string(rune(0b0110)))
	actual := XOR([]byte(string(rune(0b0011))), []byte(string(rune(0b0101))))
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v\nGot: %v\n", expected, actual)
	}
}

func TestEncryptAESECB(t *testing.T) {
	input := "Deez nuts"
	key := "MY_TEST_KEY_HEHE"
	encrypted := EncryptAESECB(input, key)
	decrypted := DecryptAESECB(encrypted, key)
	if encrypted != decrypted {
		t.Errorf("Encypted: %v\nDecrypted: %v\n", encrypted, decrypted)
	}
}

func TestRandomAESKey(t *testing.T) {
	fmt.Println(RandomAESKey())
}
