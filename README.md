# Automated Deployment for Apache Pulsar on Azure AKS

Bootstrap an Azure Kubernetes Service with an Apache Pulsar cluster on it.

Fully automated deployment, just run to deploy all:

```bash
./bin/task
```

[Taksfile](https://taskfile.dev) is used as a task runner which makes this deployment ci/cd platform independent

This is used to setup a minimal event-driven eco-system fast.

## Preconditions

### taskfile binary

The taskfile binary is mandatory for the deployment
Install the binary via script or see [use alternative setup instructions]

```bash
sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d
```

[use alternative setup instructions]: https://taskfile.dev/#/installation

### Azure authentication

The azurerm `Terraform` provider is used for deployment.

By default the authentication via [Azure CLI is used].

```bash
az login --tenant "my_tenant"
az account set --subscription "my_subscription"
```

The use of service principals credentials is also possible via [service principal auth].

```bash
    $ export ARM_CLIENT_ID="00000000-0000-0000-0000-000000000000"
    $ export ARM_SUBSCRIPTION_ID="00000000-0000-0000-0000-000000000000"
    $ export ARM_TENANT_ID="00000000-0000-0000-0000-000000000000"
```

[Azure CLI is used]: https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/guides/azure_cli
[service principal auth]: https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/guides/service_principal_client_certificate#configuring-the-service-principal-in-terraform

### Terraform

Terraform is used to deploy Azure resources so the binary needs to be [downloaded](https://www.terraform.io/downloads.html) first.

## Helm

Helm is used to deploy Pulsar on Kubernetes, so a [helm installation/binary](https://helm.sh/docs/intro/install/) is also needed.

## Go

An example Client is written in Go.

## Usage

To deploy everything run the configured default task.

```bash
./bin/task
```

## Deployment tasks

To run tasks by hand you can list all available tasks first.

```bash
./bin/task --list
```

To deploy Azure infrastructure use

```bash
./bin/task deploy:cloud-infra
```

To trigger Pulsar deployment with Helm

```bash
./bin/task deploy:pulsar
```
