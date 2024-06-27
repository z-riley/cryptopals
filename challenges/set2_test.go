package challenges

import "testing"

func TestPadPKCS7(t *testing.T) {
	expected := "YELLOW SUBMARINE\x04\x04\x04\x04"
	actual := PadPKCS7("YELLOW SUBMARINE", 20)
	if expected != actual {
		t.Errorf("Expected: %v\nGot: %v\n", expected, actual)
	}
}
