# Release Tracks & Components

The gcloud CLI release track system (alpha, beta, preview, GA) and component management.

> **Source of truth**: [gcloud components docs](https://cloud.google.com/sdk/docs/components) and [gcloud release levels](https://cloud.google.com/sdk/docs/gcloud).

## Release Tracks

Commands in gcloud are categorized by release level, indicating their stability and support status.

| Release Level | Command Prefix | Label | Installed by Default | Stability |
|---------------|---------------|-------|---------------------|-----------|
| **General Availability** | (none) | — | Yes | Fully stable, production-ready |
| **Preview** | `preview` | *(PREVIEW)* | No | Feature-complete but no SLAs, test environments only |
| **Beta** | `beta` | *(BETA)* | No | Functionally complete, some outstanding issues |
| **Alpha** | `alpha` | *(ALPHA)* | No | Early release, may change without notice |

### Usage Pattern

```bash
# GA (stable)
gcloud compute instances list

# Preview
gcloud preview compute instances list

# Beta
gcloud beta compute instances list

# Alpha
gcloud alpha compute instances list
```

Some commands **only exist** at alpha or beta levels and never graduate to GA.

### Installing Release Track Components

```bash
gcloud components install alpha
gcloud components install beta
gcloud components install preview
```

If you run an uninstalled alpha/beta/preview command interactively, gcloud prompts you to install it on-the-spot.

### Release Track Behavior

- **GA commands**: Breaking changes are announced in release notes with migration periods
- **Preview commands**: No SLAs or support commitments; average lifetime ~6 months before graduation or removal
- **Beta commands**: Breaking changes can be made without notice
- **Alpha commands**: No guarantees — may change or be removed at any time

## Component Management

The gcloud CLI is modular — components can be installed or removed as needed.

### Key Commands

| Command | Purpose |
|---------|---------|
| `gcloud components list` | List installed and available components |
| `gcloud components install [ID]` | Install one or more components |
| `gcloud components remove [ID]` | Remove one or more components |
| `gcloud components update` | Update all components to latest version |
| `gcloud components reinstall` | Reinstall the CLI with current components |
| `gcloud components post-process` | Post-installation steps (hidden) |

### Installed by Default

| Component | Description |
|-----------|-------------|
| `gcloud` | GA gcloud commands |
| `bq` | BigQuery command-line tool |
| `gsutil` | Cloud Storage command-line tool |
| `core` | Core dependencies |

### Notable Optional Components

| Component | Description |
|-----------|-------------|
| `alpha` | Alpha gcloud commands |
| `beta` | Beta gcloud commands |
| `preview` | Preview gcloud commands |
| `kubectl` | Kubernetes command-line tool |
| `gke-gcloud-auth-plugin` | GKE authentication plugin |
| `app-engine-java` | App Engine Java SDK |
| `app-engine-go` | App Engine Go SDK |
| `app-engine-python` | App Engine Python SDK |
| `cbt` | Cloud Bigtable command-line tool |
| `cloud-build-local` | Local Cloud Build runner |
| `cloud-run-proxy` | Cloud Run proxy |
| `skaffold` | Skaffold tool |
| `bigtable-emulator` | Bigtable local emulator |
| `datastore-emulator` | Datastore local emulator |
| `firestore-emulator` | Firestore local emulator |
| `pubsub-emulator` | Pub/Sub local emulator |
| `spanner-emulator` | Spanner local emulator |

### Listing Components

```bash
gcloud components list
gcloud components list --show-versions
gcloud components list --show-platform
gcloud components list --only-local-state
gcloud components list --filter="state.name:Installed"
```

### Updating

```bash
# Update all components
gcloud components update

# Update to a specific version
gcloud components update --version=450.0.0
```

### Component Manager Disabled with Package Managers

If the gcloud CLI was installed via APT (`google-cloud-cli`) or YUM/DNF, the component manager is **disabled**. Use the package manager instead:

```bash
# APT
sudo apt-get install google-cloud-cli-gke-gcloud-auth-plugin
sudo apt-get install google-cloud-cli-pubsub-emulator

# YUM/DNF
sudo dnf install google-cloud-cli-gke-gcloud-auth-plugin
```

This is a common gotcha — see [gotchas.md](gotchas.md).

## Version Information

```bash
gcloud version                    # Print version for all components
gcloud version --format=json      # JSON format
gcloud info                       # Display environment info
```

## References

- [gcloud components](https://cloud.google.com/sdk/docs/components)
- [gcloud release levels](https://cloud.google.com/sdk/docs/gcloud)

## Related Skills

- For the `gcloud info` command and environment details, see [architecture.md](architecture.md).
- For the component manager disabled gotcha, see [gotchas.md](gotchas.md).
