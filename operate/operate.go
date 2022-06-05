package operate

import "github.com/gavinmathias/string-reduce/verify"

// Reduce takes "inputString" and removes "consecutive" consecutive occurences of any characters
//
// It will recurse until all consecutive "consecutive" characters are removed.
//
// The final reduced string "outputString" is returned.
func Reduce(inputString string, consecutive int) (outputString string) {
	inputCharacters := []rune(inputString)
	var outputCharacters []rune
	outputString = inputString

	found, index := verify.ContainsConsecutiveCharacters(inputString, consecutive)
	if found {

		outputCharacters = removeConsecutiveRunesAtIndex(inputCharacters, index, consecutive)
		outputString = string(outputCharacters)
	} else {
		outputString = inputString
	}
	if outputString != inputString {
		return Reduce(outputString, consecutive)
	} else {
		return outputString
	}
}

// removeRuneAtIndex removes a single rune from a rune array "input" at the index "index"
//
// To prevent panics, we return "input" rune array if the index is out of bounds.
func removeRuneAtIndex(input []rune, index int) []rune {
	if index+1 > len(input) {
		return input
	}
	return append(input[:index], input[index+1:]...)
}

// removeConsecutiveRunesAtIndex starts at "index" of a rune array "input" and removes "consecutive" consecutive characters
//
// It returns the resulting rune array. If "input" is empty, it returns it unchanged
func removeConsecutiveRunesAtIndex(input []rune, index int, consecutive int) []rune {
	if len(input) == 0 {
		return input
	}
	if consecutive == 1 {
		return removeRuneAtIndex(input, index)
	} else {
		input = removeRuneAtIndex(input, index)
		return removeConsecutiveRunesAtIndex(input, index, consecutive-1)
	}
}
