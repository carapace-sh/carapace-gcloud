# Gotchas & Edge Cases

Common pitfalls, non-obvious behavior, and edge cases in the gcloud CLI.

## Authentication

### `gcloud auth login` Does NOT Set Up ADC

`gcloud auth login` provides credentials for **gcloud CLI only**. Client libraries (Terraform, Python `google-auth`, Node.js, Go, etc.) use **Application Default Credentials (ADC)**, which is a separate system.

- `gcloud auth login` → stores in `credentials.db` (SQLite)
- `gcloud auth application-default login` → stores in `application_default_credentials.json`

If client libraries fail with credential errors despite being logged in via `gcloud auth login`, you need to run `gcloud auth application-default login` separately.

See [auth.md](auth.md) for details.

### Credentials Stored in Two Places

| Credential Type | Storage Location |
|----------------|-----------------|
| gcloud CLI credentials | `~/.config/gcloud/credentials.db` (SQLite) |
| ADC (for client libraries) | `~/.config/gcloud/application_default_credentials.json` (JSON) |

## Project Identity

### Project ID vs Project Name vs Project Number

| Identifier | Mutable | Example | Usage |
|-----------|---------|---------|-------|
| **Project ID** | No | `my-project-123` | Used in gcloud commands and flags |
| **Project Number** | No (auto-assigned) | `464036093014` | Can be used interchangeably with ID |
| **Name** | Yes | `My Project` | Display only — **not** accepted by commands |

Using the project **name** in commands where the **ID** is expected is a common error.

### `--project` Dual Role

The `--project` flag serves **two roles**:
1. Specifies the project of the resource to operate on
2. Specifies the project for API enablement check, quota, and billing

Use `--billing-project` (or `billing/quota_project` property) to decouple these:

```bash
gcloud compute instances list --project=my-resource-project --billing-project=my-quota-project
```

## stdout vs stderr

- **stdout**: Successful command output — safe for scripting and piping
- **stderr**: Prompts, warnings, errors — **do not script against these** (wording can change between versions)

The gcloud team explicitly warns: *"Do not script against responses written to stderr because these responses aren't stable."*

## Filter Gotchas

### `:` Operator Behavior Change

The `:` operator is changing from substring matching to word matching. A warning is displayed when both implementations would return different results. Avoid relying on substring matching behavior.

### `=` Operator Inconsistency

The `=` operator behaves inconsistently across APIs — for some it means strict equality, for others it acts like `:` (contains). This is being phased out.

### Shell Quoting

Most filter expressions need shell quoting. The interaction between shell quotes and filter string literals can be confusing:

```bash
# Correct: single quotes for shell, double quotes inside for filter literals
--filter='zone="(us-central1-a us-central1-b)"'

# Or vice versa
--filter="zone='us-central1-a'"
```

### Scientific Notation

Values like `30e5504145` are interpreted as numbers (`30 × 10^5504145`). If you need to filter on a value that looks like scientific notation, quote it:

```bash
--filter="'tags:30e5504145'"
```

## Processing Order

The `--flatten`, `--sort-by`, `--filter`, and `--limit` flags are applied in a **fixed order**:

```
--flatten → --sort-by → --filter → --limit
```

This means:
- You can filter on flattened fields (filter runs after flatten)
- You can sort before filtering (sort runs before filter)
- `--limit` applies after all other processing

See [output-formatting.md](output-formatting.md) for details.

## Quiet Mode & Destructive Commands

- Commands like `gcloud projects delete` are **interactive by default** and require confirmation
- In scripts, always use `--quiet` / `-q` (or `CLOUDSDK_CORE_DISABLE_PROMPTS=1`)
- With `--quiet`, if input is required and there's no default, an **error is raised** (the command doesn't hang)

## Component Manager Disabled with Package Managers

If installed via APT (`google-cloud-cli`) or YUM/DNF, `gcloud components install` **will not work**. The component manager is disabled.

Instead, use the system package manager:

```bash
# APT
sudo apt-get install google-cloud-cli-kubectl
sudo apt-get install google-cloud-cli-pubsub-emulator

# YUM/DNF
sudo dnf install google-cloud-cli-kubectl
```

This is one of the most commonly encountered issues. See [release-tracks.md](release-tracks.md).

## Cloud Shell Configuration Persistence

In Google Cloud Shell, gcloud configurations are stored in a **temporary directory** and **do not persist** across sessions. Each new Cloud Shell session starts with a fresh default configuration.

To persist configuration, set properties via environment variables in `~/.bashrc` or use `gcloud config set` at the start of each session.

## Configuration Storage Location

- Default: `~/.config/gcloud/` (macOS/Linux), `%APPDATA%\gcloud` (Windows)
- Override with `CLOUDSDK_CONFIG` environment variable
- Find current location: `gcloud info --format='value(config.paths.global_config_dir)'`

## Python Version Requirements

- Requires Python **3.10+** (check `gcloud topic startup` for current supported range)
- The x86_64 Linux package includes a **bundled Python interpreter** preferred by default
- Check: `python3 -V` or `gcloud topic startup`

## Impersonation Delegation Chain

The delegation chain requires each service account to have `roles/iam.serviceAccountTokenCreator` on the next:

```
--impersonate-service-account=sa1,sa2,sa3
```

- Active account → token creator on `sa1`
- `sa1` → token creator on `sa2`
- `sa2` → token creator on `sa3`

If any link in the chain is missing the role, impersonation fails with a permission error.

## `gcloud init` with Many Projects

When running `gcloud init` with access to 200+ projects, you'll be prompted to enter a project ID, create one, or list projects. The list may be paginated. For scripting, set the project directly:

```bash
gcloud config set project PROJECT_ID
```

## Windows-Specific Issues

- Requires **Windows 8.1+** or **Windows Server 2012+**
- The `find` command must be in `PATH` (usually `C:\WINDOWS\system32`)
- If uninstalling, a **reboot is required** before reinstalling
- If unzipping fails, run the installer as **administrator**

## Deprecated `gcloud docker` Command

`gcloud docker` is deprecated. Use `gcloud auth configure-docker` to register gcloud as a Docker credential helper, then use `docker` directly.

## `--no-` Flag Pattern

Many boolean flags have a `--no-` prefixed inverse (e.g., `--async` / `--no-async`). In the carapace-gcloud YAML specs, these are marked with the `&` suffix (e.g., `--no-async&`). Both forms set the same underlying property — the `--no-` form sets it to `false`.

## References

- [gcloud scripting guide](https://cloud.google.com/sdk/docs/scripting-gcloud)
- [gcloud troubleshooting](https://cloud.google.com/sdk/docs/troubleshooting)

## Related Skills

- For authentication details, see [auth.md](auth.md).
- For filter syntax, see [output-formatting.md](output-formatting.md).
- For component management, see [release-tracks.md](release-tracks.md).
- For configuration, see [config-properties.md](config-properties.md).
