---
title: Custom Resource Annotations in Helm-based Operators
linkTitle: Custom Resource Annotations
weight: 300
description: Use custom resource annotations to configure how reconciliation works.
---

## `helm.sdk.operatorframework.io/upgrade-force`

This annotation can be set to `"true"` on custom resources to enable the chart to be upgraded with the
`helm upgrade --force` option. For more info see the [Helm Upgrade documentation](https://helm.sh/docs/helm/helm_upgrade/)
and this [explanation](https://github.com/helm/helm/issues/7082#issuecomment-559558318) of `--force` behavior.

**Example**

```yaml
apiVersion: example.com/v1alpha1
kind: Nginx
metadata:
  name: nginx-sample
  annotations:
    helm.sdk.operatorframework.io/upgrade-force: "true"
spec:
  replicaCount: 2
  service:
    port: 8080
```

Setting this annotation to `true` and making a change to trigger an upgrade (e.g. setting `spec.replicaCount: 3`)
will cause the custom resource to be reconciled and upgraded with the `force` option. This can be verified in the
log message when an upgrade succeeds:

```
{"level":"info","ts":1591198931.1703992,"logger":"helm.controller","msg":"Upgraded release","namespace":"helm-nginx","name":"example-nginx","apiVersion":"cache.example.com/v1alpha1","kind":"Nginx","release":"example-nginx","forceUpgrade":true}
```

## `helm.sdk.operatorframework.io/rollback-force`

The operator will assume this annotation to be `"true"`, if it is not set or set to an invalid value.
This annotation can be set to `"false"` to run a possible `helm rollback` after a failed release
without the `--force` option. For more info see the [Helm Rollback documentation](https://helm.sh/docs/helm/helm_rollback/).

Setting this can be required if a chart contains resources that have immutable fields after their first creation.
See [helm/helm#6378](https://github.com/helm/helm/issues/6378#issuecomment-735812673) for a more detailed description
of this issue.

This for example applies to a `Service` of type `ClusterIP`. The `ClusterIP` can be empty on creation, but kubernetes will fill
it with an actual IP, afterwards, the field is immutable. Helm (with `--force`) however will try to set it to `""`, which
results in an error.

**Example**

```yaml
apiVersion: example.com/v1alpha1
kind: Nginx
metadata:
  name: nginx-sample
  annotations:
    helm.sdk.operatorframework.io/rollback-force: "false"
spec:
  replicaCount: 2
  service:
    port: 8080
```
