# Configuration & Properties

The gcloud CLI configuration system: properties, named configurations, environment variables, and precedence.

> **Source of truth**: [gcloud properties docs](https://cloud.google.com/sdk/docs/properties) and [gcloud configurations docs](https://cloud.google.com/sdk/docs/configurations).

## Properties

Properties are key-value pairs organized in **sections** that govern gcloud CLI behavior. They are stored in named configurations.

### Setting Properties

```bash
gcloud config set project PROJECT_ID
gcloud config set compute/zone us-east1-b
gcloud config set compute/region us-central1
gcloud config set disable_prompts true
gcloud config set accessibility/screen_reader true
```

### Getting Properties

```bash
gcloud config get project
gcloud config get compute/zone
```

### Listing Properties

```bash
gcloud config list                          # List all set properties
gcloud config list --all                    # Include unset properties
gcloud config list --format='text(core.project)'
```

### Unsetting Properties

```bash
gcloud config unset compute/zone
gcloud config unset disable_usage_reporting
```

## Key Property Sections

| Section | Properties | Description |
|---------|-----------|-------------|
| `core` | `account`, `project`, `disable_usage_reporting`, `verbosity`, `disable_prompts`, `default_format`, `format`, `log_http`, `trace_token`, `user_output_enabled`, `show_structured_logs` | Core CLI behavior |
| `compute` | `zone`, `region` | Default compute region/zone |
| `auth` | `access_token_file`, `impersonate_service_account` | Authentication settings |
| `billing` | `quota_project` | Billing/quota project |
| `accessibility` | `screen_reader` | Accessibility settings |
| `metrics` | `command_name` | Metrics settings |

### Common Properties Reference

| Property | Type | Default | Description |
|----------|------|---------|-------------|
| `core/account` | String | — | Active user account email |
| `core/project` | String | — | Active project ID |
| `core/verbosity` | Enum | `warning` | Log verbosity: `debug`, `info`, `warning`, `error`, `critical`, `none` |
| `core/disable_prompts` | Boolean | `false` | Disable interactive prompts |
| `core/format` | String | — | Output format override |
| `core/default_format` | String | — | Default output format (lower priority than `core/format`) |
| `core/log_http` | Boolean | `false` | Log HTTP requests/responses to stderr |
| `core/user_output_enabled` | Boolean | `true` | Print user-intended output |
| `compute/zone` | String | — | Default compute zone (e.g., `us-central1-a`) |
| `compute/region` | String | — | Default compute region (e.g., `us-central1`) |
| `auth/impersonate_service_account` | String | — | Service account to impersonate |
| `auth/access_token_file` | String | — | File path for access token |
| `billing/quota_project` | String | — | Project for quota billing |

## Environment Variables

Every property can also be set via an environment variable using the pattern `CLOUDSDK_SECTION_PROPERTY`:

```bash
export CLOUDSDK_CORE_PROJECT=my-project
export CLOUDSDK_COMPUTE_ZONE=us-central1-a
export CLOUDSDK_COMPUTE_REGION=us-central1
export CLOUDSDK_CORE_DISABLE_PROMPTS=1
export CLOUDSDK_CORE_VERBOSITY=debug
export CLOUDSDK_AUTH_IMPERSONATE_SERVICE_ACCOUNT=sa@project.iam.gserviceaccount.com
```

### Special Environment Variables

| Variable | Purpose |
|----------|---------|
| `CLOUDSDK_ACTIVE_CONFIG_NAME` | Active configuration name (same as `--configuration`) |
| `CLOUDSDK_CONFIG` | Override the config directory (default: `~/.config/gcloud/`) |
| `CLOUDSDK_CORE_PROJECT` | Override `core/project` |
| `CLOUDSDK_CORE_ACCOUNT` | Override `core/account` |
| `CLOUDSDK_CORE_DISABLE_PROMPTS` | Override `core/disable_prompts` |
| `CLOUDSDK_COMPUTE_ZONE` | Override `compute/zone` |
| `CLOUDSDK_COMPUTE_REGION` | Override `compute/region` |

## Precedence

The full precedence order (highest to lowest):

1. **Command-line flags** — `--project=my-project`
2. **Environment variables** — `CLOUDSDK_CORE_PROJECT=my-project`
3. **Configuration properties** — `core/project` in the active configuration
4. **Defaults** — built-in defaults

This means:
- A `--project` flag always wins
- An env var overrides a config property
- Config properties persist across sessions; flags and env vars are per-invocation

## Named Configurations (Profiles)

A **configuration** is a named set of properties, acting like a profile. The initial configuration is named `default`.

### Configuration Commands

| Command | Purpose |
|---------|---------|
| `gcloud config configurations create NAME` | Create a new named configuration |
| `gcloud config configurations activate NAME` | Switch active configuration |
| `gcloud config configurations list` | List all configurations |
| `gcloud config configurations describe NAME` | View properties in a configuration |
| `gcloud config configurations delete NAME` | Delete a configuration (cannot delete active one) |
| `gcloud config configurations rename NAME` | Rename a configuration |

### Using a Configuration for a Single Command

```bash
gcloud auth list --configuration=my-config
```

Or via environment variable:

```bash
CLOUDSDK_ACTIVE_CONFIG_NAME=my-config gcloud auth list
```

### Creating a Configuration and Switching

```bash
# Create a new configuration
gcloud config configurations create work-project

# Activate it
gcloud config configurations activate work-project

# Set properties in the new configuration
gcloud config set project my-work-project
gcloud config set account my-work-email@gmail.com

# Switch back to default
gcloud config configurations activate default
```

### Storage Location

Configurations are stored in `~/.config/gcloud/` (macOS/Linux) or `%APPDATA%\gcloud` (Windows).

Override the config directory:

```bash
export CLOUDSDK_CONFIG=/path/to/custom/config/dir
```

Find the current config directory:

```bash
gcloud info --format='value(config.paths.global_config_dir)'
```

## Config Helper

The `gcloud config config-helper` command (hidden) provides auth and config data to external tools:

```bash
gcloud config config-helper --format=json
```

This outputs all configuration data, credentials, and project information in a machine-readable format. Useful for tooling that needs to consume gcloud's configuration state.

## Installation-Wide Properties

By default, `gcloud config set` updates the active configuration. To set a property installation-wide:

```bash
gcloud config set --installation property value
```

This affects all configurations, not just the active one.

## References

- [gcloud properties](https://cloud.google.com/sdk/docs/properties)
- [gcloud configurations](https://cloud.google.com/sdk/docs/configurations)
- `gcloud topic configurations`

## Related Skills

- For global flags that override these properties, see [global-flags.md](global-flags.md).
- For authentication configuration, see [auth.md](auth.md).
- For project/resource configuration, see [resource-hierarchy.md](resource-hierarchy.md).
