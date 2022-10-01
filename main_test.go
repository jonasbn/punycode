package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArguments(T *testing.T) {
	// We manipuate the Args to set them up for the testcases
	// After this test we restore the initial args
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	cases := []struct {
		Name           string
		Args           []string
		ExpectedExit   int
		ExpectedOutput string
	}{
		{"no arguments", []string{""}, 1, ""},
		{"single ASCII argument", []string{"test"}, 0, "test\n"},
		{"single punycode encoded and lower cased argument", []string{"xn--kdplg-orai3l"}, 0, "kødpålæg\n"},
		{"single punycode encoded and lower cased argument", []string{"kødpålæg"}, 0, "xn--kdplg-orai3l\n"},
		{"single punycode encoded and lower cased argument", []string{"xn--MASSEDELGGELSESVBEN-5ebm60b"}, 0, "MASSEØDELÆGGELSESVÅBEN\n"},
		{"single punycode encoded and lower cased argument", []string{"MASSEØDELÆGGELSESVÅBEN"}, 0, "xn--MASSEDELGGELSESVBEN-5ebm60b\n"},
		{"single punycode encoded and lower cased argument", []string{"xn-MASSEDELGGELSESVBEN-5ebm60b"}, 0, "xn-MASSEDELGGELSESVBEN-5ebm60b\n"},
		{"multiple punycode encoded and lower and upper cased arguments", []string{"xn--kdplg-orai3l", "xn--BLBRGRD-3pak7p"}, 0, "kødpålæg\n"},
		{"multiple punycode encoded and lower and upper cased arguments", []string{"kødpålæg", "BLÅBÆRGRØD"}, 0, "xn--kdplg-orai3l\n"},
	}

	for _, tc := range cases {
		// we need a value to set Args[0] to cause flag begins parsing at Args[1]
		os.Args = append([]string{tc.Name}, tc.Args...)
		//actualExit := realMain()
		var actualExit = 0

		actualOutput := captureOutput(func() {
			actualExit = realMain()
		})

		if tc.ExpectedExit != actualExit {
			T.Errorf("Wrong exit code for args: %v, expected: %v, got: %v",
				tc.Args, tc.ExpectedExit, actualExit)
		}

		assert.Equal(T, tc.ExpectedOutput, actualOutput)
	}

}

// REF:
// https://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string
// https://stackoverflow.com/questions/26804642/how-to-test-a-functions-output-stdout-stderr-in-unit-tests
func captureOutput(f func()) string {

	originalStdout := os.Stdout // keep backup of the original stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing will not block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = originalStdout // restoring the original stdout
	outputStr := <-outC

	return outputStr
}
