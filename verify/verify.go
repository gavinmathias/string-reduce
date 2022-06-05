package verify

import "unicode"

// IsComprisedOFDigits returns true if all characters in a string "inputString" are digits and false otherwise
func IsComprisedOFDigits(inputString string) bool {
	inputCharacters := []rune(inputString)
	if len(inputCharacters) == 0 {
		return false
	}
	for i := 0; i < len(inputCharacters); i++ {
		rune := inputCharacters[i]
		if !unicode.IsDigit(rune) {
			return false
		}
	}
	return true
}

// ContainsConsecutiveCharacters checks "inputString" for consecutive "consecutive" occurences
// of any character. If found, it returns found equal to true and the index "index" of the first occurrence.
// Otherwise it returns found equal to false and -1 for index.
func ContainsConsecutiveCharacters(inputString string, consecutive int) (found bool, index int) {
	inputCharacters := []rune(inputString)
	characters := make([]string, len(inputCharacters))
	for i := 0; i < len(inputCharacters); i++ {
		characters[i] = string(inputCharacters[i])
	}
	found = false
NoFoundConsecutiveCharacters:
	for i := 0; i < len(inputCharacters); i++ {
		if i+consecutive > len(inputCharacters) {
			break
		}
		for j := 1; j < consecutive; j++ {
			if characters[i] != characters[i+j] {
				continue NoFoundConsecutiveCharacters
			}
		}
		found = true
		return found, i
	}
	found = false
	return found, -1
}
