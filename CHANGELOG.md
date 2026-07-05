# Change log for punycode

## 2026-07-05 0.18.0 Maintenance release

- CI: Added dependabot groups (major-updates + minor-and-patch for github-actions; minor-and-patch for gomod) and 7-day cooldown
- CI: Switched workflows to `go-version-file: go.mod`; pinned to stable Go
- CI: Updated GitHub Actions: actions/checkout to 7.0.0, actions/setup-go to 6.5.0, github/codeql-action to v4.36.3, goreleaser/goreleaser-action to 7.2.3, spellcheck to 0.63.0
- Dependency: `golang.org/x/net` bumped from 0.51.0 to 0.56.0
- Added CLAUDE.md with codebase guidance for Claude Code
- Added `docs/TODO.md` with dependabot improvement backlog
- Fixed: pre-commit lint and format issues in `main_test.go` and `punycode.go`
- Removed CodeSee integration (service discontinued)

## 2026-03-02 0.17.0 Maintenance release

- Dependency: `golang.org/x/net` bumped to 0.51.0
- Go version updated to 1.25.7
- GitHub Actions pinned to commit SHAs for supply-chain security
- CI: Updated github/codeql-action to v4.32.5, actions/setup-go to 6.3.0, goreleaser/goreleaser-action to 7.0.0, spellcheck to 0.59.0, actions/checkout to 6.0.2, goveralls to 1.11.0

## 2025-12-30 0.16.0 Maintenance release

- Updated to outdated automated release process configuration

## 2025-12-30 0.15.0 Maintenance release

### Breaking Changes

- **Zero-width character handling**: The library now uses `golang.org/x/net/idna` instead of the previous `puny` library. As a result, zero-width joiners (ZWJ) in emoji sequences are now preserved during punycode decoding, which is the correct behavior according to Unicode and IDNA2008 standards. Previously, these characters were stripped using the `go-zero-width` library. Users relying on the previous behavior of removing zero-width characters will need to handle this themselves if needed. For example, decoding `xn--1ug6825plhas9r` now returns `🧑🏾‍🎨` (with ZWJ preserved) instead of the emoji without the ZWJ. See PR [#100](https://github.com/jonasbn/punycode/pull/100)

### Changes

- Migrated from `gitlab.com/golang-commonmark/puny` to `golang.org/x/net/idna` for IDNA operations
- Removed dependency on `github.com/trubitsyn/go-zero-width`
- Removed debug output sections from code and tests
- Improved code organization with package-level variable initialization

## 2023-12-04 0.14.0 Maintenance release

- Patch from @dependabot bumping dependency [x/net](https://github.com/golang/net) from version 0.18.0 to 0.19.0, see PR [#56](https://github.com/jonasbn/punycode/pull/56)

## 2023-11-13 0.13.0 Maintenance release

- Patch from @dependabot bumping dependency [x/net](https://github.com/golang/net) from version 0.17.0 to 0.18.0, see PR [#54](https://github.com/jonasbn/punycode/pull/54)

## 2023-10-17 0.12.0 Maintenance release

- Patch from @dependabot bumping dependency [x/net](https://github.com/golang/net) from version 0.16.0 to 0.17.0, see PR [#53](https://github.com/jonasbn/punycode/pull/53)

- Patch from @dependabot bumping dependency [x/net](https://github.com/golang/net) from version 0.15.0 to 0.16.0, see PR [#52](https://github.com/jonasbn/punycode/pull/52)

## 2023-09-12 0.11.0 Maintenance release

- Patch from @dependabot bumping dependency [x/net](https://github.com/golang/net) from version 0.14.0 to 0.15.0, see PR [#48](https://github.com/jonasbn/punycode/pull/48)

## 2023-08-09 0.10.0 Maintenance release

- Patch from @dependabot bumping dependency [x/net](https://github.com/golang/net) from version 0.12.0 to 0.14.0, see PR [#46](https://github.com/jonasbn/punycode/pull/46)

## 2023-07-10 0.9.0 Maintenance release

- Updated goreleaser configuration, I had missed that some options were deprecated

## 2023-07-10 0.8.0 Maintenance release

- Patch from @dependabot bumping dependency [x/net](https://github.com/golang/net) from version 0.11.0 to 0.12.0, see PR [#44](https://github.com/jonasbn/punycode/pull/44)

## 2023-06-19 0.7.0 Maintenance release

- Patch from @dependabot bumping dependency [x/net](https://github.com/golang/net) from version 0.10.0 to 0.11.0, see PR [#42](https://github.com/jonasbn/punycode/pull/42)

## 2023-06-16 0.6.0 Maintenance release

- Patch from @dependabot bumping dependency [stretchr/testify](https://github.com/stretchr/testify) from version 1.8.3 to 1.8.4, see PR [#41](https://github.com/jonasbn/punycode/pull/41)

## 2023-05-23 0.5.0 Maintenance release

- Patch from @dependabot bumping dependency [stretchr/testify](https://github.com/stretchr/testify) from version 1.8.2 to 1.8.3, see PR [#40](https://github.com/jonasbn/punycode/pull/40)

## 2023-05-15 0.4.0 Maintenance release

- Patch from @dependabot bumping dependency [x/net](https://github.com/x/net) from version 0.9.0 to 0.10.0, see PR [#38](https://github.com/jonasbn/punycode/pull/38)

## 2023-01-16 0.3.0 Maintenance release

- Minor changes to the code to improve the test coverage. Test suite extended to accommodate testing of STDIN scenarios

## 2023-01-02 0.2.0 Maintenance release

- Patch from @dependabot bumping dependency [stretchr/testify](https://github.com/stretchr/testify) from version 1.8.0 to 1.8.1, see PR [#28](https://github.com/jonasbn/punycode/pull/28)

## 2022-10-09 0.1.1 Bug fix release

- Patch by @isviridov, correcting wrongly types package name, see PR [#24](https://github.com/jonasbn/punycode/pull/24)

## 2022-02-20 0.1.0 Initial feature release

- Initial working version
- A merge of jonasbn/punydecode and jonasbn/punyencode, both archived now
