# Output Formatting

The gcloud CLI output processing pipeline: `--format`, `--filter`, `--flatten`, `--sort-by`, `--limit`, projections, and transforms.

> **Source of truth**: `gcloud topic formats`, `gcloud topic filters`, `gcloud topic projections`, and [gcloud formatting docs](https://cloud.google.com/sdk/gcloud/reference/topic/formats).

## Processing Pipeline

Output is processed in a fixed order:

```
API Response
  → --flatten     (flatten repeated fields into separate records)
    → --sort-by    (sort records)
      → --filter    (filter records by expression)
        → --limit    (limit number of results)
          → --format   (format for display)
            → stdout
```

This order matters — filtering happens **after** flattening, so you can filter on flattened fields.

## `--format`

The format expression has three parts: `NAME[ATTRIBUTES](PROJECTION)`

### Supported Formats

| Format | Description | Requires Projection? |
|--------|-------------|---------------------|
| `json` | JSON output | No |
| `yaml` | YAML output (default for most commands) | No |
| `csv` | Comma-separated values | Yes |
| `table` | Aligned columns with headings | Yes |
| `value` | Tab-separated values (for scripting) | Yes |
| `text` / `flattened` | `key: value` pairs per line | No |
| `list` | Ordered list of items | No |
| `config` | Config-style dictionary of dictionaries | No |
| `diff` | Unified diff of first two projection columns | Yes |
| `multi` | Multiple sub-formats within one command | No |
| `get` | Value without transforms | No |
| `object` | Bypass JSON serialization | No |
| `none` / `disable` | Suppress output | No |

### Format Attributes

Attributes are specified in brackets `[...]` after the format name:

```bash
--format="table[box,title=Instances](name:sort=1, zone:label=zone, status)"
--format="table[all-box,no-heading](name, id)"
--format="csv[separator=','][no-heading](projectId)"
--format="yaml[no-undefined]"
```

### Projections

Projections select and transform fields, specified in parentheses `(...)`:

```bash
--format="table(name, zone, status)"
--format="value(name, selfLink)"
--format="json(name, id, status)"
```

### Common Format Examples

```bash
# JSON output
gcloud compute instances list --format=json

# YAML output
gcloud projects describe my-project --format=yaml

# Table with custom headings
gcloud compute instances list \
  --format="table[box,title=Instances](name:sort=1, zone:label=Zone, status:label=Status)"

# Values only (for scripting)
gcloud projects list --format="value(projectId)"
gcloud compute instances list --format="value(name,zone)" | awk '{print $1, $2}'

# CSV
gcloud projects list --format="csv[no-heading](projectId,name)"

# Flattened key: value
gcloud projects describe my-project --format=flattened

# Suppress output
gcloud compute instances delete my-instance --format=none --quiet
```

### Projection Transforms

Transforms modify field values within projections:

```bash
# Date/time transform
--format="table(name, creationTimestamp.date(tz=LOCAL))"

# URL path segment extraction (scope)
--format="value(selfLink.scope())"

# Resource URI
--format="value(uri())"

# Nested scope/segment
--format="table(name.scope(keys).segment(0):label='keyID')"

# List join
--format="value(disks[].interface.list())"

# Basename transform
--format="flattened(name,serviceAccounts[].scopes[].basename())"

# Strip prefix
--format="value(name.segment(0))"

# Convert to boolean
--format="table(name, archived.bool())"
```

## `--filter`

Filter expressions select which resources to print.

### Operators

| Operator | Meaning | Example |
|----------|---------|---------|
| `:` | Contains / word match | `zone:us-central1` |
| `=` | Equality (or any-of with parentheses) | `zone=us-central1-a`, `zone=(a b)` |
| `!=` | Not equal | `zone!=us-central1-a` |
| `<`, `<=`, `>=`, `>` | Comparison (numeric or lexicographic) | `createTime>=2018-01-15` |
| `~` | Regex match (Python `re` syntax) | `zone ~ ^us` |
| `!~` | Regex not-match | `zone !~ ^us` |
| `:*` | Key defined (exists) | `labels.my-label:*` |
| `-key:*` | Key NOT defined | `-labels.my-label:*` |
| `AND`, `OR`, `NOT` | Logic (must be uppercase) | `zone ~ us AND -machineType:f1-micro` |

### Filter Examples

```bash
# Zone starts with "us" and not f1-micro machines
gcloud compute instances list --filter="zone ~ ^us AND -machineType:f1-micro"

# Multiple values (any-of)
gcloud compute instances list --filter="zone=(us-central1-a us-central1-b)"

# Label filtering
gcloud compute instances list --filter="labels.env=test AND labels.version=alpha"

# Date comparison
gcloud compute instances list --filter="createTime>=2018-01-15T12:00:00"

# ISO 8601 duration (last 2 weeks)
gcloud compute instances list --filter="createTime>-P2W"

# Tag filtering
gcloud compute instances list --filter="tags.items=(my-tag,my-other-tag)"

# Key existence
gcloud compute instances list --filter="labels.my-label:*"

# Key absence
gcloud compute instances list --filter="-labels.my-label:*"
```

### Filter Gotchas

- The `:` operator is **changing** — the deprecated default does substring matching, the new implementation does word matching. A warning appears when both would return different results.
- The `=` operator behaves **inconsistently** across different APIs — for some it means equality, for others it acts like `:`. This is being phased out.
- Most filter expressions need shell quoting. Use `'...'` for shell quotes and `"..."` for filter string literals.
- Scientific notation gotcha: `30e5504145` is interpreted as `30 × 10^5504145`, so values like this need quoting: `--filter="'tags:30e5504145'"`.

See [gotchas.md](gotchas.md) for more.

## `--flatten`

Flattens repeated/nested fields into separate records. Each record with N elements in the flattened field expands to N records.

```bash
# Flatten IAM bindings into separate records
gcloud projects get-iam-policy my-project \
  --flatten=bindings \
  --filter=bindings.role:roles/editor \
  --format='value(bindings.members)'

# Flatten multiple levels
gcloud projects get-iam-policy my-project \
  --flatten=bindings \
  --flatten=bindings.members \
  --format='table(bindings.role, bindings.members)'
```

### Flatten Syntax

```bash
--flatten=abc.def          # Flatten abc.def[] into separate records
--flatten=abc.def[].ghi    # Flatten to abc.def.ghi
```

## `--sort-by`

Sort results by one or more fields:

```bash
# Ascending
gcloud compute instances list --sort-by=createTime

# Descending (prefix with ~)
gcloud compute instances list --sort-by=~createTime

# Multiple fields
gcloud compute instances list --sort-by=zone,name
```

## `--limit`

Limit the number of results:

```bash
gcloud compute instances list --limit=10
gcloud projects list --limit=5
```

## `--page-size`

Control the API page size (number of results per API request):

```bash
gcloud compute instances list --page-size=50
```

This is different from `--limit` — `--page-size` controls API pagination, `--limit` controls total output.

## `--uri`

Print resource URIs instead of default output:

```bash
gcloud compute instances list --uri
```

## Combined Example

```bash
gcloud compute instances list \
  --flatten=networkInterfaces \
  --filter="networkInterfaces.accessConfigs[0].natIP:*" \
  --sort-by=~createTime \
  --limit=10 \
  --format="table(name:label='Instance', zone:label='Zone', \
    networkInterfaces.accessConfigs[0].natIP:label='External IP', \
    status:label='Status')"
```

## References

- `gcloud topic formats`
- `gcloud topic filters`
- `gcloud topic projections`
- `gcloud topic resource-keys`
- [gcloud formatting reference](https://cloud.google.com/sdk/gcloud/reference/topic/formats)

## Related Skills

- For global flags including `--format`, `--filter`, `--flatten`, see [global-flags.md](global-flags.md).
- For filter gotchas and edge cases, see [gotchas.md](gotchas.md).
