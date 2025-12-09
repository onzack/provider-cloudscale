# Provider Cloudscale

`provider-cloudscale` is a [Crossplane](https://crossplane.io/) provider
cloudscale that is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the [Cloudscale
API](https://www.cloudscale.ch/en/api/v1).

## Getting Started

This cloudscale serves as a starting point for generating a new [Crossplane Provider](https://docs.crossplane.io/latest/packages/providers/) using the [`upjet`](https://github.com/crossplane/upjet) tooling.

## Installation

### 1. Install the Provider

Install the provider into your Crossplane cluster:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-cloudscale
spec:
  package: ghcr.io/onzack/provider-cloudscale:v0.1.10
```

### 2. Create a Secret with your cloudscale.ch API Token

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: cloudscale-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "token": "your-cloudscale-api-token"
    }
```

### 3. Create a ProviderConfig

```yaml
apiVersion: cloudscale.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: cloudscale-creds
      namespace: crossplane-system
      key: credentials
```

You can obtain your API token from the [cloudscale.ch Control Panel](https://control.cloudscale.ch/).

## Examples

Example manifests for all supported resources are available in the repository:

- **[examples/](examples/)** – Handcrafted examples for ProviderConfig and common resources (Network, Subnet)
- **[examples-generated/](examples-generated/)** – Auto-generated examples for all managed resources:
  - Server, Volume, Network, Subnet
  - Floating IP, Custom Image
  - Load Balancer, Load Balancer Pool, Pool Member, Listener, Health Monitor
  - Server Group, Objects User

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/onzack/provider-cloudscale/issues).
