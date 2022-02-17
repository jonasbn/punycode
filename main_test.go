package main

import (
	"os"
	"testing"
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
		{"multiple punycode encoded and lower and upper cased arguments", []string{"xn--kdplg-orai3l", "xn--BLBRGRD-3pak7p"}, 0, "kødpålæg BLÅBÆRGRØD\n"},
		{"multiple punycode encoded and lower and upper cased arguments", []string{"kødpålæg", "BLÅBÆRGRØD"}, 0, "xn--kdplg-orai3l BLÅBÆRGRØD\n"},
	}

	for _, tc := range cases {
		// we need a value to set Args[0] to cause flag begins parsing at Args[1]
		os.Args = append([]string{tc.Name}, tc.Args...)
		actualExit := realMain()
		if tc.ExpectedExit != actualExit {
			T.Errorf("Wrong exit code for args: %v, expected: %v, got: %v",
				tc.Args, tc.ExpectedExit, actualExit)
		}
	}

}
