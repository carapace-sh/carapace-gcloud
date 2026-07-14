# Authentication

The gcloud CLI authentication system: user accounts, service accounts, Application Default Credentials (ADC), impersonation, and Docker credential helper.

> **Source of truth**: [gcloud auth reference](https://cloud.google.com/sdk/gcloud/reference/auth) and [ADC documentation](https://cloud.google.com/docs/authentication/application-default-credentials).

## User Account Authentication

### `gcloud auth login`

Authenticate as a user with a browser-based OAuth2 flow. Credentials are stored in a SQLite database (`credentials.db`) under `~/.config/gcloud/`.

```bash
gcloud auth login                        # Browser-based auth
gcloud auth login --no-launch-browser    # No browser (device code flow)
gcloud auth login --no-launch-browser    # Alternative no-browser flow
```

These credentials are used **only by gcloud CLI itself**, not by client libraries or SDKs.

### `gcloud auth list`

List all credentialed accounts:

```bash
gcloud auth list
```

The active account is marked with `*`.

### `gcloud auth revoke`

Remove credentials:

```bash
gcloud auth revoke                       # Revoke active account
gcloud auth revoke ACCOUNT_EMAIL         # Revoke specific account
gcloud auth revoke --all                 # Revoke all accounts
```

### `gcloud auth print-access-token`

Print an access token for the current account:

```bash
gcloud auth print-access-token
gcloud auth print-access-token --scopes=https://www.googleapis.com/auth/cloud-platform
```

## Service Account Authentication

### `gcloud auth activate-service-account`

Authorize access to Google Cloud with a service account:

```bash
gcloud auth activate-service-account --key-file=/path/to/key.json
```

For `.p12` files:

```bash
gcloud auth activate-service-account --key-file=/path/to/key.p12 \
  --prompt-for-password
```

Or with a password file:

```bash
gcloud auth activate-service-account --key-file=/path/to/key.p12 \
  --password-file=/path/to/password.txt
```

This is the recommended method for production and scripting environments.

## Application Default Credentials (ADC)

ADC is a **separate credential system** used by Cloud Client Libraries, Google API Client Libraries, and SDKs (e.g., Terraform, Python `google-auth`, Node.js, Go).

### `gcloud auth application-default login`

Set up ADC with user credentials:

```bash
gcloud auth application-default login                    # Browser-based
gcloud auth application-default login --no-launch-browser # No browser
gcloud auth application-default login --no-launch-browser
```

Credentials are stored as a JSON file at a well-known location (`~/.config/gcloud/application_default_credentials.json`).

### ADC with Service Account Impersonation

```bash
gcloud auth application-default login \
  --impersonate-service-account=SA@PROJECT.iam.gserviceaccount.com
```

### ADC Print Access Token

```bash
gcloud auth application-default print-access-token
gcloud auth application-default print-access-token --scopes=SCOPE_LIST
```

### ADC Revoke

```bash
gcloud auth application-default revoke
```

### ADC Set Quota Project

```bash
gcloud auth application-default set-quota-project PROJECT_ID
```

### Key Distinction: gcloud auth login vs ADC

| | `gcloud auth login` | `gcloud auth application-default login` |
|---|---|---|
| **Used by** | gcloud CLI commands | Client libraries, SDKs, Terraform |
| **Storage** | `credentials.db` (SQLite) | `application_default_credentials.json` (JSON file) |
| **Scope** | gcloud CLI only | All Google Cloud client libraries |

This is a common source of confusion. See [gotchas.md](gotchas.md) for details.

### ADC Lookup Order

When client libraries look for credentials, the order is:

1. `GOOGLE_APPLICATION_CREDENTIALS` environment variable (path to a service account key JSON)
2. ADC file created by `gcloud auth application-default login`
3. Built-in service account (on Compute Engine, GKE, Cloud Run, Cloud Functions, etc.)

## Impersonation

### `--impersonate-service-account`

Make API requests as a service account without creating or downloading a key:

```bash
gcloud --impersonate-service-account=sa@project.iam.gserviceaccount.com compute instances list
```

Or set it as a persistent property:

```bash
gcloud config set auth/impersonate_service_account sa@project.iam.gserviceaccount.com
```

### Delegation Chains

Supports comma-separated delegation chains:

```bash
gcloud --impersonate-service-account=sa1,sa2,sa3 compute instances list
```

Requirements:
- The active account must have `roles/iam.serviceAccountTokenCreator` on `sa1`
- `sa1` must have the same role on `sa2`
- `sa2` must have the same role on `sa3`

The final token is for `sa3`, with the delegation chain `active → sa1 → sa2 → sa3`.

## Enterprise Certificate Authentication

Manage enterprise certificate configurations for smart card / PKCS #11 authentication:

```bash
# Linux
gcloud auth enterprise-certificate-config create linux \
  --label=LABEL --module=/path/to/module --slot=SLOT

# macOS
gcloud auth enterprise-certificate-config create macos \
  --issuer="Certificate Issuer"

# Windows
gcloud auth enterprise-certificate-config create windows \
  --issuer="Certificate Issuer"
```

## Docker Credential Helper

Register `gcloud` as a Docker credential helper for Google Container Registry (GCR) and Artifact Registry:

```bash
# GCR only (deprecated)
gcloud auth configure-docker

# GCR + Artifact Registry
gcloud auth configure-docker --include-artifact-registry
```

This modifies `~/.docker/config.json` to use `gcloud` as a credential helper for `gcr.io` and `pkg.dev` domains.

## Credential File Override

For advanced use cases, a credential file can be specified:

```bash
gcloud --credential-file-override=/path/to/credentials.json ...
```

This is an internal/hidden flag and typically not needed for normal usage.

## References

- [gcloud auth reference](https://cloud.google.com/sdk/gcloud/reference/auth)
- [Application Default Credentials](https://cloud.google.com/docs/authentication/application-default-credentials)
- [Provide credentials via ADC](https://cloud.google.com/docs/authentication/provide-credentials-adc)

## Related Skills

- For global flags related to auth (`--account`, `--impersonate-service-account`, `--access-token-file`), see [global-flags.md](global-flags.md).
- For auth-related config properties (`auth/access_token_file`, `auth/impersonate_service_account`), see [config-properties.md](config-properties.md).
- For the ADC confusion gotcha, see [gotchas.md](gotchas.md).
