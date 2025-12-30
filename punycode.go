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

	match, err := regexp.MatchString("^xn--", inputString)

	if err != nil {
		log.Println(err)
		return ""
	}

	var p *idna.Profile = idna.New()

	/* DEBUG OUTPUT
	fmt.Printf("Bytes: %v\n", []byte(inputString))
	fmt.Printf("Runes: %U\n", []rune(inputString))
	fmt.Printf("Length in bytes: %d, runes: %d\n", len(inputString), len([]rune(inputString)))
	fmt.Printf(inputString + "\n")
	*/

	if match {
		unicodeString, _ := p.ToUnicode(inputString)

		outputString = unicodeString

	} else {
		outputString, err = p.ToASCII(inputString)

		if err != nil {
			log.Println(err)
			return ""
		}
	}

	/* DEBUG OUTPUT
	fmt.Printf("Bytes: %v\n", []byte(outputString))
	fmt.Printf("Runes: %U\n", []rune(outputString))
	fmt.Printf("Length in bytes: %d, runes: %d\n", len(outputString), len([]rune(outputString)))
	fmt.Printf(outputString + "\n")
	*/

	return outputString
}
