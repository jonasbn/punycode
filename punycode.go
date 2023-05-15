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

	var argString string

	if len(argsWithoutProg) <= 0 {
		var err error

		argString, err = readStdin(os.Stdin)

		if err != nil {
			log.Println(err)
			return 2
		}
	} else {
		argString = os.Args[1] // we only take a single parameter, the string to process
	}

	if argString != "" {

		match, _ := regexp.MatchString("^xn--", argString)

		if match {
			unicodeString := puny.ToUnicode(argString)

			if zeroWidth.HasZeroWidthCharacters(unicodeString) {
				fmt.Println(zeroWidth.RemoveZeroWidthJoiner(unicodeString))
			} else {
				fmt.Println(unicodeString)
			}
			return 0

		} else {
			punycodeString := puny.ToASCII(argString)
			fmt.Println(punycodeString)
			return 0
		}

		return 3
	}

	return 1
}

func readStdin(stdin io.Reader) (string, error) {
	scanner := bufio.NewScanner(stdin)

	var argString string

	for scanner.Scan() {
		argString = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return argString, nil
}
