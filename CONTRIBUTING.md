# Contributing to osctx

All contributions are welcome — bug fixes, features, documentation, tests.

## Prerequisites

| Tool | Purpose |
|------|---------|
| Go 1.26+ | Build and test |
| [fzf](https://github.com/junegunn/fzf) | Manual testing of interactive selection |
| [goreleaser](https://goreleaser.com/install/) | Validate build config |

```sh
brew install go fzf goreleaser
```

## Development setup

```sh
git clone https://github.com/yoyrandao/osctx.git
cd osctx
go mod download
```

Add the shell wrapper to your session for manual testing:

```sh
osctx() { eval "$(command go run . "$@")"; }
```

## Running checks

```sh
go build ./...            # build
go test ./...             # tests
goreleaser check          # validate .goreleaser.yml
```

## Submitting changes

1. Fork the repo and create a branch from `main`.
2. Make your changes. Add or update tests where relevant.
3. Ensure all checks above pass.
4. Open a pull request against `main`.

CI runs build, tests, and lint automatically on every PR.

## stdout vs stderr contract

The one non-obvious invariant: **only** the shell export statement goes to stdout; everything else (prompts, lists, error messages) goes to stderr. This is required because the shell wrapper sources stdout:

```sh
osctx() { eval "$(command osctx "$@")"; }
```

Breaking this contract will silently cause the shell function to execute garbage.

## Releases

Releases are published automatically by GoReleaser when a `v*` tag is pushed. Only maintainers push tags.

```sh
git tag v1.2.3
git push --tags
```

GoReleaser builds cross-platform binaries (Linux, macOS, Windows — amd64/arm64), creates archives, and publishes them as a GitHub release.
