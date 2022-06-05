package operate

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type OperateTestSuite struct {
	suite.Suite
	RemoveRuneAtIndexTests             []RemoveRuneAtIndexTest
	RemoveConsecutiveRunesAtIndexTests []RemoveConsecutiveRunesAtIndexTest
	ReduceTests                        []ReduceTest
}

type RemoveRuneAtIndexTest struct {
	input  []rune
	output []rune
	index  int
}

type RemoveConsecutiveRunesAtIndexTest struct {
	input       []rune
	output      []rune
	index       int
	consecutive int
}

type ReduceTest struct {
	input       string
	output      string
	consecutive int
}

func (suite *OperateTestSuite) SetupTest() {
	suite.RemoveRuneAtIndexTests = []RemoveRuneAtIndexTest{
		{[]rune("112233444321"), []rune("12233444321"), 0},
		{[]rune("1124321"), []rune("112421"), 4},
		{[]rune("8484383733"), []rune("848433733"), 5},
		{[]rune("3839372022"), []rune("3839372022"), 50},
	}

	suite.RemoveConsecutiveRunesAtIndexTests = []RemoveConsecutiveRunesAtIndexTest{
		{[]rune(""), []rune(""), 5, 13},
		{[]rune("112233444321"), []rune("233444321"), 0, 3},
		{[]rune("1124321"), []rune("11241"), 4, 2},
		{[]rune("8484383733"), []rune("848433"), 5, 4},
		{[]rune("3839372022"), []rune("3839372022"), 50, 23},
	}

	suite.ReduceTests = []ReduceTest{
		{"112233444321", "", 3},
		{"11223344431", "11221", 3},
		{"83992827282902202020", "", 1},
		{"982729202089999888", "9827292020", 4},
	}
}

func (suite *OperateTestSuite) TestRemoveRuneAtIndex() {
	for _, test := range suite.RemoveRuneAtIndexTests {
		output := removeRuneAtIndex(test.input, test.index)
		suite.Equal(string(output), string(test.output))
	}
}

func (suite *OperateTestSuite) TestRemoveConsecutiveRunesAtIndex() {
	for _, test := range suite.RemoveConsecutiveRunesAtIndexTests {
		output := removeConsecutiveRunesAtIndex(test.input, test.index, test.consecutive)
		suite.Equal(string(output), string(test.output))
	}
}

func (suite *OperateTestSuite) TestReduce() {
	for _, test := range suite.ReduceTests {
		output := Reduce(test.input, test.consecutive)
		suite.Equal(output, test.output)
	}
}
func TestOperateTestSuite(t *testing.T) {
	suite.Run(t, new(OperateTestSuite))
}
