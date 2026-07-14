---
name: gcloud
description: >
  Use when working with the Google Cloud CLI (gcloud) — command structure, global flags,
  configuration system, authentication, resource hierarchy, output formatting, release tracks,
  and command groups. Triggers on: "gcloud", "gcloud cli", "google cloud cli", "gcloud config",
  "gcloud auth", "gcloud compute", "gcloud storage", "gcloud sql", "gcloud container", "gcloud iam",
  "gcloud projects", "gcloud organizations", "gcloud billing", "gcloud components", "gcloud alpha",
  "gcloud beta", "gcloud preview", "Application Default Credentials", "ADC", "service account",
  "impersonate-service-account", "--filter", "--format", "--flatten", "CLOUDSDK_",
  "gcloud configurations", "gcloud properties", "gcloud topic", "gsutil", "gcloud init".
user-invocable: true
---

# gcloud CLI In-Depth Reference

Comprehensive reference for the Google Cloud CLI (`gcloud`) — the unified command-line tool for managing Google Cloud Platform resources and services.

## Data Flow

```
gcloud command line
  → flag parsing (global flags + command flags)
    → configuration resolution (flags > env vars > config properties)
      → credential lookup (account/service-account/ADC)
        → API request (with project, billing, impersonation)
          → response processing (--flatten → --sort-by → --filter → --limit)
            → output formatting (--format)
              → stdout (results) / stderr (warnings, prompts)
```

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| command structure, command pattern, command groups, dispatch, CLI architecture, gcloud layout, component, entity, operation | [references/architecture.md](references/architecture.md) |
| service groups, command groups, compute, storage, sql, container, iam, app, functions, run, pubsub, logging, monitoring, bigtable, dataproc, dataflow, spanner, kms, dns, builds, deploy | [references/command-groups.md](references/command-groups.md) |
| global flags, --project, --account, --format, --filter, --flatten, --quiet, --verbosity, --billing-project, --impersonate-service-account, --flags-file, --log-http, --access-token-file, --trace-token, --configuration, flag precedence | [references/global-flags.md](references/global-flags.md) |
| config, properties, configurations, profiles, gcloud config set, gcloud config get, gcloud config list, CLOUDSDK_, environment variables, CLOUDSDK_CONFIG, named configurations | [references/config-properties.md](references/config-properties.md) |
| auth, authentication, credentials, service account, ADC, Application Default Credentials, gcloud auth login, gcloud auth activate-service-account, impersonation, docker credential helper, access token, refresh token | [references/auth.md](references/auth.md) |
| resource hierarchy, organization, folder, project, project ID, project number, billing, IAM policy, gcloud projects, gcloud organizations, gcloud billing, gcloud resource-manager | [references/resource-hierarchy.md](references/resource-hierarchy.md) |
| output formatting, --format, --filter, --flatten, --sort-by, --limit, projections, transforms, json, yaml, table, csv, value, text, filter expressions, filter operators | [references/output-formatting.md](references/output-formatting.md) |
| release tracks, alpha, beta, preview, GA, general availability, components, gcloud components install, gcloud components list, component manager, kubectl, optional components | [references/release-tracks.md](references/release-tracks.md) |
| gotchas, edge cases, known issues, pitfalls, stdout vs stderr, project ID vs name, quiet mode, filter gotchas, component manager disabled, Python version, Cloud Shell, ADC confusion | [references/gotchas.md](references/gotchas.md) |

## Quick Guide

- **How does the gcloud command structure work?** → [references/architecture.md](references/architecture.md)
- **What service groups are available?** → [references/command-groups.md](references/command-groups.md)
- **What global flags exist and how do they work?** → [references/global-flags.md](references/global-flags.md)
- **How do I configure gcloud properties?** → [references/config-properties.md](references/config-properties.md)
- **How do I authenticate with gcloud?** → [references/auth.md](references/auth.md)
- **How does the resource hierarchy work?** → [references/resource-hierarchy.md](references/resource-hierarchy.md)
- **How do I format and filter output?** → [references/output-formatting.md](references/output-formatting.md)
- **What are the release tracks?** → [references/release-tracks.md](references/release-tracks.md)
- **What are common gotchas and pitfalls?** → [references/gotchas.md](references/gotchas.md)
- **How do flag overrides and precedence work?** → [references/global-flags.md](references/global-flags.md) and [references/config-properties.md](references/config-properties.md)
- **How do I set up ADC for client libraries?** → [references/auth.md](references/auth.md)
- **How do I use impersonation?** → [references/auth.md](references/auth.md)

## Cross-Project References

- For shell completion integration with gcloud (carapace-bridge, carapace-spec), see the **carapace** skill and **carapace-dev** skill.
- For cobra command structure patterns used by carapace-gcloud, see the **cobra** skill.
- For YAML spec format used by carapace-spec-gcloud, see the **carapace** skill → spec documentation.
