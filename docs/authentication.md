---
title: Authentication
description: Azure credentials for Azure Service Operator.
---

# Authentication

ASO authenticates to Azure as a service principal or a managed identity. The credential can be
set at three levels, and the most specific one wins:

1. **Per resource** — the `serviceoperator.azure.com/credential-from` annotation, which the
   blueprints expose as the `credentialFrom` field.
2. **Per namespace** — a Secret named `aso-credential` in that namespace.
3. **Global** — the operator's own credential, cluster-wide.

Leaving `credentialFrom` empty inherits from the namespace, then the global credential, so a
single-tenant cluster needs nothing set per Composition.

Workload Identity is preferred over a client secret wherever the cluster supports it: it avoids
a long-lived credential entirely. Grant the identity only the roles the blueprints you publish
need — it is the identity every Composition provisions through.

## Subscription and resource group

Neither is a credential setting. `spec.owner` names the parent (usually a `ResourceGroup`), and
the subscription comes from the credential itself.
