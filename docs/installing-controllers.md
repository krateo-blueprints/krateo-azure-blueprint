---
title: Installing Azure Service Operator
description: How Azure Service Operator is installed for the Krateo Azure blueprints.
---

# Installing Azure Service Operator

The blueprints in this repo render native ASO resources, so Azure Service Operator must be
running in the cluster before any Composition can be provisioned.

ASO publishes its own Helm chart and a single-YAML release bundle. Install the CRDs and the
operator, then configure credentials as described in
[`authentication.md`](authentication.md).

> **Planned:** ASO will ship here as a Krateo blueprint, the way Config Connector does in
> `krateo-gcp-blueprint`. Every dependency in Krateo is a blueprint; until that lands, ASO is a
> cluster-admin prerequisite.

## Which resources ASO can manage

ASO only reconciles the resource types whose CRDs you install. The release bundle ships all of
them; the Helm chart lets you select a subset with `crdPattern`. Install at least the CRDs for
the blueprints you intend to publish, plus `resources.azure.com` for `ResourceGroup`, which is
the `owner` of nearly everything else.
