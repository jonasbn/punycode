package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	"golang.org/x/net/idna"
)

func main() {
	os.Exit(realMain())
}

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
		argString = os.Args[1] // we only take a single parameter, the string to decode
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
