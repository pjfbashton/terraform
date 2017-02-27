package kubernetes

import (
	"github.com/hashicorp/terraform/helper/schema"
)

var namespaceSchema = map[string]*schema.Schema{
	"metadata": &schema.Schema{
		Type:        schema.TypeList,
		Description: "Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
		Required:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"annotations": &schema.Schema{
					Type:        schema.TypeMap,
					Description: "An unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
					Optional:    true,
				},
				"generate_name": &schema.Schema{
					Type:          schema.TypeString,
					Description:   "An optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.\n\nIf this field is specified and the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry (optionally after the time indicated in the Retry-After header).\n\nApplied only if Name is not specified. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#idempotency",
					Optional:      true,
					ForceNew:      true,
					ConflictsWith: []string{"metadata.name"},
				},
				"generation": &schema.Schema{
					Type:        schema.TypeInt,
					Description: "A sequence number representing a specific generation of the desired state. Populated by the system. Read-only.",
					Computed:    true,
				},
				"labels": &schema.Schema{
					Type:        schema.TypeMap,
					Description: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
					Optional:    true,
				},
				"name": &schema.Schema{
					Type:          schema.TypeString,
					Description:   "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
					Optional:      true,
					ForceNew:      true,
					Computed:      true,
					ConflictsWith: []string{"metadata.generate_name"},
				},
				"resource_version": &schema.Schema{
					Type:        schema.TypeString,
					Description: "An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.\n\nPopulated by the system. Read-only. Value must be treated as opaque by clients and . More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#concurrency-control-and-consistency",
					Computed:    true,
				},
				"self_link": &schema.Schema{
					Type:        schema.TypeString,
					Description: "A URL representing this object. Populated by the system. Read-only.",
					Computed:    true,
				},
				"uid": &schema.Schema{
					Type:        schema.TypeString,
					Description: "The unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.\n\nPopulated by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
					Computed:    true,
				},
			},
		},
	},
	"spec": &schema.Schema{
		Type:        schema.TypeList,
		Description: "Spec defines the behavior of the Namespace. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
		Optional:    true,
		ForceNew:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"finalizers": &schema.Schema{
					Type:        schema.TypeSet,
					Description: "Finalizers is an opaque list of values that must be empty to permanently remove object from storage. More info: http://releases.k8s.io/HEAD/docs/design/namespaces.md#finalizers",
					Optional:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
					Set:         schema.HashString,
				},
			},
		},
	},
}
