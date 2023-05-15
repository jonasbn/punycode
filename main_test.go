package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	zeroWidth "github.com/trubitsyn/go-zero-width"
)

func TestArguments(t *testing.T) {
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
		{"single ASCII string argument", []string{"test"}, 0, "test\n"},
		{"single punycode encoded and lower cased string argument", []string{"xn--kdplg-orai3l"}, 0, "kødpålæg\n"},
		{"single unencoded string and lower cased string argument", []string{"kødpålæg"}, 0, "xn--kdplg-orai3l\n"},
		{"single punycode encoded and lower cased string argument", []string{"xn--MASSEDELGGELSESVBEN-5ebm60b"}, 0, "MASSEØDELÆGGELSESVÅBEN\n"},
		{"single unencoded string and upper cased string argument", []string{"MASSEØDELÆGGELSESVÅBEN"}, 0, "xn--MASSEDELGGELSESVBEN-5ebm60b\n"},
		{"single punycode encoded and upper cased string argument", []string{"xn-MASSEDELGGELSESVBEN-5ebm60b"}, 0, "xn-MASSEDELGGELSESVBEN-5ebm60b\n"},
		{"multiple lower and upper cased punycode encoded string arguments", []string{"xn--kdplg-orai3l", "xn--BLBRGRD-3pak7p"}, 0, "kødpålæg\n"},
		{"multiple lower and upper cased unencoded string arguments", []string{"kødpålæg", "BLÅBÆRGRØD"}, 0, "xn--kdplg-orai3l\n"},
		{"stand alone punycode indicator", []string{"xn--"}, 0, "\n"},
		{"challenge with zwj", []string{"xn--8k8hlfr9n"}, 0, zeroWidth.RemoveZeroWidthCharacters("🧑🏾‍🎨\n")},
		{"challenge with zwj", []string{"🧑🏾‍🎨"}, 0, "xn--8k8hlfr9n\n"},
	}

	for _, tc := range cases {
		// we need a value to set Args[0] to cause flag begins parsing at Args[1]
		os.Args = append([]string{tc.Name}, tc.Args...)
		// actualExit := realMain()
		var actualExit = 0

		actualOutput := captureOutput(func() {
			actualExit = realMain()
		})

		if tc.ExpectedExit != actualExit {
			t.Errorf("Wrong exit code for args: %v, expected: %v, got: %v",
				tc.Args, tc.ExpectedExit, actualExit)
		}

		assert.Equal(t, tc.ExpectedOutput, actualOutput)
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

// REF:
// https://petersouter.xyz/testing-and-mocking-stdin-in-golang/
func TestStdin(t *testing.T) {

	cases := []struct {
		Name           string
		Input          string
		Err            error
		ExpectedOutput string
	}{
		{"empty string", "", nil, ""},
		{"single punycode encoded and lower cased input", "xn--kdplg-orai3l", nil, "xn--kdplg-orai3l"},
		{"single multibyte string and lower cased input", "kødoplæg", nil, "kødoplæg"},
		{"single punycode encoded and upper cased input", "xn-MASSEDELGGELSESVBEN-5ebm60b", nil, "xn-MASSEDELGGELSESVBEN-5ebm60b"},
		{"single multibyte string and upper cased input", "MASSEØDELÆGGELSESVÅBEN", nil, "MASSEØDELÆGGELSESVÅBEN"},
		{"single ASCII string input", "test", nil, "test"},
		{"challenge with zwj", "xn--8k8hlfr9n", nil, "🧑🏾‍🎨"},
	}

	for _, tc := range cases {

		var stdin bytes.Buffer
		stdin.Write([]byte(tc.Input))

		result, err := readStdin(&stdin)

		assert.NoError(t, err)
		assert.Equal(t, tc.ExpectedOutput, result)
	}
}
