# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Build
go build ./...

# Test
go test -v ./...

# Run a single test
go test -v -run TestArguments ./...
go test -v -run TestStdin ./...

# Format
gofumpt -w .
goimports -w .

# Lint
golangci-lint run
go vet ./...
```

## Architecture

This is a small, single-package CLI tool (`main` package) with production code in `punycode.go` and tests in `main_test.go`.

**Entry point pattern:** `main()` calls `realMain()` and passes the exit code to `os.Exit()`. This indirection makes the real logic testable — tests call `realMain()` directly rather than `main()`.

**Input handling:** `realMain()` checks whether `os.Args` has arguments (uses first arg via `readArgs()`) or falls back to `readStdin()`, which reads STDIN line-by-line but uses only the final line read. Multiple arguments are silently ignored — only the first is used.

**Conversion logic:** `convertString()` uses a regex (`^xn--`) to detect punycode input. Punycode strings are decoded using `golang.org/x/net/idna`; Unicode strings are encoded to punycode using the same package.

**Exit codes:** 0 = success, 1 = empty output / invalid input / no input, 2 = STDIN read error.

## Testing Approach

Tests use `captureOutput()` (redirects `os.Stdout` to a pipe) to capture what the CLI prints. The test suite covers encoding, decoding, emoji with zero-width joiners (ZWJ sequences), empty input, and STDIN input. The testify library is used for assertions.

## Dependencies

- `golang.org/x/net/idna` — punycode encode/decode (migrated from `gitlab.com/golang-commonmark/puny` in v0.15.0; the new library preserves ZWJ sequences in emoji)
- `github.com/stretchr/testify` — test assertions

## Pre-commit Hooks

The project uses pre-commit hooks (`.pre-commit-config.yaml`). Go-focused hooks run `go-fmt-import`, `gofumpt`, `go-vet`, `go-lint`, `go-unit-tests`, `go-static-check`, and `golangci-lint`; the configuration also includes general `pre-commit-hooks` checks (for example, trailing whitespace and end-of-file fixes) and `markdownlint`. Run `pre-commit install` once after cloning to enable them.

Direct commits to `main` are blocked by the `no-commit-to-branch` hook — work on a feature branch instead.

**Gotchas:**

- `golangci-lint` pre-commit hook is disabled (Bahjat hook exits 1 on "0 issues." — a stdout capture bug). Run manually: `golangci-lint run`
- `staticcheck` must match or exceed the Go toolchain version; update with `go install honnef.co/go/tools/cmd/staticcheck@latest` if it fails with a version mismatch.
- New technical terms in `.md` files must be added to `.github/spellcheck-wordlist.txt` or the spellcheck CI job will fail.

## Releases

Releases are automated via GoReleaser (`.goreleaser.yml`) triggered by `v*` tags. Binaries are built for Linux, Windows, and macOS with `CGO_ENABLED=0`. The release workflow uses the built-in `GITHUB_TOKEN`; ensure workflow permissions allow creating releases.

Branch protection requires 1 approving review. All PRs are created under the repo owner's profile and cannot be self-approved — use `gh pr merge <N> --squash --delete-branch --admin` to merge.

## Backlog

`docs/TODO.md` tracks future improvements (dependabot tuning, tooling, CI). Check it before starting maintenance work.
