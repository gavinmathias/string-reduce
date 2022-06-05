//go:build !test
// +build !test

package main

import (
	"flag"
	"github.com/gavinmathias/string-reduce/build"
	"github.com/gavinmathias/string-reduce/operate"
	"github.com/gavinmathias/string-reduce/verify"
	"log"
	"os"
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
		log.Printf("Build Version: %s", build.Version)
		log.Printf("Build Time: %s", build.Time)
		os.Exit(0)
	}
}

func main() {
	if digitsOnly && !verify.IsComprisedOFDigits(inputString) {
		log.Fatalf("Input string %s is not comprised entirely of digits. Exiting", inputString)
	}

	outputString := operate.Reduce(inputString, consecutive)
	log.Printf("Input string %s has been reduced to %s", inputString, outputString)
}
