# Command Groups

The major gcloud service groups, their subcommands, and common usage patterns.

> **Source of truth**: `gcloud --help` and the YAML specs in this project (`cmd/carapace-gcloud/cmd/gcloud/gcloud.*.yaml`).

## Core Infrastructure

### `compute` — Compute Engine

Create and manipulate Compute Engine resources (VMs, networks, disks, load balancers).

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| `instances` | `create`, `list`, `describe`, `delete`, `start`, `stop`, `reset`, `ssh`, `scp` | `--zone`, `--machine-type`, `--image`, `--preemptible`, `--no-restart-on-failure` |
| `disks` | `create`, `list`, `describe`, `delete`, `resize`, `snapshot` | `--zone`, `--size`, `--type`, `--source-snapshot` |
| `images` | `create`, `list`, `describe`, `delete`, `deprecate` | `--source-disk`, `--family` |
| `firewall-rules` | `create`, `list`, `describe`, `delete`, `update` | `--action`, `--allow`, `--rules`, `--network`, `--source-ranges` |
| `networks` | `create`, `list`, `describe`, `delete`, `update` | `--subnet-mode`, `--bgp-routing-mode` |
| `subnets` | `create`, `list`, `describe`, `delete`, `update` | `--range`, `--region`, `--network` |
| `addresses` | `create`, `list`, `describe`, `delete`, `move` | `--region`, `--global`, `--ip-version` |
| `instance-templates` | `create`, `list`, `describe`, `delete` | `--machine-type`, `--image` |
| `instance-groups` | `create`, `list`, `describe`, `delete` | `--zone`, `--region` |
| `zones` | `list`, `describe` | — |
| `regions` | `list`, `describe` | — |

### `storage` — Cloud Storage

Create and manage Cloud Storage buckets and objects.

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| `buckets` | `create`, `list`, `describe`, `delete`, `update`, `add-iam-policy-binding` | `--location`, `--uniform-bucket-level-access` |
| `objects` | `list`, `describe`, `update`, `compose`, `add-iam-policy-binding` | `--bucket` |
| `batch-operations` | `jobs create`, `jobs list`, `jobs cancel` | `--bucket`, `--source` |

### `container` — GKE (Kubernetes Engine)

Deploy and manage clusters of machines for running containers.

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| `clusters` | `create`, `list`, `describe`, `delete`, `update`, `resize`, `get-credentials` | `--zone`, `--region`, `--num-nodes`, `--machine-type`, `--cluster-version` |
| `node-pools` | `create`, `list`, `describe`, `delete`, `update` | `--cluster`, `--num-nodes`, `--machine-type` |
| `fleet` | `list`, `describe` | `--project` |
| `ai` | `profiles list`, `profiles manifests create` | `--model`, `--model-server` |

### `sql` — Cloud SQL

Create and manage Google Cloud SQL databases.

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| `instances` | `create`, `list`, `describe`, `delete`, `patch`, `clone`, `restart` | `--database-version`, `--tier`, `--region`, `--async` |
| `databases` | `create`, `list`, `describe`, `delete`, `patch` | `--instance` |
| `backups` | `create`, `list`, `describe`, `delete`, `restore` | `--instance`, `--async` |
| `users` | `create`, `list`, `delete`, `set-password` | `--instance` |
| `ssl-certs` (deprecated) | `create`, `list`, `describe`, `delete` | `--instance` |
| `export` | `sql` | `--database`, `--async`, `--offload`, `--parallel` |
| `import` | `sql` | `--database`, `--async`, `--parallel` |

### `run` — Cloud Run

Manage your Cloud Run applications.

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| `services` | `create`, `list`, `describe`, `delete`, `update`, `deploy` | `--image`, `--region`, `--platform`, `--port`, `--memory` |
| `jobs` | `create`, `list`, `describe`, `delete`, `execute`, `update` | `--image`, `--region`, `--tasks` |
| `revisions` | `list`, `describe`, `delete` | `--region` |

## Security & Identity

### `iam` — Identity and Access Management

Manage IAM service accounts and keys.

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| `service-accounts` | `create`, `list`, `describe`, `delete`, `enable`, `disable`, `add-iam-policy-binding` | `--display-name` |
| `service-accounts keys` | `create`, `list`, `delete`, `enable`, `disable`, `upload` | `--iam-account`, `--key-file-type` |
| `workforce-pools` | `create`, `list`, `describe`, `delete`, `update` | `--location`, `--description` |
| `workload-identity-pools` | `create`, `list`, `describe`, `delete`, `update` | `--location` |
| `oauth-clients` | `create`, `list`, `describe`, `delete` | `--location`, `--client-type` |
| `list-grantable-roles` | — | `--filter` |
| `list-testable-permissions` | — | `--filter` |

### `kms` — Key Management Service

Manage cryptographic keys in the cloud.

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| `keyrings` | `create`, `list`, `describe` | `--location` |
| `keys` | `create`, `list`, `describe`, `delete`, `update`, `enable`, `disable` | `--keyring`, `--location`, `--purpose`, `--rotation-period` |
| `import-jobs` | `create`, `list`, `describe` | `--keyring`, `--location` |
| `encrypt` | — | `--key`, `--keyring`, `--location`, `--plaintext-file` |
| `decrypt` | — | `--key`, `--keyring`, `--location`, `--ciphertext-file` |

### `secrets` — Secret Manager

Manage secrets on Google Cloud.

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| (top-level) | `create`, `list`, `describe`, `delete`, `update` | `--replication-policy`, `--locations` |
| `versions` | `access`, `add`, `disable`, `enable`, `destroy`, `list`, `describe` | `--secret` |

## Data & Analytics

### `bigtable` — Cloud Bigtable

Manage your Cloud Bigtable storage.

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| `instances` | `create`, `list`, `describe`, `delete`, `update` | `--display-name`, `--cluster-storage-type` |
| `clusters` | `create`, `list`, `describe`, `update` | `--instance`, `--storage-type` |
| `tables` | `create`, `list`, `describe`, `delete` | `--instance` |

### `spanner` — Cloud Spanner

Command groups for Cloud Spanner.

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| `instances` | `create`, `list`, `describe`, `delete`, `update` | `--config`, `--description`, `--nodes` |
| `databases` | `create`, `list`, `describe`, `delete`, `update` | `--instance`, `--ddl` |
| `backups` | `create`, `list`, `describe`, `delete` | `--instance`, `--database` |

### `dataproc` — Cloud Dataproc

Create and manage Google Cloud Dataproc clusters and jobs.

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| `clusters` | `create`, `list`, `describe`, `delete`, `update`, `start`, `stop`, `diagnose` | `--region`, `--image-version`, `--master-machine-type` |
| `jobs` | `submit`, `list`, `describe`, `delete`, `kill`, `wait` | `--region`, `--cluster` |

### `dataflow` — Cloud Dataflow

Manage Google Cloud Dataflow resources.

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| `jobs` | `run`, `list`, `describe`, `cancel`, `show`, `drain` | `--region`, `--gcs-location`, `--parameters` |
| `flex-template` | `run`, `build` | `--region`, `--template-file` |

### `pubsub` — Cloud Pub/Sub

Manage Cloud Pub/Sub topics, subscriptions, and snapshots.

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| `topics` | `create`, `list`, `describe`, `delete`, `update`, `publish` | `--message-retention-duration` |
| `subscriptions` | `create`, `list`, `describe`, `delete`, `update`, `pull`, `ack` | `--topic`, `--ack-deadline` |
| `snapshots` | `create`, `list`, `describe`, `delete` | `--subscription` |

## Application Platforms

### `app` — App Engine

Manage your App Engine deployments.

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| `deploy` | — | `--image-url`, `--version`, `--no-promote` |
| `versions` | `list`, `browse`, `start`, `stop`, `delete`, `migrate` | `--service` |
| `services` | `list`, `describe`, `delete`, `browse` | — |
| `instances` | `list`, `describe`, `delete`, `enable-debug`, `disable-debug`, `ssh`, `scp` | `--service`, `--version` |
| `logs` | `read`, `tail` | `--service`, `--version` |
| `create` | — | `--service-account` |

### `functions` — Cloud Functions

Manage Google Cloud Functions.

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| `deploy` | — | `--trigger-http`, `--trigger-topic`, `--trigger-bucket`, `--runtime`, `--entry-point`, `--source` |
| `list` | — | `--region`, `--filter` |
| `describe` | — | `--region` |
| `delete` | — | `--region` |
| `call` | — | `--region`, `--data` |
| `logs` | `read`, `tail` | `--region` |

## Management & Operations

### `logging` — Cloud Logging

Manage Cloud Logging.

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| `logs` | `list`, `describe`, `delete` | — |
| `sinks` | `create`, `list`, `describe`, `delete`, `update` | `--log-filter`, `--destination` |
| `metrics` | `create`, `list`, `describe`, `delete`, `update` | `--log-filter`, `--description` |
| `tail` (beta-only) | — | `--filter`, `--format` |
| `read` | — | `--filter`, `--limit` |

### `monitoring` — Cloud Monitoring

Manage Cloud Monitoring dashboards.

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| `dashboards` | `create`, `list`, `describe`, `delete` | `--config` |
| `policies` (alpha-only) | `list`, `describe` | `--filter` |
| `channels` (alpha-only) | `create`, `list`, `describe`, `delete`, `update` | `--type`, `--description` |

### `builds` — Cloud Build

Create and manage builds for Google Cloud Build.

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| `submit` | — | `--tag`, `--config`, `--substitutions`, `--no-cache` |
| `list` | — | `--filter`, `--limit` |
| `describe` | — | — |
| `log` | — | `--stream` |
| `triggers` | `create`, `list`, `describe`, `delete`, `update`, `run` | `--repository`, `--branch-pattern`, `--build-config` |

### `deploy` — Cloud Deploy

Create and manage Cloud Deploy resources.

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| `delivery-pipelines` | `create`, `list`, `describe`, `delete`, `update` | `--region`, `--description` |
| `targets` | `create`, `list`, `describe`, `delete`, `update` | `--region`, `--execution-config` |
| `releases` | `create`, `list`, `describe`, `promote`, `abandon` | `--delivery-pipeline`, `--region` |

## Networking

### `dns` — Cloud DNS

Manage your Cloud DNS managed-zones and record-sets.

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| `managed-zones` | `create`, `list`, `describe`, `delete`, `update` | `--dns-name`, `--description`, `--visibility` |
| `record-sets` | `create`, `list`, `describe`, `delete`, `update`, `import`, `export` | `--zone`, `--type`, `--ttl`, `--rrdatas` |
| `policies` | `create`, `list`, `describe`, `delete`, `update` | `--description` |

### `network-connectivity` — Network Connectivity

Manage Network Connectivity resources.

| Sub-group | Key Commands |
|-----------|-------------|
| `hubs` | `create`, `list`, `describe`, `delete`, `update` |
| `spokes` | `create`, `list`, `describe`, `delete` |

## Serverless & Eventing

### `tasks` — Cloud Tasks

Manage Cloud Tasks queues and tasks.

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| `queues` | `create`, `list`, `describe`, `delete`, `update`, `pause`, `resume` | `--location`, `--max-attempts` |
| `create-http-task` | — | `--queue`, `--location`, `--url` |
| `create-app-engine-task` | — | `--queue`, `--location` |

### `scheduler` — Cloud Scheduler

Manage Cloud Scheduler jobs and schedules.

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| `jobs` | `create`, `list`, `describe`, `delete`, `update`, `pause`, `resume` | `--location`, `--schedule`, `--time-zone`, `--http-method` |

### `workflows` — Cloud Workflows

Manage your Cloud Workflows resources.

| Sub-group | Key Commands | Notable Flags |
|-----------|-------------|---------------|
| `deploy` | — | `--source`, `--location`, `--call-log-level` |
| `list` | — | `--location`, `--filter` |
| `describe` | — | `--location` |
| `delete` | — | `--location` |
| `execute` | — | `--location`, `--data` |

## References

- [gcloud command reference](https://cloud.google.com/sdk/gcloud/reference)
- YAML specs in this project: `cmd/carapace-gcloud/cmd/gcloud/gcloud.*.yaml`

## Related Skills

- For command structure and dispatch model, see [architecture.md](architecture.md).
- For global flags used across all groups, see [global-flags.md](global-flags.md).
