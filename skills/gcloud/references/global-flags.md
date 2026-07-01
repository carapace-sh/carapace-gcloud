# Global Flags

The flags available on every `gcloud` command invocation, their types, behavior, and property overrides.

> **Source of truth**: `gcloud --help` and [gcloud reference](https://cloud.google.com/sdk/gcloud/reference).

## Flag Precedence

Flags override properties. The precedence order is:

1. **Command-line flags** (highest)
2. **Environment variables** (`CLOUDSDK_SECTION_PROPERTY`)
3. **Configuration file properties** (lowest)

For example, `--project=my-project` overrides `CLOUDSDK_CORE_PROJECT`, which overrides the `core/project` config property.

## Primary Global Flags

| Flag | Type | Default | Property Override | Description |
|------|------|---------|-------------------|-------------|
| `--account` | String (email) | — | `core/account` | Google Cloud user account to use for invocation |
| `--billing-project` | String (project ID) | — | `billing/quota_project` | Project charged quota for API operations |
| `--configuration` | String (name) | — | `CLOUDSDK_ACTIVE_CONFIG_NAME` env var | Named configuration to use for this invocation |
| `--flags-file` | File path (YAML/JSON) | — | — | File containing `--flag:value` dictionary |
| `--flatten` | List of keys | — | — | Flatten `name[]` resource slices into separate records |
| `--format` | String (enum) | command-specific | `core/format`, `core/default_format` | Output format for command results |
| `--help` | Boolean | — | — | Display detailed help |
| `--project` | String (project ID) | current project | `core/project` | Google Cloud project ID for invocation |
| `--quiet` / `-q` | Boolean | — | `core/disable_prompts` | Disable all interactive prompts |
| `--verbosity` | String (enum) | `warning` | `core/verbosity` | Log verbosity level |
| `--version` / `-v` | Boolean | — | — | Print version info (global level only) |
| `-h` | Boolean | — | — | Print summary help |

## Authentication & Authorization Flags

| Flag | Type | Property Override | Description |
|------|------|-------------------|-------------|
| `--access-token-file` | File path | `auth/access_token_file` | File containing a raw access token; ignores active account credentials |
| `--impersonate-service-account` | String (email or comma-separated chain) | `auth/impersonate_service_account` | Make API requests as the given service account(s) via token impersonation |
| `--log-http` | Boolean | `core/log_http` | Log all HTTP server requests and responses to stderr |
| `--trace-token` | String | `core/trace_token` | Token used to route traces of service requests |

## Output Control Flags

| Flag | Type | Property Override | Description |
|------|------|-------------------|-------------|
| `--user-output-enabled` | Boolean (tri-state) | `core/user_output_enabled` | Print user-intended output to console |
| `--no-user-output-enabled` | Boolean | — | Suppress user-intended output |

## Hidden / Internal Flags

These flags exist in the gcloud CLI but are hidden from help output:

| Flag | Type | Description |
|------|------|-------------|
| `--authority-selector` | String | Internal — needs help text |
| `--authorization-token-file` | String | Internal — needs help text |
| `--credential-file-override` | String | Internal — needs help text |
| `--document` | String | Internal — should be hidden |
| `--force-endpoint-mode` | String | Regional endpoint mode: `legacy`, `auto`, `regional-only` |
| `--http-timeout` | String | Internal — needs help text |
| `--no-log-http` | Boolean | Inverse of `--log-http` |
| `--universe-domain` | String | Universe domain to target |

## Verbosity Values

The `--verbosity` flag accepts these values:

| Value | Level |
|-------|-------|
| `debug` | Most verbose — includes debug messages |
| `info` | Informational messages |
| `warning` | Default — warnings and errors |
| `error` | Errors only |
| `critical` | Critical errors only |
| `none` | No logging output |

## Format Values

The `--format` flag accepts these format names. See [output-formatting.md](output-formatting.md) for full details.

| Format | Description |
|--------|-------------|
| `json` | JSON output |
| `yaml` | YAML output (default for most commands) |
| `csv` | Comma-separated values (requires projection) |
| `table` | Aligned columns with headings |
| `value` | Tab-separated values (for scripting) |
| `text` / `flattened` | Key: value pairs per line |
| `list` | Ordered list of items |
| `config` | Config-style dictionary |
| `diff` | Unified diff |
| `multi` | Multiple sub-formats |
| `get` | Value without transforms |
| `object` | Bypass JSON serialization |
| `none` / `disable` | Suppress output |

## Key Flag Behaviors

### `--project` Dual Role

The `--project` flag serves **two roles**:
1. Specifies the project of the resource to operate on
2. Specifies the project for API enablement check, quota, and billing

Use `--billing-project` to decouple quota/billing from the target project.

### `--impersonate-service-account` Delegation Chains

Supports a comma-separated delegation chain: `--impersonate-service-account=sa1,sa2,sa3`

- The active account must have `roles/iam.serviceAccountTokenCreator` on `sa1`
- `sa1` must have the same role on `sa2`
- `sa2` must have the same role on `sa3`

### `--quiet` in Scripts

In scripts, always use `--quiet` or `-q` to avoid hanging on interactive prompts. For destructive commands like `gcloud projects delete`, this skips confirmation. If input is required and there's no default, an error is raised.

Equivalent: `CLOUDSDK_CORE_DISABLE_PROMPTS=1` or `gcloud config set disable_prompts true`.

### `--flags-file` Format

A YAML or JSON file specifying a `--flag: value` dictionary:

```yaml
# flags.yaml
--zone: us-central1-a
--machine-type: e2-medium
--preemptible: true
```

```bash
gcloud compute instances create my-instance --flags-file=flags.yaml
```

Each `--flags-file` argument is replaced by its constituent flags. See `gcloud topic flags-file`.

## References

- [gcloud reference — global flags](https://cloud.google.com/sdk/gcloud/reference)
- `gcloud topic flags-file`

## Related Skills

- For configuration properties that flags override, see [config-properties.md](config-properties.md).
- For output formatting details (`--format`, `--filter`, `--flatten`), see [output-formatting.md](output-formatting.md).
- For authentication flags usage, see [auth.md](auth.md).
