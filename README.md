# punycode

[![Go Report](https://goreportcard.com/badge/github.com/jonasbn/punycode)](https://goreportcard.com/report/github.com/jonasbn/punycode)
[![Coverage Status](https://coveralls.io/repos/github/jonasbn/punycode/badge.svg?branch=main)](https://coveralls.io/github/jonasbn/punycode?branch=main)

A simple CLI tool to decode a punycode encoded string

## Usage

### Decoding a single string, meaning conversion from punycode to human readable text

```bash
punycode xn--blbrgrd-fxak7p
```

Will emit

```text
blåbærgrød
```

As an alternative to provided arguments, you can pipe text into `punycode`

```bash
echo xn--blbrgrd-fxak7p | punycode
```

Will emit

```text
blåbærgrød
```

### Encoding a single string, meaning conversion from human readable text to punycode

```bash
punycode kødpålæg
```

Will emit

```text
xn--kdplg-orai3l
```

As an alternative to provided arguments, you can pipe text into `punycode`

```bash
echo kødpålæg | punycode
```

Will emit

```text
xn--kdplg-orai3l
```

## Installation

Installation is easy using Go

```bash
go install github.com/jonasbn/punycode@latest
```

If you want a particular version, please see [Go Modules Reference][MOD]

## Description

TODO

## Diagnostics

TODO

## Exit Status

- `0` success, provided string has been decoded and printed

- `1` failure no argument provided or data from STDIN

- `2` failure reading from STDIN

- `3` encoding or decoding failed

## Dependencies

This utility requires:

- [golang.org/x/net/idna][goidna]

In addition to a few of the standard libraries

## Bugs and Limitations

There a no known bugs, please see the GitHub repository [issues section](https://github.com/jonasbn/punycode/issues) for a up to date listing.

### Only support for Unicode

The utility is limited to decoding to Unicode (version 13) from Punycode.

Please see [golang.org/x/net/idna][goidna] for details.

### Only a single argument

`punycode` only takes a single argument.

```bash
punycode xn--blbrgrd-fxak7p
blåbærgrød
```

So providing more than one argument will only return the first argument

```bash
punycode xn--blbrgrd-fxak7p xn--kdplg-orai3l
blåbærgrød
```

## Author

- jonasbn

## Acknowledgements

- @isviridov, bug fix contribution

## Motivation

This utility was created, when in the process of learning Go. I have worked in the DNS and domain name business for a decade so it was only natural to work on something I _know_ when learning Go.

This particular repository touched the following topics:

1. Learning to make CLI tools
1. Making an executable distributable and installable component
1. Reading data from the CLI
1. Reading data from STDIN
1. Testing a CLI tool / Main function in Go

All of the above was covered in: [punyencode][punyencode] and [punydecode][punydecode]

The `punycode` is a merge of the two, which then opened up for more things to learn.

1. Using a regular expressions
1. Capturing and testing output emitted to STDOUT

See the resources and references below for resources on the above topics.

## Resources and References

1. [Wikipedia: Punycode](https://en.wikipedia.org/wiki/Punycode)
1. [Go Modules Reference][MOD]
1. [GitHub: punydecode][punydecode]
1. [GitHub: punyencode][punyencode]
1. [golang.org/x/net/idna][goidna]
1. [Go By Example: Regular Expressions](https://gobyexample.com/regular-expressions)
1. [yourbasic.org/golang: Read a file (stdin) line by line](https://yourbasic.org/golang/read-file-line-by-line/)
1. [Blog post: Test the main function in Go by Johannes Malsam](https://mj-go.in/golang/test-the-main-function-in-go)
1. [StackOverflow: "In Go, how do I capture stdout of a function into a string?"](https://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string)
1. [StackOverflow: "How to test a function's output (stdout/stderr) in unit tests"](https://stackoverflow.com/questions/26804642/how-to-test-a-functions-output-stdout-stderr-in-unit-tests)

## License and Copyright

Copyright Jonas Brømsø (jonasbn) 2022-2023

MIT License, see separate `LICENSE` file

[MOD]: https://go.dev/ref/mod#go-install
[punydecode]: https://github.com/jonasbn/punydecode
[punyencode]: https://github.com/jonasbn/punyencode
[goidna]: https://pkg.go.dev/golang.org/x/net/idna
