package verify

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type VerifyTestSuite struct {
	suite.Suite
	IsComprisedOFDigitsTests           []IsComprisedOFDigitsTest
	ContainsConsecutiveCharactersTests []ContainsConsecutiveCharactersTest
}

type IsComprisedOFDigitsTest struct {
	input  string
	output bool
}

type ContainsConsecutiveCharactersTest struct {
	input       string
	consecutive int
	found       bool
	index       int
}

func (suite *VerifyTestSuite) SetupTest() {
	suite.IsComprisedOFDigitsTests = []IsComprisedOFDigitsTest{
		{"112233444321", true},
		{"82892929f993837399", false},
		{"8", true},
		{"", false},
		{"ajisdijd02+++==[]]'", false},
	}

	suite.ContainsConsecutiveCharactersTests = []ContainsConsecutiveCharactersTest{
		{"112233444321", 3, true, 6},
		{"", 3, false, -1},
		{"92782782777777980202", 7, false, -1},
		{"jkdkdjsnsksl;'';//;,...++++ksks9ks", 4, true, 23},
		{"0", 1, true, 0},
	}
}

func (suite *VerifyTestSuite) TestIsComprisedOFDigits() {
	for _, test := range suite.IsComprisedOFDigitsTests {
		output := IsComprisedOFDigits(string(test.input))
		suite.Equal(output, test.output)
	}
}

func (suite *VerifyTestSuite) TestContainsConsecutiveCharacters() {
	for _, test := range suite.ContainsConsecutiveCharactersTests {
		found, index := ContainsConsecutiveCharacters(string(test.input), test.consecutive)
		suite.Equal(index, test.index)
		suite.Equal(found, test.found)
	}
}

func TestVerifyTestSuite(t *testing.T) {
	suite.Run(t, new(VerifyTestSuite))
}
