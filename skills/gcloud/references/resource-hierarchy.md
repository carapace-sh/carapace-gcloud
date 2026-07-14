# Resource Hierarchy

The Google Cloud resource hierarchy: organizations, folders, projects, and service resources. How gcloud manages them.

> **Source of truth**: [Cloud Resource Hierarchy](https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy).

## Hierarchy Structure

```
Organization (root)
    ├── Folder(s)
    │       ├── Folder(s) (nested)
    │       └── Project(s)
    └── Project(s)
            └── Service Resources (VMs, buckets, databases, etc.)
```

IAM policies applied at any level are **inherited by all child resources**. This is the key design principle — policy inheritance flows downward.

## Organization

The root node of the resource hierarchy, representing a company or entity.

- Requires a **Google Workspace** or **Cloud Identity** account
- The Google Workspace super admin can assign the **Organization Administrator** role
- Once an Organization exists, managed users must create projects **within** it

### Key Commands

```bash
gcloud organizations list
gcloud organizations describe ORG_ID
gcloud organizations add-iam-policy-binding ORG_ID \
  --member=user:email@gmail.com --role=roles/resourcemanager.organizationAdmin
```

### Attributes

| Attribute | Description |
|-----------|-------------|
| `organizations/ORG_ID` | Resource path |
| `displayName` | Human-readable name |
| `lifecycleState` | `ACTIVE` or `DELETE_REQUESTED` |
| `owner.directoryCustomerId` | Workspace/Cloud Identity customer ID |

## Folders

Optional grouping mechanism within an organization. Can contain projects and other folders (nesting).

- Requires an Organization to exist
- Acts as a policy inheritance point
- Useful for departmental/team separation

### Key Commands

```bash
gcloud resource-manager folders create --display-name="Engineering" \
  --organization=ORG_ID
gcloud resource-manager folders list --organization=ORG_ID
gcloud resource-manager folders describe FOLDER_ID
gcloud resource-manager folders update FOLDER_ID --display-name="New Name"
gcloud resource-manager folders delete FOLDER_ID
```

## Projects

The **fundamental organizing entity** in Google Cloud — required to use any service.

### Project Identifiers

A project has **three identifiers** — confusing them is a common mistake:

| Identifier | Type | Mutable | Example | Usage |
|-----------|------|---------|---------|-------|
| **Project ID** | String | No (immutable) | `my-project-123` | Used in commands and gcloud flags |
| **Project Number** | Integer | No (auto-assigned) | `464036093014` | Can be used interchangeably with ID |
| **Name** | String | Yes (mutable) | `My Project` | Display name only — **not** for commands |

### Key Commands

```bash
# List projects
gcloud projects list
gcloud projects list --filter="lifecycleState:ACTIVE"
gcloud projects list --format="table(projectId,name,projectNumber)"

# Describe a project
gcloud projects describe PROJECT_ID
gcloud projects describe 464036093014  # by number

# Create a project
gcloud projects create my-project-123 \
  --name="My Project" --organization=ORG_ID
gcloud projects create my-project-123 --folder=FOLDER_ID

# Delete a project
gcloud projects delete PROJECT_ID

# Update
gcloud projects update PROJECT_ID --name="Updated Name"
```

### Project IAM

```bash
# Add IAM policy binding
gcloud projects add-iam-policy-binding PROJECT_ID \
  --member=user:email@gmail.com --role=roles/editor

# Remove IAM policy binding
gcloud projects remove-iam-policy-binding PROJECT_ID \
  --member=user:email@gmail.com --role=roles/editor

# Get IAM policy
gcloud projects get-iam-policy PROJECT_ID

# Get ancestors (hierarchy path)
gcloud projects get-ancestors PROJECT_ID

# Get IAM policies for project and all ancestors
gcloud projects get-ancestors-iam-policy PROJECT_ID
```

### Project Creation Flags

| Flag | Description |
|------|-------------|
| `--name` | Display name for the project |
| `--organization` | ID of the organization to use as parent |
| `--folder` | ID of the folder to use as parent |
| `--enable-cloud-apis` | Enable `cloudapis.googleapis.com` during creation |
| `--set-as-default` | Set as `core/project` property after creation |
| `--labels` | List of `KEY=VALUE` label pairs |
| `--tags` | List of `KEY=VALUE` tag pairs to bind |

## Billing

### Key Commands

```bash
# List billing accounts
gcloud billing accounts list

# List projects linked to a billing account
gcloud billing projects list --billing-account=ACCOUNT_ID

# Link a project to a billing account
gcloud billing projects link PROJECT_ID --billing-account=ACCOUNT_ID

# Unlink a project from billing
gcloud billing projects unlink PROJECT_ID

# Describe a billing account
gcloud billing accounts describe ACCOUNT_ID
```

### Billing vs Quota Project

The `--project` flag serves two roles:
1. Specifies the project of the resource to operate on
2. Specifies the project for API quota and billing checks

Use `--billing-project` (or `billing/quota_project` property) to decouple these:

```bash
# Operate on resources in my-resource-project,
# but charge quota to my-quota-project
gcloud compute instances list \
  --project=my-resource-project \
  --billing-project=my-quota-project
```

## Resource Manager

The `resource-manager` command group manages cloud resources at the organization and folder level:

```bash
gcloud resource-manager folders create ...
gcloud resource-manager folders list ...
gcloud resource-manager liens list ...
gcloud resource-manager org-policies list ...
```

## Policy Inheritance

IAM policies are inherited downward through the hierarchy:

```
Organization policy (applies to all)
  └── Folder policy (inherited + folder-specific)
       └── Project policy (inherited + project-specific)
            └── Resource policy (inherited + resource-specific)
```

Effective policy for a resource = union of all policies in its ancestry chain.

## References

- [Cloud Resource Hierarchy](https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy)
- [gcloud projects reference](https://cloud.google.com/sdk/gcloud/reference/projects)
- [gcloud organizations reference](https://cloud.google.com/sdk/gcloud/reference/organizations)

## Related Skills

- For billing-related flags (`--billing-project`), see [global-flags.md](global-flags.md).
- For project-related config properties (`core/project`), see [config-properties.md](config-properties.md).
- For the project ID vs name vs number gotcha, see [gotchas.md](gotchas.md).
