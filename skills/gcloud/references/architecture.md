# gcloud CLI Architecture

The structural design and command dispatch model of the Google Cloud CLI.

## Command Pattern

Most `gcloud` commands follow a consistent pattern:

```
gcloud + [release level] + component + entity + operation + [positional args] + [flags]
```

Example:

```
gcloud compute instances create example-instance-1 --zone=us-central1-a
```

| Part | Description | Example |
|------|-------------|---------|
| **Release level** | Optional: `alpha`, `beta`, `preview`, or omitted for GA | `gcloud alpha compute` |
| **Component** | Google Cloud service group | `compute`, `app`, `storage`, `iam`, `container` |
| **Entity** | Plural form of the resource | `instances`, `disks`, `firewall-rules`, `zones` |
| **Operation** | Imperative verb | `create`, `list`, `describe`, `delete`, `update`, `deploy` |
| **Positional args** | Order-specific required arguments | `<INSTANCE_NAMES>` |
| **Flags** | `--flag-name(=value)` | `--zone=us-central1-a`, `--preemptible` |

## Hierarchical Command Tree

The gcloud CLI is a hierarchical tree of commands. Non-leaf nodes are **command groups** (representing products or features), and leaf nodes are **commands** (representing operations).

```
gcloud
├── auth                    (command group)
│   ├── login               (command)
│   ├── activate-service-account
│   ├── application-default  (sub-group)
│   │   ├── login
│   │   ├── print-access-token
│   │   └── revoke
│   ├── configure-docker
│   └── ...
├── compute                 (command group)
│   ├── instances           (sub-group)
│   │   ├── create
│   │   ├── list
│   │   ├── describe
│   │   ├── delete
│   │   └── ...
│   ├── disks
│   ├── firewall-rules
│   └── ...
├── config
├── container
├── iam
├── projects
├── storage
├── sql
└── ... (100+ service groups)
```

## Command Group Categories

### Core Infrastructure

| Group | Description | Key Entities |
|-------|-------------|--------------|
| `compute` | Compute Engine | instances, disks, images, firewall-rules, networks, subnets, addresses |
| `storage` | Cloud Storage | buckets, objects, batch-operations |
| `container` | GKE (Kubernetes) | clusters, node-pools, fleets |
| `sql` | Cloud SQL | instances, databases, backups |
| `run` | Cloud Run | services, jobs, revisions |

### Security & Identity

| Group | Description | Key Entities |
|-------|-------------|--------------|
| `iam` | IAM service accounts | service-accounts, keys, workforce-pools |
| `auth` | OAuth2 credentials | login, activate-service-account, application-default |
| `kms` | Key Management Service | keyrings, keys, import-jobs |
| `secrets` | Secret Manager | secrets, versions |
| `privateca` | Private Certificate Authorities | certificates, certificate-authorities |

### Data & Analytics

| Group | Description | Key Entities |
|-------|-------------|--------------|
| `bigtable` | Cloud Bigtable | instances, clusters, tables |
| `spanner` | Cloud Spanner | instances, databases, backups |
| `dataproc` | Cloud Dataproc | clusters, jobs |
| `dataflow` | Cloud Dataflow | jobs, flex-template |
| `pubsub` | Cloud Pub/Sub | topics, subscriptions, snapshots |

### Management & Operations

| Group | Description | Key Entities |
|-------|-------------|--------------|
| `projects` | Project management | create, delete, describe, list, add-iam-policy-binding |
| `organizations` | Organization management | create, describe, list, add-iam-policy-binding |
| `billing` | Billing accounts | accounts, projects, links |
| `config` | CLI properties | set, get, list, configurations |
| `components` | CLI components | install, list, update, remove |
| `logging` | Cloud Logging | logs, sinks, metrics |
| `monitoring` | Cloud Monitoring | dashboards, alerts, policies |

## Common Operations

Most resource entities support a consistent set of CRUD operations:

| Operation | Description | Common Flags |
|-----------|-------------|--------------|
| `create` | Create a new resource | `--zone`, `--region`, `--async` |
| `list` | List resources | `--filter`, `--limit`, `--sort-by`, `--page-size`, `--uri` |
| `describe` | Show details of a resource | `--zone`, `--region` |
| `delete` | Delete a resource | `--async`, `--quiet` |
| `update` | Modify a resource | `--zone`, `--region` |
| `export` | Export resource data | `--destination` |

## Dispatch Model

When a gcloud command is invoked:

1. **Parse the command line** — gcloud splits the input into the command path (e.g., `compute instances create`), positional arguments, and flags
2. **Resolve the command** — the command path is matched against the command tree to find the appropriate handler
3. **Resolve configuration** — global flags, environment variables, and config properties are merged using precedence rules (see [global-flags.md](global-flags.md) and [config-properties.md](config-properties.md))
4. **Authenticate** — credentials are looked up based on the resolved account, service account impersonation, or ADC (see [auth.md](auth.md))
5. **Execute the API call** — the command handler makes the appropriate Google Cloud API request with the resolved project, billing project, and other context
6. **Process the response** — output is processed through `--flatten` → `--sort-by` → `--filter` → `--limit` (see [output-formatting.md](output-formatting.md))
7. **Format output** — the response is formatted according to `--format` and written to stdout (results) or stderr (warnings, prompts)

## Top-Level Commands (Non-Service)

These are standalone commands outside the service group hierarchy:

| Command | Purpose |
|---------|---------|
| `gcloud init` | Initialize or reinitialize gcloud (interactive setup) |
| `gcloud version` | Print version information for all components |
| `gcloud info` | Display information about the current gcloud environment |
| `gcloud help` | Search gcloud help text |
| `gcloud feedback` | Provide feedback to the Google Cloud CLI team |
| `gcloud survey` | Invoke a customer satisfaction survey |
| `gcloud cheat-sheet` | Display gcloud cheat sheet |
| `gcloud topic` | Supplementary help on topics (formats, filters, configurations, etc.) |

## Topic Commands

`gcloud topic` provides supplementary documentation on cross-cutting concepts:

| Topic | Covers |
|-------|--------|
| `gcloud topic formats` | Output format specifications |
| `gcloud topic filters` | Filter expression syntax |
| `gcloud topic projections` | Field projection and transforms |
| `gcloud topic configurations` | Named configurations |
| `gcloud topic flags-file` | YAML/JSON flags file format |
| `gcloud topic resource-keys` | Resource key naming |
| `gcloud topic startup` | Python interpreter configuration |
| `gcloud topic accessibility` | Screen reader support |
| `gcloud topic escaping` | Shell escaping rules |

## References

- [Official gcloud reference](https://cloud.google.com/sdk/gcloud/reference)
- [gcloud cheat sheet](https://cloud.google.com/sdk/docs/cheatsheet)

## Related Skills

- For the carapace-gcloud completion provider's use of this structure, see the **carapace** skill.
- For details on specific service groups, see [command-groups.md](command-groups.md).
- For global flags, see [global-flags.md](global-flags.md).
