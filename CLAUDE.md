# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`osctx` is a CLI tool (inspired by kubectx/kubens) for interactively switching between OpenStack clouds defined in `clouds.yaml`. Module: `github.com/yoyrandao/osctx`, written in Go 1.26.2.

## Commands

```bash
go build ./...          # Build
go test ./...           # Run all tests
go test ./... -run Name # Run a single test
go vet ./...            # Vet
```

## Architecture

**Core constraint**: A subprocess cannot set environment variables in its parent shell. Therefore, `osctx` outputs `export OS_CLOUD=<name>` to **stdout** and everything else to **stderr**. Users must wrap the binary in a shell function to source the output:
```bash
osctx() { eval "$(command osctx "$@")"; }
```

**Commands (via spf13/cobra)**:
- `osctx` — interactive cloud selection via fzf; falls back to index-based selection if fzf is not installed
- `osctx -` — switch to previous cloud
- `osctx ls` — list all clouds from `clouds.yaml`
- `osctx current` — print current cloud
- `osctx unset` — clear current cloud

**External dependency**: `fzf` binary for interactive fuzzy selection (optional; fallback to numbered list).

**Design doc**: [docs/DESIGN.md](docs/DESIGN.md)
