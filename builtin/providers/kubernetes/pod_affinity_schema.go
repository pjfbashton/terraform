package kubernetes

import (
	"github.com/hashicorp/terraform/helper/schema"
)

var podAffinitySchema = map[string]*schema.Schema{
	"node_affinity": &schema.Schema{
		Type:        schema.TypeList,
		Description: "Describes node affinity scheduling rules for the pod.",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"preferred_during_scheduling_ignored_during_execution": &schema.Schema{
					Type:        schema.TypeSet,
					Description: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding \"weight\" to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
					Optional:    true,
					Elem: &schema.Schema{Type: schema.TypeList,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"preference": &schema.Schema{
									Type:        schema.TypeList,
									Description: "A node selector term, associated with the corresponding weight.",
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"match_expressions": &schema.Schema{
												Type:        schema.TypeSet,
												Description: "Required. A list of node selector requirements. The requirements are ANDed.",
												Required:    true,
												Elem: &schema.Schema{Type: schema.TypeList,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": &schema.Schema{
																Type:        schema.TypeString,
																Description: "The label key that the selector applies to.",
																Optional:    true,
															},
															"operator": &schema.Schema{
																Type:        schema.TypeString,
																Description: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																Optional:    true,
															},
															"values": &schema.Schema{
																Type:        schema.TypeSet,
																Description: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																Optional:    true,
																Elem:        &schema.Schema{Type: schema.TypeString},
																Set:         schema.HashString,
															},
														},
													}},
											},
										},
									},
								},
								"weight": &schema.Schema{
									Type:        schema.TypeInt,
									Description: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
									Optional:    true,
								},
							},
						}},
				},
				"required_during_scheduling_ignored_during_execution": &schema.Schema{
					Type:        schema.TypeList,
					Description: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
					Optional:    true,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"node_selector_terms": &schema.Schema{
								Type:        schema.TypeSet,
								Description: "Required. A list of node selector terms. The terms are ORed.",
								Required:    true,
								Elem: &schema.Schema{Type: schema.TypeList,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"match_expressions": &schema.Schema{
												Type:        schema.TypeSet,
												Description: "Required. A list of node selector requirements. The requirements are ANDed.",
												Required:    true,
												Elem: &schema.Schema{Type: schema.TypeList,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": &schema.Schema{
																Type:        schema.TypeString,
																Description: "The label key that the selector applies to.",
																Optional:    true,
															},
															"operator": &schema.Schema{
																Type:        schema.TypeString,
																Description: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																Optional:    true,
															},
															"values": &schema.Schema{
																Type:        schema.TypeSet,
																Description: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																Optional:    true,
																Elem:        &schema.Schema{Type: schema.TypeString},
																Set:         schema.HashString,
															},
														},
													}},
											},
										},
									}},
							},
						},
					},
				},
			},
		},
	},
	"pod_affinity": &schema.Schema{
		Type:        schema.TypeList,
		Description: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"preferred_during_scheduling_ignored_during_execution": &schema.Schema{
					Type:        schema.TypeSet,
					Description: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding \"weight\" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
					Optional:    true,
					Elem: &schema.Schema{Type: schema.TypeList,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"pod_affinity_term": &schema.Schema{
									Type:        schema.TypeList,
									Description: "Required. A pod affinity term, associated with the corresponding weight.",
									Required:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"label_selector": &schema.Schema{
												Type:        schema.TypeList,
												Description: "A label query over a set of resources, in this case pods.",
												Optional:    true,
												MaxItems:    1,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"match_expressions": &schema.Schema{
															Type:        schema.TypeSet,
															Description: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															Optional:    true,
															Elem: &schema.Schema{Type: schema.TypeList,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"key": &schema.Schema{
																			Type:        schema.TypeString,
																			Description: "key is the label key that the selector applies to.",
																			Optional:    true,
																		},
																		"operator": &schema.Schema{
																			Type:        schema.TypeString,
																			Description: "operator represents a key's relationship to a set of values. Valid operators ard In, NotIn, Exists and DoesNotExist.",
																			Optional:    true,
																		},
																		"values": &schema.Schema{
																			Type:        schema.TypeSet,
																			Description: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																			Optional:    true,
																			Elem:        &schema.Schema{Type: schema.TypeString},
																			Set:         schema.HashString,
																		},
																	},
																}},
														},
														"match_labels": &schema.Schema{
															Type:        schema.TypeMap,
															Description: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.",
															Optional:    true,
														},
													},
												},
											},
											"namespaces": &schema.Schema{
												Type:        schema.TypeSet,
												Description: "namespaces specifies which namespaces the labelSelector applies to (matches against); nil list means \"this pod's namespace,\" empty list means \"all namespaces\" The json tag here is not \"omitempty\" since we need to distinguish nil and empty. See https://golang.org/pkg/encoding/json/#Marshal for more details.",
												Optional:    true,
												Elem:        &schema.Schema{Type: schema.TypeString},
												Set:         schema.HashString,
											},
											"topology_key": &schema.Schema{
												Type:        schema.TypeString,
												Description: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. For PreferredDuringScheduling pod anti-affinity, empty topologyKey is interpreted as \"all topologies\" (\"all topologies\" here means all the topologyKeys indicated by scheduler command-line argument --failure-domains); for affinity and for RequiredDuringScheduling pod anti-affinity, empty topologyKey is not allowed.",
												Optional:    true,
											},
										},
									},
								},
								"weight": &schema.Schema{
									Type:        schema.TypeInt,
									Description: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
									Optional:    true,
								},
							},
						}},
				},
			},
		},
	},
	"pod_anti_affinity": &schema.Schema{
		Type:        schema.TypeList,
		Description: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"preferred_during_scheduling_ignored_during_execution": &schema.Schema{
					Type:        schema.TypeSet,
					Description: "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding \"weight\" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
					Optional:    true,
					Elem: &schema.Schema{Type: schema.TypeList,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"pod_affinity_term": &schema.Schema{
									Type:        schema.TypeList,
									Description: "Required. A pod affinity term, associated with the corresponding weight.",
									Required:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"label_selector": &schema.Schema{
												Type:        schema.TypeList,
												Description: "A label query over a set of resources, in this case pods.",
												Optional:    true,
												MaxItems:    1,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"match_expressions": &schema.Schema{
															Type:        schema.TypeSet,
															Description: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															Optional:    true,
															Elem: &schema.Schema{Type: schema.TypeList,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"key": &schema.Schema{
																			Type:        schema.TypeString,
																			Description: "key is the label key that the selector applies to.",
																			Optional:    true,
																		},
																		"operator": &schema.Schema{
																			Type:        schema.TypeString,
																			Description: "operator represents a key's relationship to a set of values. Valid operators ard In, NotIn, Exists and DoesNotExist.",
																			Optional:    true,
																		},
																		"values": &schema.Schema{
																			Type:        schema.TypeSet,
																			Description: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																			Optional:    true,
																			Elem:        &schema.Schema{Type: schema.TypeString},
																			Set:         schema.HashString,
																		},
																	},
																}},
														},
														"match_labels": &schema.Schema{
															Type:        schema.TypeMap,
															Description: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.",
															Optional:    true,
														},
													},
												},
											},
											"namespaces": &schema.Schema{
												Type:        schema.TypeSet,
												Description: "namespaces specifies which namespaces the labelSelector applies to (matches against); nil list means \"this pod's namespace,\" empty list means \"all namespaces\" The json tag here is not \"omitempty\" since we need to distinguish nil and empty. See https://golang.org/pkg/encoding/json/#Marshal for more details.",
												Optional:    true,
												Elem:        &schema.Schema{Type: schema.TypeString},
												Set:         schema.HashString,
											},
											"topology_key": &schema.Schema{
												Type:        schema.TypeString,
												Description: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. For PreferredDuringScheduling pod anti-affinity, empty topologyKey is interpreted as \"all topologies\" (\"all topologies\" here means all the topologyKeys indicated by scheduler command-line argument --failure-domains); for affinity and for RequiredDuringScheduling pod anti-affinity, empty topologyKey is not allowed.",
												Optional:    true,
											},
										},
									},
								},
								"weight": &schema.Schema{
									Type:        schema.TypeInt,
									Description: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
									Optional:    true,
								},
							},
						},
					},
				},
			},
		},
	},
}
