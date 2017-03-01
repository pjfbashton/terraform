---
layout: "kubernetes"
page_title: "Kubernetes: kubernetes_namespace"
sidebar_current: "docs-kubernetes-namespace"
description: |-
  TODO
---

# kubernetes_namespace

TODO


## Example Usage

```
resource "kubernetes_namespace" "example" {
  // TODO
}
```

## Argument Reference

The following arguments are supported:

* `metadata` - (Required) Standard namespace's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata

## Nested Blocks

### `metadata`

#### Arguments


* `annotations` - (Optional) An unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. More info: http://kubernetes.io/docs/user-guide/annotations
* `generate_name` - (Optional) Prefix, used by the server, to generate a unique name ONLY IF the `name` field has not been provided. This value will also be combined with a unique suffix. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#idempotency
* `labels` - (Optional) Map of string keys and values that can be used to organize and categorize (scope and select) namespaces. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels
* `name` - (Optional) Name of the namespace, must be unique. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names

#### Attributes


* `generation` - A sequence number representing a specific generation of the desired state.
* `resource_version` - An opaque value that represents the internal version of this namespace that can be used by clients to determine when namespaces have changed. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#concurrency-control-and-consistency
* `self_link` - A URL representing this namespace.
* `uid` - The unique in time and space value for this namespace. More info: http://kubernetes.io/docs/user-guide/identifiers#uids

## Import

KMS Keys can be imported using the `id`, e.g.

```
$ terraform import aws_kms_key.a arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
```