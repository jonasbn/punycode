package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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
		{"single punycode encoded string argument", []string{"xn--kdplg-orai3l"}, 0, "kødpålæg\n"},
		{"single unencoded string argument", []string{"kødpålæg"}, 0, "xn--kdplg-orai3l\n"},
		{"multiple lower and upper cased punycode encoded string arguments", []string{"xn--kdplg-orai3l", "xn--BLBRGRD-3pak7p"}, 0, "kødpålæg\n"},
		{"multiple lower and upper cased unencoded string arguments", []string{"kødpålæg", "BLÅBÆRGRØD"}, 0, "xn--kdplg-orai3l\n"},
		{"stand alone punycode indicator", []string{"xn--"}, 1, ""},
		{"single punycode encoded zero width (zwj)", []string{"xn--1ug6825plhas9r"}, 0, "🧑🏾‍🎨\n"},
		{"single unencoded zero width string (zwj)", []string{"🧑🏾‍🎨"}, 0, "xn--1ug6825plhas9r\n"},
	}

	for _, tc := range cases {
		// we need a value to set Args[0] to cause flag begins parsing at Args[1]
		os.Args = append([]string{tc.Name}, tc.Args...)
		actualExit := 0

		actualOutput := captureOutput(func() {
			actualExit = realMain()
		})

		if tc.ExpectedExit != actualExit {
			t.Errorf("Wrong exit code for args: %v, expected: %v, got: %v",
				tc.Args, tc.ExpectedExit, actualExit)
		}

		assert.Equal(t, tc.ExpectedOutput, actualOutput, tc.Name)
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
		_, _ = io.Copy(&buf, r)
		outC <- buf.String()
	}()

	_ = w.Close()
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
		Error          error
		ExpectedOutput string
	}{
		{"empty string", "", nil, ""},
		{"single punycode encoded input", "xn--kdplg-orai3l", nil, "kødpålæg"},
		{"single multibyte string input", "kødpålæg", nil, "xn--kdplg-orai3l"},
		{"single ASCII string input", "test", nil, "test"},
		{"single punycode encoded input with zwj", "xn--1ug6825plhas9r", nil, "🧑🏾‍🎨"},
		{"single multibyte string input with zwj", "🧑🏾‍🎨", nil, "xn--1ug6825plhas9r"},
	}

	for _, tc := range cases {

		var stdin bytes.Buffer
		stdin.Write([]byte(tc.Input))

		outputString, err := readStdin(&stdin)

		assert.NoError(t, err)

		assert.Equal(t, tc.ExpectedOutput, outputString, tc.Name)
	}
}
