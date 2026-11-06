# IaCConsole CLI — Configuration Management for OpenTofu and Terraform

[![GitHub stars](https://img.shields.io/github/stars/IaCConsole/iacconsole-cli)](https://github.com/IaCConsole/iacconsole-cli)
[![GitHub license](https://img.shields.io/github/license/IaCConsole/iacconsole-cli)](https://github.com/IaCConsole/iacconsole-cli/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/iacconsole/iacconsole-cli)](https://goreportcard.com/report/github.com/iacconsole/iacconsole-cli)
[![GitHub release](https://img.shields.io/github/v/release/IaCConsole/iacconsole-cli)](https://github.com/IaCConsole/iacconsole-cli/releases)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.26-blue)](https://golang.org/)
[![Downloads](https://img.shields.io/github/downloads/IaCConsole/iacconsole-cli/total)](https://github.com/IaCConsole/iacconsole-cli/releases)

IaCConsole CLI is a free, open-source configuration management tool for OpenTofu and Terraform. It is part of [IaCConsole](https://github.com/IaCConsole), a set of open-source tools for DevOps, DevSecOps, and FinOps engineers.

The CLI dynamically orchestrates OpenTofu or Terraform deployments. It separates infrastructure configuration from code so teams can reuse modules across environments (dev, staging, production) without duplication. Configuration can come from local files or from the free hosted IaCConsole CMDB and Internal Developer Portal at [app.iacconsole.com](https://app.iacconsole.com).

## Table of Contents

- [Key Features](#key-features)
- [MCP Integration](#mcp-integration)
- [Use Cases](#use-cases)
- [Related Open-Source Tools](#related-open-source-tools)
- [Installation](#installation)
- [Quick Start](#quick-start-locally-with-demo-configuration)
- [Jenkins CI Example](#quick-start-with-jenkins-and-demo-configuration)
- [Quick Start with AI Coding Assistants](#quick-start-with-ai-coding-assistants)
- [Usage](#usage)
- [Configuration Storage](#configuration-storage)
- [License](#license)

## Key Features

- Environment configuration is stored outside the Terraform/OpenTofu code.
- Terraform/OpenTofu code, called a "unit", stays generic enough to deploy the same resources with different configurations.
- `tfvars` and variables are generated automatically in a temporary folder alongside the selected Terraform code ("unit").
- Once the temporary folder is ready, the CLI runs `terraform` or `tofu` with the specified parameters.
- State files are kept separate for each environment, with remote state paths derived from the configured dimensions. Deployed configurations are stored in different `tfstate` files in remote storage such as S3 or GCS.

## MCP Integration

IaCConsole CLI supports the [Model Context Protocol (MCP)](https://modelcontextprotocol.io/), allowing AI coding assistants to interact directly with your infrastructure configuration. This enables AI agents to query your CMDB, understand your infrastructure layout, and assist with complex remediation tasks safely.

- **Context-aware**: Assistants can retrieve dimension data, unit configurations, and inventory status directly through the protocol.
- **AI-Native Workflow**: Use your preferred AI assistant to plan, generate, and troubleshoot your infrastructure without leaving your editor.

## Use Cases

- Infrastructure as Code management across multiple environments
- CI/CD pipeline integration for automated deployments
- Configuration management for development, staging, and production
- Terraform/OpenTofu workflow automation
- Multi-cloud infrastructure orchestration

## Related Open-Source Tools

IaCConsole publishes four free, open-source repositories:

- [prescaler](https://github.com/IaCConsole/prescaler) — a Kubernetes operator that pre-scales workloads from CPU usage so pods are ready before traffic arrives. Go.
- [iacconsole-cli](https://github.com/IaCConsole/iacconsole-cli) — this repository. Configuration management for OpenTofu and Terraform. Go, Apache-2.0.
- [go-entrypoint](https://github.com/IaCConsole/go-entrypoint) — an entrypoint binary that starts your app with secrets from Vault and forwards signals correctly. Go, Apache-2.0.
- [go-entrypoint-updater](https://github.com/IaCConsole/go-entrypoint-updater) — keeps go-entrypoint up to date on Kubernetes nodes so every pod uses the current build. Go.

The CLI also works with the free hosted console at [app.iacconsole.com](https://app.iacconsole.com), which provides a CMDB, an Internal Developer Portal, and AI assistants for DevOps, DevSecOps, and FinOps work.

## Installation

### Download Pre-built Binaries

1. Download the latest release from [GitHub Releases](https://github.com/IaCConsole/iacconsole-cli/releases) (version >= 0.5.0)
2. Extract the binary for your platform (Linux/macOS)
3. Make it executable: `chmod +x iacconsole-cli`
4. Move to your PATH: `sudo mv iacconsole-cli /usr/local/bin/`

### Build from Source

```bash
git clone https://github.com/IaCConsole/iacconsole-cli.git
cd iacconsole-cli
go build -o bin/iacconsole-cli .
```

### Prerequisites

- [OpenTofu](https://opentofu.org/docs/intro/install/) or [Terraform](https://developer.hashicorp.com/terraform/install)
- Go 1.26+ (for building from source)

## Quick start locally with demo configuration

1. Install [OpenTofu](https://opentofu.org/docs/intro/install/)
2. Execute to generate a simple demo configuration connected to the demo account in the [IaCConsole API (CMDB)](#configuration-management-database-cmdb---iacconsole-api):

```bash
iacconsole-cli init
```

To generate with [file-based inventory](#file-based-configuration-storage-inventory-files):

```bash
iacconsole-cli init --iacconsole-db=false
```

3. Follow the on-screen instructions

## Quick start with Jenkins and demo configuration

You will need any Kubernetes cluster (ideally a small temporary one such as Podman + Kind or minikube).
Check that the kubectl context points to this test cluster, not production.

For a full end-to-end example of using IaCConsole CLI in a CI/CD pipeline, see the pre-configured Jenkins deployment in [examples/jenkins/README.md](examples/jenkins/README.md).

## Quick start with AI Coding Assistants

Getting started with IaCConsole CLI is even easier using AI coding assistants:

1. Open the repository in your preferred editor with an AI assistant installed:

- GitHub Copilot in VS Code
- Cursor
- Claude/Anthropic
- WindSurf
- Cline

2. Ask questions like:

- "How do I set up IaCConsole for AWS resources?"
- "Help me create a new unit for a GCP instance"
- "How do I use the IaCConsole API (CMDB)?"
- "Show me how to pass environment variables to my terraform code"

These instructions provide context about:

- The architecture and key concepts of IaCConsole CLI
- How to work with units, dimensions, and inventory sources
- Integration with the IaCConsole API (CMDB) and IaCConsole Web UI for centralized configuration
- Project-specific conventions and workflows

For more complex tasks, you can ask the AI assistant to guide you step-by-step through creating configurations, setting up the backend, or troubleshooting deployment issues.

## What about alternative tools?

Yes, there are several robust Infrastructure as Code (IaC) orchestration tools available for Terraform and OpenTofu. IaCConsole CLI is designed to complement these by focusing on dynamic CMDB integration, AI-native workflows, and automated lifecycle management.

### Comparison Overview

| Feature | IaCConsole CLI | Terragrunt | Terramate | Atmos |
| :--- | :--- | :--- | :--- | :--- |
| **Primary Focus** | AI-Native / CMDB Integration | DRY / Module Orchestration | Code Gen / Orchestration | Stack / Orchestration |
| **CMDB/IDP Native** | Yes | No | No | No |
| **AI/MCP Integration** | Built-in | Requires external | Requires external | Requires external |
| **Architecture** | Dynamic / External | Hierarchical | Hierarchical | Hierarchical |

IaCConsole CLI is inspired by tools like [Cubtera](https://cubtera.com/), aiming to provide a lightweight, open-source orchestration layer that integrates directly with your development environment through the Model Context Protocol (MCP).

## Usage

Organization with AWS resources and state stored in S3

```bash
./iacconsole-cli exec --config examples/.iacconsolerc -o demo-org -d account:test-account -d datacenter:staging1 -u vpc -- init
./iacconsole-cli exec --config examples/.iacconsolerc -o demo-org -d account:test-account -d datacenter:staging1 -u vpc -- plan
./iacconsole-cli exec --config examples/.iacconsolerc -o demo-org -d account:test-account -d datacenter:staging1 -u vpc -- apply
```

Organization with Google Cloud resources and state stored in Google Cloud Storage

```bash
./iacconsole-cli exec --config examples/.iacconsolerc -o gcp-org -d account:free-tier -u free_instance -- init
./iacconsole-cli exec --config examples/.iacconsolerc -o gcp-org -d account:free-tier -u free_instance -- plan
./iacconsole-cli exec --config examples/.iacconsolerc -o gcp-org -d account:free-tier -u free_instance -- apply
```

- Everything after `--` will be passed as parameters to the `cmd_to_exec`
- `-c` = to remove temp dir after any `iacconsole-cli` execution (after `apply` or `destroy` and exit code=0 temp dir removed automatically)
- `-o` = name of the `organization` (subfolder in **Inventory**, **units** folders and in `.iacconsolerc` config section)
- `-d` = `dimension` to attach to tofu/terraform. You may specify as many `-d` pairs as you need!
- `-u` = name of the `unit` in the `units` folder

## Unit Manifest

The optional `unit_manifest.json` file in the `unit` folder provides options for IaCConsole CLI.

Currently it supports `dimensions`: a list of the required or expected dimensions from the **Inventory Store**.

[unit_manifest.json example](examples/units/demo-org/vpc/unit_manifest.json)

## Configuration Storage

### Configuration Management Database (CMDB) - IaCConsole API

IaCConsole API is an OpenAPI-powered Configuration Management Database with a web-based **IaCConsole Web UI** for centrally managing infrastructure configurations. The hosted API and web UI at [app.iacconsole.com](https://app.iacconsole.com) are free to use.

You can set the env variable `IACCONSOLE_TOKEN` to point to the IaCConsole API (CMDB).
You can generate an API Token in the **Settings > API Tokens** section of the [IaCConsole Web App](https://app.iacconsole.com).

```bash
export IACCONSOLE_TOKEN='iacc_your-token-uuid'
```

With the correct `IACCONSOLE_TOKEN`, the CLI will connect and receive all the required dimension data from the IaCConsole API (CMDB).
An additional parameter could be passed to iacconsole-cli `-w workspacename`. In general, `workspacename` is the branch name of the source repo where the dimension is stored. If the IaCConsole API (CMDB) does not find the dimension with the specified `workspacename`, it will try to return the dimension from the `master` workspace/branch!

**IaCConsole API (CMDB)** provides additional features for your CI/CD pipelines:

- **IaCConsole Web UI** - Visual console for managing configurations at [https://app.iacconsole.com](https://app.iacconsole.com)
- **OpenAPI RESTful API** - Programmatic access to configuration data with full API documentation
- **CI/CD Integration** - Fetch configurations directly in pipelines (e.g., [first-app.json](examples/inventory/demo-org/application/first-app.json) for validation)
- **Dynamic Dropdowns** - Get dimension lists for [Jenkins drop-downs](examples/jenkins/README.md) to select deployment targets

**API Documentation:**
[IaCConsole API guide - endpoints and examples](https://iacconsole.com/docs/api)

**OpenAPI Reference:**
[Swagger API docs - full RESTful API documentation and examples](https://app.swaggerhub.com/apis-docs/altuhovsu/iacconsole-api/1.0.0)

To upload/update dimensions to the IaCConsole API from your Inventory Files repo you can use [inventory-to-iacconsole-db.sh script example](examples/inventory-to-iacconsole-db.sh) and execute it like `bash examples/inventory-to-iacconsole-db.sh examples/inventory/`

### File-based Configuration Storage (Inventory Files)

If the env variable `IACCONSOLE_TOKEN` is not set, the CLI uses file-based configuration storage (typically a dedicated Git repository), specified by the path configured in `inventory_path`.

Examples:

- [staging1.json in Inventory Files](examples/inventory/demo-org/datacenter/staging1.json)
- [dim_defaults.json in Inventory Files](examples/inventory/demo-org/datacenter/dim_defaults.json)

### Dimensions usage in tf-code

When you set dimensions with the CLI flags `-d datacenter:staging1`, IaCConsole CLI provides the following variables inside the Terraform/OpenTofu code:

- var.iacconsole_datacenter_name = will contain string `staging1`
- var.iacconsole_datacenter_data = will contain the whole object from `staging1.json`
- var.iacconsole_datacenter_defaults = will contain the whole object from `dim_defaults.json` IF the file `dim_defaults.json` exists!

Examples:

- [datacenter object with defaults used in tf-code](examples/units/demo-org/vpc/main.tf#L5)

## Passing environment variables from shell

For example, to pass a variable (such as an AWS region) from the shell to the Terraform code, set it and use it:

**Environment variables must start with the `iacconsole_envvar_` prefix.**

```bash
export iacconsole_envvar_awsregion=us-east-1
```

In the TF code:

```
provider "aws" {
    region = var.iacconsole_envvar_awsregion
}
```

[Env variables used in code example](examples/units/demo-org/vpc/providers.tf#L3)

## $HOME/.iacconsolerc

The config file path (YAML format) can be provided with the `--config` flag, for example:

```bash
iacconsole-cli --config path_to_config/.iacconsolerc exec -o demo-org -d account:test-account -d datacenter:staging1 -u vpc -- init
```

If the `--config` flag is not set, the CLI tries to load the file from the default location, `$HOME/.iacconsolerc`.

[.iacconsolerc example](examples/.iacconsolerc):

```yaml
defaults:
  units_path: examples/units
  shared_modules_path: examples/shared-modules
  inventory_path: examples/inventory
  cmd_to_exec: tofu
  backend:
    bucket: default-tfstates
    key: $iacconsole_state_path
    region: us-east-2
gcp-org:
  backend:
    bucket: gcp-tfstates
    prefix: $iacconsole_state_path
```

- `units_path` = relative path to the folder with Terraform code (`unit`)
- `shared_modules_path` = relative path to the folder with shared Terraform modules that any `unit` can use
- `inventory_path` = relative path to the folder with JSON inventory files
- `cmd_to_exec` = name of the binary to execute (`tofu` or `terraform`)
- `backend` = config values for the backend provider. All child key/value pairs are passed to `init`, and `$iacconsole_state_path` is replaced with the generated path.
  For example, when you execute `iacconsole-cli exec ... -- init`, IaCConsole CLI runs `init -backend-config=bucket=gcp-tfstates -backend-config=prefix=account_free-tier/free_instance.tfstate`.

At least the following must be set in the config file, with key/value pairs specific to the backend provider used by the organization:

```yaml
defaults:
  backend:
    bucket: default-tfstates
    key: $iacconsole_state_path
```

Other options have built-in defaults:

```yaml
defaults:
  inventory_path: "examples/inventory"
  shared_modules_path: ""
  units_path: "examples/units"
  cmd_to_exec: "tofu"
```

## Shared modules support

It is good practice to move generic Terraform code into modules and reuse those modules across multiple units.

The path to the folder with such private shared modules is configured by the `shared_modules_path` parameter in the `.iacconsolerc` configuration file.

This folder is mounted or linked into every temporary folder, so you can reference any module by a short path:

```
//use shared-module
module "vpc" {
  source = "./shared-modules/create_vpc"
}
```

Examples:

- [Shared module for VPC creation](examples/shared-modules/create_vpc)
- [Shared module for VPC creation used in code](examples/units/demo-org/vpc/main.tf#L3)

## Remote state (Terraform Backend where state data files are stored)

AWS, Google Cloud, and some other backends are supported. You can configure any backend provider in the [config file](#homeiacconsolerc).

Your Terraform code (`unit`) should contain at least the following for [AWS S3](examples/units/demo-org/vpc/versions.tf#L4):

```
terraform {
  backend "s3" {}
}
```

For [Google Cloud Storage](examples/units/gcp-org/free_instance/versions.tf#L4):

```
terraform {
  backend "gcs" {}
}
```

If for the `demo-org` config `bucket` is set, then `$iacconsole_state_path` will be like: `dimName1_dimValue1/dimNameN_dimValueN/unitName.tfstate`

If for the `demo-org` config `bucket` is NOT set, then `$iacconsole_state_path` will be like `org_demo-org/dimName1_dimValue1/dimNameN_dimValueN/unitName.tfstate`

This could be useful if you want to store by default tfstate for all the organizations in the same/default bucket `default-tfstates` but for some specific organization you need to store tfstates in a dedicated bucket `demo-org-tfstates`

## Data Source Configuration (data "terraform_remote_state")

To simplify data source configuration (`data "terraform_remote_state" "tfstate" { }`), it helps to have backend config values available as tfvars.

`var.iacconsole_backend_config` contains all parameters from the [`backend` section of the config](#homeiacconsolerc).

For example, for AWS S3 ([examples/units/demo-org/vpc/data.tf](examples/units/demo-org/vpc/data.tf)):

```
data "terraform_remote_state" "network" {
  backend = "s3"
  config = {
    bucket = var.iacconsole_backend_config.bucket
    key    = "network/terraform.tfstate"
    region = var.iacconsole_backend_config.region
  }
}
```

[And for GCS](examples/units/gcp-org/free_instance/data.tf):

```
data "terraform_remote_state" "free_instance" {
  backend = "gcs"
  config = {
    bucket  = var.iacconsole_backend_config.bucket
    prefix  = "account_free-tier/free_instance.tfstate"
  }
}
```

You set `key` or `prefix` to another unit's tfstate whose outputs you want to use.

## $HOME/.tofurc

We recommend enabling `plugin_cache_dir` to reuse providers.

[.tofurc example](examples/.tofurc):

```ini
plugin_cache_dir   = "$HOME/.terraform.d/plugin-cache"
plugin_cache_may_break_dependency_lock_file = true
```

Do not forget to create the plugin-cache dir: `mkdir "$HOME/.terraform.d/plugin-cache"`

## License

IaCConsole CLI is licensed with Apache License Version 2.0.
Please read the LICENSE file for more details.
