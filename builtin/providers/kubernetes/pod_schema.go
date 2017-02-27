package kubernetes

import (
	"github.com/hashicorp/terraform/helper/schema"
)

var podSchema = map[string]*schema.Schema{
	"metadata": metadataSchema,
	"spec": &schema.Schema{
		Type:        schema.TypeList,
		Description: "Specification of the desired behavior of the pod. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
		Required:    true,
		ForceNew:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: podSpecSchema,
		},
	},
}
