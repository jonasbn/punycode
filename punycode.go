// punycode is a utility to encode and decode to and from punycode on the command line
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	zeroWidth "github.com/trubitsyn/go-zero-width"
	"gitlab.com/golang-commonmark/puny"
)

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

	match, _ := regexp.MatchString("^xn--", inputString)

	if match {
		unicodeString := puny.ToUnicode(inputString)

		if zeroWidth.HasZeroWidthCharacters(unicodeString) {
			outputString = zeroWidth.RemoveZeroWidthJoiner(unicodeString)
		} else {
			outputString = unicodeString
		}
	} else {
		outputString = puny.ToASCII(inputString)
	}

	return outputString
}
