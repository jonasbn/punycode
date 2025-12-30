// punycode is a utility to encode and decode to and from punycode on the command line
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"golang.org/x/net/idna"
)

// punycodePrefix is a compiled regex pattern to match the punycode prefix "xn--"
var punycodePrefix = regexp.MustCompile("^xn--")

// profile is the IDNA profile used for punycode conversions, initialized once at package level.
// Using idna.New() creates a profile with default options suitable for general-purpose
// bidirectional conversion between Unicode and ASCII (punycode) representations.
var profile = idna.New()

// main function is a wrapper on the realMain function and emits OS exit code based on wrapped function
func main() {
	os.Exit(realMain())
}

// realMain function wrapped so it is testable, can read arguments from CLI og STDING returns integer, which can be translated into OS exit code
func realMain() int {
	argsWithoutProg := os.Args[1:]

	var outputString string

	if len(argsWithoutProg) <= 0 {
		var err error

		outputString, err = readStdin(os.Stdin)

		if err != nil {
			log.Println(err)
			return 2
		}

	} else {
		outputString = readArgs()
	}

	if outputString != "" {
		fmt.Printf("%s\n", outputString)
		return 0
	} else {
		return 1
	}
}

func readArgs() string {
	inputString := os.Args[1] // we only take a single parameter, the string to process

	return convertString(inputString)
}

func readStdin(stdin io.Reader) (string, error) {
	scanner := bufio.NewScanner(stdin)

	var inputString string

	for scanner.Scan() {
		inputString = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return convertString(inputString), nil
}

func convertString(inputString string) string {

	var outputString string
	var err error
	var unicodeString string

	match := punycodePrefix.MatchString(inputString)

	if match {
		unicodeString, err = profile.ToUnicode(inputString)

		if err != nil {
			log.Println(err)
			return ""
		}
		outputString = unicodeString

	} else {
		outputString, err = profile.ToASCII(inputString)

		if err != nil {
			log.Println(err)
			return ""
		}
	}

	return outputString
}
