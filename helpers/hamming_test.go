package helpers

import "testing"

func TestHammingDistanceStr(t *testing.T) {
	expected := 37
	actual := HammingDistance("this is a test", "wokka wokka!!!")

	if expected != actual {
		t.Errorf("Expected: %d\nGot: %d\n", expected, actual)
	}
}

func TestHammingDistanceRune(t *testing.T) {
	expected := 6
	actual := hammingDistanceByte(0b11110000, 0b11001111)

	if expected != actual {
		t.Errorf("Expected: %d\nGot: %d\n", expected, actual)
	}
}
