package helpers

import "reflect"

// HammingDistance calculates the Hamming Distance between two variables
func HammingDistance(a, b any) int {

	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		panic("HammingDistance() types must be the same")
	}

	distance := 0
	switch a.(type) {
	case []byte:
		a, b := a.([]byte), b.([]byte)
		if len(a) != len(b) {
			panic("HammingDistance() bytes have different lengths")
		}
		for i := range a {
			distance += hammingDistanceByte(a[i], b[i])
		}
		return distance
	case string:
		a, b := a.(string), b.(string)
		if len(a) != len(b) {
			panic("HammingDistance() strings have different lengths")
		}
		for i := range a {
			distance += hammingDistanceByte(a[i], b[i])
		}
		return distance
	case byte:
		distance += hammingDistanceByte(a.(byte), b.(byte))
	default:
		panic("HammingDistance() unsupported type")
	}
	return distance
}

// hammingDistanceByte calculates the Hamming Distance between two bytes
func hammingDistanceByte(b1, b2 byte) int {
	// XOR to find bits which are different
	diff := b1 ^ b2
	// Count the number of set bits
	var count int32
	for diff != 0 {
		count += int32(diff) & 1
		diff >>= 1
	}
	return int(count)
}
