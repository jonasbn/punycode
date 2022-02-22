// punycode is a utility to encode and decode to and from punycode on the command line
package main

import (
	"bufio"
	"fmt"
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

	var argString string

	if len(argsWithoutProg) <= 0 {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			argString = scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			log.Println(err)
			return 2
		}
	} else {
		argString = os.Args[1] // we only take a single parameter, the string to process
	}

	if argString != "" {

		match, _ := regexp.MatchString("^xn--", argString)

		if match {
			unicodeString, err := idna.ToUnicode(argString)
			if err == nil {
				fmt.Printf("%s\n", unicodeString)
				return 0
			}

		} else {
			punycodeString, err := idna.ToASCII(argString)
			if err == nil {
				fmt.Printf("%s\n", punycodeString)
				return 0
			}
		}

		return 3
	}

	return 1
}
