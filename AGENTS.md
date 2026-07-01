# AGENTS.md

## Project Overview

`carapace-gcloud` is a shell completion provider for the Google Cloud CLI (`gcloud`). It enriches the gcloud completer by combining static YAML command specs (embedded at build time) with dynamic completions bridged from gcloud's own completer at runtime.

## Architecture

There are two separate binaries:

### `carapace-gcloud` (the completer)

- **Entry point**: `cmd/carapace-gcloud/main.go`
- **Root command**: `cmd/carapace-gcloud/cmd/root.go` — defines the `gcloud` root command with all global persistent flags (matching gcloud's own), then dynamically registers one cobra subcommand per service (e.g. `compute`, `sql`, `storage`) from the `gcloud.Services()` map.
- **YAML specs**: `cmd/carapace-gcloud/cmd/gcloud/gcloud.*.yaml` — one file per gcloud service group, embedded via `//go:embed *.yaml`. Each spec is a `carapace-spec` `Command` YAML describing subcommands, flags, and static flag completions.
- **Spec loading**: `cmd/carapace-gcloud/cmd/gcloud/gcloud.go` — `Services()` returns the service→description map; `Get(name)` reads and unmarshals a specific YAML spec.
- **Generated file**: `cmd/carapace-gcloud/cmd/gcloud/gcloud_generated.go` — auto-generated `init()` populating the `services` map. **Do not edit by hand**; regenerate with `go generate` (see below).
- **Bridge**: `cmd/carapace-gcloud/common/bridge.go` — `ActionBridgeGcloudCompleter()` delegates flag and positional completion to gcloud's native completer via `carapace-bridge`. The `PreInvoke` hook in `root.go` replaces non-bool flag actions and all positional actions with this bridge action.
- **Usage suppression**: `rootCmd.SetUsageFunc` returns nil to suppress cobra usage output (this is a completer, not a real CLI).

### `carapace-spec-gcloud` (the spec generator)

- **Entry point**: `cmd/carapace-spec-gcloud/main.go`
- **Root command**: `cmd/carapace-spec-gcloud/cmd/root.go` — reads a JSON file (gcloud's internal command model from `gcloud alpha interactive`), converts it to `carapace-spec` YAML, and writes one file per service group.
- **Command model**: `cmd/carapace-spec-gcloud/cmd/command.go` — defines `Cli` and `Command` structs matching gcloud's JSON schema, with `ToSpecCommand()` converting to `carapace-spec` format. Uses `sentences` tokenizer to extract first sentence of flag descriptions.
- **Flags**: `--target` (output directory), `--stdout` (print to stdout), `--no-doc` (strip documentation). `--target` and `--stdout` are mutually exclusive.

### Data flow

1. **Spec generation**: `gcloud alpha interactive` produces JSON → `carapace-spec-gcloud` converts it to YAML specs in `cmd/gcloud/`
2. **Code generation**: `go generate` scans YAML files and produces `gcloud_generated.go` with the service map
3. **Runtime**: `carapace-gcloud` embeds YAML specs, registers cobra commands from the service map, loads specs lazily in `PreRun`, converts them to cobra commands via `spec.Command.ToCobra()`, and bridges dynamic completions to gcloud's native completer

## Commands

```sh
# Build the completer
go build -o cmd/carapace-gcloud/carapace-gcloud ./cmd/carapace-gcloud

# Build the spec generator
go build -o cmd/carapace-spec-gcloud/carapace-spec-gcloud ./cmd/carapace-spec-gcloud

# Regenerate gcloud_generated.go from YAML specs (run from repo root)
go generate ./cmd/carapace-gcloud

# Run tests
go test ./...

# Docker: build image that runs `gcloud alpha interactive` to produce the JSON command model
docker compose build
docker compose run --rm gcloud
```

## The `go generate` pipeline

`main.go` has `//go:generate sh -c "go run -C ./generate ."`. This runs `cmd/carapace-gcloud/generate/main.go`, which:
1. Reads all `gcloud.*.yaml` files in `cmd/gcloud/`
2. Extracts `name` and `description` from each
3. Writes `gcloud_generated.go` with an `init()` that populates the `services` map
4. Runs `go fmt` on the generated file

## Key Dependencies

- **`carapace`** — shell completion engine; provides `carapace.Gen()`, `ActionCallback`, `ActionMap`, etc.
- **`carapace-spec`** — YAML-driven command specs; `spec.Command` unmarshals YAML and converts to cobra commands via `ToCobra()`. Also provides `spec.Register(rootCmd)` for spec-based invocation.
- **`carapace-bridge`** — bridges completions from other completers; `bridge.ActionGcloud()` invokes gcloud's own completion.
- **`pflag` (replaced)** — uses `carapace-sh/carapace-pflag` (a fork) via `replace` directive in `go.mod`.

## Gotchas

- **`gcloud_generated.go` is generated**: Always run `go generate ./cmd/carapace-gcloud` after adding, removing, or renaming YAML spec files. Never hand-edit it.
- **YAML spec files are embedded**: The `//go:embed *.yaml` directive in `gcloud.go` embeds all `.yaml` files in the `gcloud/` directory at build time. Adding/removing YAML files requires rebuilding.
- **Specs are loaded lazily**: Service specs are only unmarshaled in `PreRun` (when the user invokes that service), not at startup. This keeps startup fast given the large number of services.
- **Bridge arg passthrough**: `bridge.go` uses `os.Args[4:]` to reconstruct the completion context — this is fragile and depends on the caller's argument structure.
- **`root.go` has incomplete TODOs**: Flag completion, positional completion, and deeper subcommand bridging are partially implemented (see comments in `root.go`).
- **`carapace-spec-gcloud` uses temp dirs with wrong prefix**: `os.MkdirTemp("", "carapace-spec-botocore-*")` — the temp dir prefix says "botocore" (likely copy-pasted from another project) but this is a gcloud project.
- **No test files**: The project currently has no tests.
- **Go version**: `go.mod` specifies `go 1.25.1`.
- **Docker/compose**: The `Dockerfile` and `compose.yaml` are for generating the gcloud command model JSON, not for running the completer.
