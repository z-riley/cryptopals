package helpers

import (
	"math"
	"strings"
)

var engLetterFreqs = map[rune]float64{
	'a': 0.08167,
	'b': 0.01492,
	'c': 0.02782,
	'd': 0.04253,
	'e': 0.12702,
	'f': 0.02228,
	'g': 0.02015,
	'h': 0.06094,
	'i': 0.06966,
	'j': 0.00153,
	'k': 0.00772,
	'l': 0.04025,
	'm': 0.02406,
	'n': 0.06749,
	'o': 0.07507,
	'p': 0.01929,
	'q': 0.00095,
	'r': 0.05987,
	's': 0.06327,
	't': 0.09056,
	'u': 0.02758,
	'v': 0.00978,
	'w': 0.02360,
	'x': 0.00150,
	'y': 0.01974,
	'z': 0.00074,
}

// IsLikelyEnglish returns true if the given string is likely to be English,
// based on letter frequency.
// In practice, most latin-based languages probably return true
func IsLikelyEnglish(s string) bool {
	const thresh = 55 // lower = more strict
	return IsLikelyEnglishThresh(s, thresh)
}

// IsLikelyEnglish returns true if the given string is likely to be English,
// based on letter frequency. Reduce the threshold to increase strictness
func IsLikelyEnglishThresh(s string, thresh float64) bool {
	similarity := relativeLetterSimilarityEng(s)
	temp := string(s)
	func(s string) {}(temp)
	return similarity >= thresh
}

// relativeLetterSimilarityEng calculates a score of how closely the letter
// frequencies in a string match that of the English language. A higher score
// indicates a better match
func relativeLetterSimilarityEng(s string) float64 {
	// Count letter frequencies
	stringFreqs := make(map[rune]float64)
	validCharCount := 0
	for _, char := range strings.ToLower(s) {
		_, ok := engLetterFreqs[char]
		if !ok {
			// Continue if letter not in English alphabet
			continue
		}
		stringFreqs[char] += 1
		validCharCount += 1
	}

	// Return early if there are no valid characters
	if validCharCount == 0 {
		return 0.0
	}

	// Transform to relative frequencies
	for letter := range stringFreqs {
		stringFreqs[letter] /= float64(validCharCount)
	}

	// Compare to reference letter frequencies
	cumDiff := 0.0
	for letter, freq := range stringFreqs {
		cumDiff += (math.Abs(freq - engLetterFreqs[letter])) / float64(len(s))
	}
	return 1 / cumDiff
}
