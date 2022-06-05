//go:build !test
// +build !test

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gavinmathias/string-reduce/build"
	"github.com/gavinmathias/string-reduce/operate"
	"github.com/gavinmathias/string-reduce/verify"
)

var inputString string
var consecutive int
var versionFlagPassed, digitsOnly bool

func init() {
	flag.StringVar(&inputString, "input", "112233444321", "The string of digits to reduce")
	flag.IntVar(&consecutive, "consecutive", 3, "The number of consecutive digits to remove from the input string of digits")
	flag.BoolVar(&versionFlagPassed, "version", false, "Print the build version and time")
	flag.BoolVar(&digitsOnly, "digits-only", false, "Only accept digits in the input string")
	flag.Parse()

	if versionFlagPassed {
		fmt.Printf("Build Version: %s\n", build.Version)
		fmt.Printf("Build Time: %s\n", build.Time)
		os.Exit(0)
	}
}

func main() {
	if digitsOnly && !verify.IsComprisedOFDigits(inputString) {
		log.Fatalf("Input string %s is not comprised entirely of digits. Exiting", inputString)
	}

	outputString := operate.Reduce(inputString, consecutive)
	fmt.Printf("Input string %s has been reduced to %s\n", inputString, outputString)
}
