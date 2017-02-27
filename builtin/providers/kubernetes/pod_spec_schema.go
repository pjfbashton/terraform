package kubernetes

import (
	"github.com/hashicorp/terraform/helper/schema"
)

var podSpecSchema = map[string]*schema.Schema{
	"active_deadline_seconds": &schema.Schema{
		Type:        schema.TypeInt,
		Description: "Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.",
		Optional:    true,
	},
	"affinity": &schema.Schema{
		Type:        schema.TypeList,
		Description: "If specified, the pod's scheduling constraints",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: podAffinitySchema,
		},
	},
	"containers": &schema.Schema{
		Type:        schema.TypeSet,
		Description: "List of containers belonging to the pod. Containers cannot currently be added or removed. There must be at least one container in a Pod. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/containers",
		Required:    true,
		MinItems:    1,
		ForceNew:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"args": &schema.Schema{
					Type:        schema.TypeSet,
					Description: "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/containers#containers-and-commands",
					Optional:    true,
					ForceNew:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
					Set:         schema.HashString,
				},
				"command": &schema.Schema{
					Type:        schema.TypeSet,
					Description: "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/containers#containers-and-commands",
					Optional:    true,
					ForceNew:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
					Set:         schema.HashString,
				},
				"env": &schema.Schema{
					Type:        schema.TypeSet,
					Description: "List of environment variables to set in the container. Cannot be updated.",
					Optional:    true,
					ForceNew:    true,
					Elem: &schema.Schema{Type: schema.TypeList,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"name": &schema.Schema{
									Type:        schema.TypeString,
									Description: "Name of the environment variable. Must be a C_IDENTIFIER.",
									Optional:    true,
								},
								"value": &schema.Schema{
									Type:        schema.TypeString,
									Description: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to \"\".",
									Optional:    true,
								},
								"value_from": &schema.Schema{
									Type:        schema.TypeList,
									Description: "Source for the environment variable's value. Cannot be used if value is not empty.",
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"config_map_key_ref": &schema.Schema{
												Type:        schema.TypeList,
												Description: "Selects a key of a ConfigMap.",
												Optional:    true,
												MaxItems:    1,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"key": &schema.Schema{
															Type:        schema.TypeString,
															Description: "The key to select.",
															Optional:    true,
														},
														"local_object_reference": &schema.Schema{
															Type:        schema.TypeList,
															Description: "Selects a key from a ConfigMap.",
															Optional:    true,
															MaxItems:    1,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"name": &schema.Schema{
																		Type:        schema.TypeString,
																		Description: "Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
																		Optional:    true,
																	},
																},
															},
														},
														"optional": &schema.Schema{
															Type:        schema.TypeBool,
															Description: "Specify whether the ConfigMap or it's key must be defined",
															Optional:    true,
														},
													},
												},
											},
											"field_ref": &schema.Schema{
												Type:        schema.TypeList,
												Description: "Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.podIP.",
												Optional:    true,
												MaxItems:    1,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"api_version": &schema.Schema{
															Type:        schema.TypeString,
															Description: "Version of the schema the FieldPath is written in terms of, defaults to \"v1\".",
															Optional:    true,
														},
														"field_path": &schema.Schema{
															Type:        schema.TypeString,
															Description: "Path of the field to select in the specified API version.",
															Optional:    true,
														},
													},
												},
											},
											"resource_field_ref": &schema.Schema{
												Type:        schema.TypeList,
												Description: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
												Optional:    true,
												MaxItems:    1,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"container_name": &schema.Schema{
															Type:        schema.TypeString,
															Description: "Container name: required for volumes, optional for env vars",
															Optional:    true,
														},
														"divisor": &schema.Schema{
															Type:        schema.TypeString,
															Description: "Specifies the output format of the exposed resources, defaults to \"1\"",
															Optional:    true,
														},
														"resource": &schema.Schema{
															Type:        schema.TypeString,
															Description: "Required: resource to select",
															Optional:    true,
														},
													},
												},
											},
											"secret_key_ref": &schema.Schema{
												Type:        schema.TypeList,
												Description: "Selects a key of a secret in the pod's namespace",
												Optional:    true,
												MaxItems:    1,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"key": &schema.Schema{
															Type:        schema.TypeString,
															Description: "The key of the secret to select from.  Must be a valid secret key.",
															Optional:    true,
														},
														"local_object_reference": &schema.Schema{
															Type:        schema.TypeList,
															Description: "SecretKeySelector selects a key of a Secret.",
															Optional:    true,
															MaxItems:    1,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"name": &schema.Schema{
																		Type:        schema.TypeString,
																		Description: "Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
																		Optional:    true,
																	},
																},
															},
														},
														"optional": &schema.Schema{
															Type:        schema.TypeBool,
															Description: "Specify whether the Secret or it's key must be defined",
															Optional:    true,
														},
													},
												},
											},
										},
									},
								},
							},
						}},
				},
				"env_from": &schema.Schema{
					Type:        schema.TypeSet,
					Description: "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. An invalid key will prevent the container from starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.",
					Optional:    true,
					ForceNew:    true,
					Elem: &schema.Schema{Type: schema.TypeList,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"config_map_ref": &schema.Schema{
									Type:        schema.TypeList,
									Description: "The ConfigMap to select from",
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"local_object_reference": &schema.Schema{
												Type:        schema.TypeList,
												Description: "ConfigMapEnvSource selects a ConfigMap to populate the environment variables with.\n\nThe contents of the target ConfigMap's Data field will represent the key-value pairs as environment variables.",
												Optional:    true,
												MaxItems:    1,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"name": &schema.Schema{
															Type:        schema.TypeString,
															Description: "Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
															Optional:    true,
														},
													},
												},
											},
											"optional": &schema.Schema{
												Type:        schema.TypeBool,
												Description: "Specify whether the ConfigMap must be defined",
												Optional:    true,
											},
										},
									},
								},
								"prefix": &schema.Schema{
									Type:        schema.TypeString,
									Description: "An optional identifer to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
									Optional:    true,
								},
								"secret_ref": &schema.Schema{
									Type:        schema.TypeList,
									Description: "The Secret to select from",
									Optional:    true,
									MaxItems:    1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"local_object_reference": &schema.Schema{
												Type:        schema.TypeList,
												Description: "SecretEnvSource selects a Secret to populate the environment variables with.\n\nThe contents of the target Secret's Data field will represent the key-value pairs as environment variables.",
												Optional:    true,
												MaxItems:    1,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"name": &schema.Schema{
															Type:        schema.TypeString,
															Description: "Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
															Optional:    true,
														},
													},
												},
											},
											"optional": &schema.Schema{
												Type:        schema.TypeBool,
												Description: "Specify whether the Secret must be defined",
												Optional:    true,
											},
										},
									},
								},
							},
						}},
				},
				"image": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Docker image name. More info: http://kubernetes.io/docs/user-guide/images",
					Required:    true,
				},
				"image_pull_policy": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/images#updating-images",
					Optional:    true,
					ForceNew:    true,
				},
				"lifecycle": &schema.Schema{
					Type:        schema.TypeList,
					Description: "Actions that the management system should take in response to container lifecycle events. Cannot be updated.",
					Optional:    true,
					ForceNew:    true,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"post_start": &schema.Schema{
								Type:        schema.TypeList,
								Description: "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: http://kubernetes.io/docs/user-guide/container-environment#hook-details",
								Optional:    true,
								MaxItems:    1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"exec": &schema.Schema{
											Type:        schema.TypeList,
											Description: "One and only one of the following should be specified. Exec specifies the action to take.",
											Optional:    true,
											MaxItems:    1,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"command": &schema.Schema{
														Type:        schema.TypeSet,
														Description: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
														Optional:    true,
														Elem:        &schema.Schema{Type: schema.TypeString},
														Set:         schema.HashString,
													},
												},
											},
										},
										"http_get": &schema.Schema{
											Type:        schema.TypeList,
											Description: "HTTPGet specifies the http request to perform.",
											Optional:    true,
											MaxItems:    1,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"host": &schema.Schema{
														Type:        schema.TypeString,
														Description: "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead.",
														Optional:    true,
													},
													"http_headers": &schema.Schema{
														Type:        schema.TypeSet,
														Description: "Custom headers to set in the request. HTTP allows repeated headers.",
														Optional:    true,
														Elem: &schema.Schema{Type: schema.TypeList,
															MaxItems: 1,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"name": &schema.Schema{
																		Type:        schema.TypeString,
																		Description: "The header field name",
																		Optional:    true,
																	},
																	"value": &schema.Schema{
																		Type:        schema.TypeString,
																		Description: "The header field value",
																		Optional:    true,
																	},
																},
															}},
													},
													"path": &schema.Schema{
														Type:        schema.TypeString,
														Description: "Path to access on the HTTP server.",
														Optional:    true,
													},
													"port": &schema.Schema{
														Type:        schema.TypeInt,
														Description: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
														Optional:    true,
													},
													"scheme": &schema.Schema{
														Type:        schema.TypeString,
														Description: "Scheme to use for connecting to the host. Defaults to HTTP.",
														Optional:    true,
													},
												},
											},
										},
										"tcp_socket": &schema.Schema{
											Type:        schema.TypeList,
											Description: "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported",
											Optional:    true,
											MaxItems:    1,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"port": &schema.Schema{
														Type:        schema.TypeInt,
														Description: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
														Optional:    true,
													},
												},
											},
										},
									},
								},
							},
							"pre_stop": &schema.Schema{
								Type:        schema.TypeList,
								Description: "PreStop is called immediately before a container is terminated. The container is terminated after the handler completes. The reason for termination is passed to the handler. Regardless of the outcome of the handler, the container is eventually terminated. Other management of the container blocks until the hook completes. More info: http://kubernetes.io/docs/user-guide/container-environment#hook-details",
								Optional:    true,
								MaxItems:    1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"exec": &schema.Schema{
											Type:        schema.TypeList,
											Description: "One and only one of the following should be specified. Exec specifies the action to take.",
											Optional:    true,
											MaxItems:    1,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"command": &schema.Schema{
														Type:        schema.TypeSet,
														Description: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
														Optional:    true,
														Elem:        &schema.Schema{Type: schema.TypeString},
														Set:         schema.HashString,
													},
												},
											},
										},
										"http_get": &schema.Schema{
											Type:        schema.TypeList,
											Description: "HTTPGet specifies the http request to perform.",
											Optional:    true,
											MaxItems:    1,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"host": &schema.Schema{
														Type:        schema.TypeString,
														Description: "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead.",
														Optional:    true,
													},
													"http_headers": &schema.Schema{
														Type:        schema.TypeSet,
														Description: "Custom headers to set in the request. HTTP allows repeated headers.",
														Optional:    true,
														Elem: &schema.Schema{Type: schema.TypeList,
															MaxItems: 1,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"name": &schema.Schema{
																		Type:        schema.TypeString,
																		Description: "The header field name",
																		Optional:    true,
																	},
																	"value": &schema.Schema{
																		Type:        schema.TypeString,
																		Description: "The header field value",
																		Optional:    true,
																	},
																},
															}},
													},
													"path": &schema.Schema{
														Type:        schema.TypeString,
														Description: "Path to access on the HTTP server.",
														Optional:    true,
													},
													"port": &schema.Schema{
														Type:        schema.TypeInt,
														Description: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
														Optional:    true,
													},
													"scheme": &schema.Schema{
														Type:        schema.TypeString,
														Description: "Scheme to use for connecting to the host. Defaults to HTTP.",
														Optional:    true,
													},
												},
											},
										},
										"tcp_socket": &schema.Schema{
											Type:        schema.TypeList,
											Description: "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported",
											Optional:    true,
											MaxItems:    1,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"port": &schema.Schema{
														Type:        schema.TypeInt,
														Description: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
														Optional:    true,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				"liveness_probe": &schema.Schema{
					Type:        schema.TypeList,
					Description: "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes",
					Optional:    true,
					ForceNew:    true,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"failure_threshold": &schema.Schema{
								Type:        schema.TypeInt,
								Description: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
								Optional:    true,
							},
							"handler": &schema.Schema{
								Type:        schema.TypeList,
								Description: "Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.",
								Optional:    true,
								MaxItems:    1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"exec": &schema.Schema{
											Type:        schema.TypeList,
											Description: "One and only one of the following should be specified. Exec specifies the action to take.",
											Optional:    true,
											MaxItems:    1,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"command": &schema.Schema{
														Type:        schema.TypeSet,
														Description: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
														Optional:    true,
														Elem:        &schema.Schema{Type: schema.TypeString},
														Set:         schema.HashString,
													},
												},
											},
										},
										"http_get": &schema.Schema{
											Type:        schema.TypeList,
											Description: "HTTPGet specifies the http request to perform.",
											Optional:    true,
											MaxItems:    1,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"host": &schema.Schema{
														Type:        schema.TypeString,
														Description: "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead.",
														Optional:    true,
													},
													"http_headers": &schema.Schema{
														Type:        schema.TypeSet,
														Description: "Custom headers to set in the request. HTTP allows repeated headers.",
														Optional:    true,
														Elem: &schema.Schema{Type: schema.TypeList,
															MaxItems: 1,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"name": &schema.Schema{
																		Type:        schema.TypeString,
																		Description: "The header field name",
																		Optional:    true,
																	},
																	"value": &schema.Schema{
																		Type:        schema.TypeString,
																		Description: "The header field value",
																		Optional:    true,
																	},
																},
															}},
													},
													"path": &schema.Schema{
														Type:        schema.TypeString,
														Description: "Path to access on the HTTP server.",
														Optional:    true,
													},
													"port": &schema.Schema{
														Type:        schema.TypeInt,
														Description: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
														Optional:    true,
													},
													"scheme": &schema.Schema{
														Type:        schema.TypeString,
														Description: "Scheme to use for connecting to the host. Defaults to HTTP.",
														Optional:    true,
													},
												},
											},
										},
										"tcp_socket": &schema.Schema{
											Type:        schema.TypeList,
											Description: "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported",
											Optional:    true,
											MaxItems:    1,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"port": &schema.Schema{
														Type:        schema.TypeInt,
														Description: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
														Optional:    true,
													},
												},
											},
										},
									},
								},
							},
							"initial_delay_seconds": &schema.Schema{
								Type:        schema.TypeInt,
								Description: "Number of seconds after the container has started before liveness probes are initiated. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes",
								Optional:    true,
							},
							"period_seconds": &schema.Schema{
								Type:        schema.TypeInt,
								Description: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
								Optional:    true,
							},
							"success_threshold": &schema.Schema{
								Type:        schema.TypeInt,
								Description: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
								Optional:    true,
							},
							"timeout_seconds": &schema.Schema{
								Type:        schema.TypeInt,
								Description: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes",
								Optional:    true,
							},
						},
					},
				},
				"name": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",
					Required:    true,
					ForceNew:    true,
				},
				"ports": &schema.Schema{
					Type:        schema.TypeSet,
					Description: "List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default \"0.0.0.0\" address inside a container will be accessible from the network. Cannot be updated.",
					Optional:    true,
					ForceNew:    true,
					Elem: &schema.Schema{Type: schema.TypeList,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"container_port": &schema.Schema{
									Type:        schema.TypeInt,
									Description: "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
									Optional:    true,
								},
								"host_ip": &schema.Schema{
									Type:        schema.TypeString,
									Description: "What host IP to bind the external port to.",
									Optional:    true,
								},
								"host_port": &schema.Schema{
									Type:        schema.TypeInt,
									Description: "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
									Optional:    true,
								},
								"name": &schema.Schema{
									Type:        schema.TypeString,
									Description: "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",
									Optional:    true,
								},
								"protocol": &schema.Schema{
									Type:        schema.TypeString,
									Description: "Protocol for port. Must be UDP or TCP. Defaults to \"TCP\".",
									Optional:    true,
								},
							},
						}},
				},
				"readiness_probe": &schema.Schema{
					Type:        schema.TypeList,
					Description: "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes",
					Optional:    true,
					ForceNew:    true,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"failure_threshold": &schema.Schema{
								Type:        schema.TypeInt,
								Description: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
								Optional:    true,
							},
							"handler": &schema.Schema{
								Type:        schema.TypeList,
								Description: "Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.",
								Optional:    true,
								MaxItems:    1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"exec": &schema.Schema{
											Type:        schema.TypeList,
											Description: "One and only one of the following should be specified. Exec specifies the action to take.",
											Optional:    true,
											MaxItems:    1,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"command": &schema.Schema{
														Type:        schema.TypeSet,
														Description: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
														Optional:    true,
														Elem:        &schema.Schema{Type: schema.TypeString},
														Set:         schema.HashString,
													},
												},
											},
										},
										"http_get": &schema.Schema{
											Type:        schema.TypeList,
											Description: "HTTPGet specifies the http request to perform.",
											Optional:    true,
											MaxItems:    1,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"host": &schema.Schema{
														Type:        schema.TypeString,
														Description: "Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead.",
														Optional:    true,
													},
													"http_headers": &schema.Schema{
														Type:        schema.TypeSet,
														Description: "Custom headers to set in the request. HTTP allows repeated headers.",
														Optional:    true,
														Elem: &schema.Schema{Type: schema.TypeList,
															MaxItems: 1,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"name": &schema.Schema{
																		Type:        schema.TypeString,
																		Description: "The header field name",
																		Optional:    true,
																	},
																	"value": &schema.Schema{
																		Type:        schema.TypeString,
																		Description: "The header field value",
																		Optional:    true,
																	},
																},
															}},
													},
													"path": &schema.Schema{
														Type:        schema.TypeString,
														Description: "Path to access on the HTTP server.",
														Optional:    true,
													},
													"port": &schema.Schema{
														Type:        schema.TypeInt,
														Description: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
														Optional:    true,
													},
													"scheme": &schema.Schema{
														Type:        schema.TypeString,
														Description: "Scheme to use for connecting to the host. Defaults to HTTP.",
														Optional:    true,
													},
												},
											},
										},
										"tcp_socket": &schema.Schema{
											Type:        schema.TypeList,
											Description: "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported",
											Optional:    true,
											MaxItems:    1,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"port": &schema.Schema{
														Type:        schema.TypeInt,
														Description: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
														Optional:    true,
													},
												},
											},
										},
									},
								},
							},
							"initial_delay_seconds": &schema.Schema{
								Type:        schema.TypeInt,
								Description: "Number of seconds after the container has started before liveness probes are initiated. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes",
								Optional:    true,
							},
							"period_seconds": &schema.Schema{
								Type:        schema.TypeInt,
								Description: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
								Optional:    true,
							},
							"success_threshold": &schema.Schema{
								Type:        schema.TypeInt,
								Description: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
								Optional:    true,
							},
							"timeout_seconds": &schema.Schema{
								Type:        schema.TypeInt,
								Description: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes",
								Optional:    true,
							},
						},
					},
				},
				"resources": &schema.Schema{
					Type:        schema.TypeList,
					Description: "Compute Resources required by this container. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#resources",
					Optional:    true,
					ForceNew:    true,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"limits": &schema.Schema{
								Type:        schema.TypeMap,
								Description: "Limits describes the maximum amount of compute resources allowed. More info: http://kubernetes.io/docs/user-guide/compute-resources/",
								Optional:    true,
							},
							"requests": &schema.Schema{
								Type:        schema.TypeMap,
								Description: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: http://kubernetes.io/docs/user-guide/compute-resources/",
								Optional:    true,
							},
						},
					},
				},
				"security_context": &schema.Schema{
					Type:        schema.TypeList,
					Description: "Security options the pod should run with. More info: http://releases.k8s.io/HEAD/docs/design/security_context.md",
					Optional:    true,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"capabilities": &schema.Schema{
								Type:        schema.TypeList,
								Description: "The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime.",
								Optional:    true,
								MaxItems:    1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"add": &schema.Schema{
											Type:        schema.TypeSet,
											Description: "Added capabilities",
											Optional:    true,
											Elem:        &schema.Schema{Type: schema.TypeString},
											Set:         schema.HashString,
										},
										"drop": &schema.Schema{
											Type:        schema.TypeSet,
											Description: "Removed capabilities",
											Optional:    true,
											Elem:        &schema.Schema{Type: schema.TypeString},
											Set:         schema.HashString,
										},
									},
								},
							},
							"privileged": &schema.Schema{
								Type:        schema.TypeBool,
								Description: "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.",
								Optional:    true,
							},
							"read_only_root_filesystem": &schema.Schema{
								Type:        schema.TypeBool,
								Description: "Whether this container has a read-only root filesystem. Default is false.",
								Optional:    true,
							},
							"run_as_non_root": &schema.Schema{
								Type:        schema.TypeBool,
								Description: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
								Optional:    true,
							},
							"run_as_user": &schema.Schema{
								Type:        schema.TypeInt,
								Description: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
								Optional:    true,
							},
							"se_linux_options": &schema.Schema{
								Type:        schema.TypeList,
								Description: "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
								Optional:    true,
								MaxItems:    1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"level": &schema.Schema{
											Type:        schema.TypeString,
											Description: "Level is SELinux level label that applies to the container.",
											Optional:    true,
										},
										"role": &schema.Schema{
											Type:        schema.TypeString,
											Description: "Role is a SELinux role label that applies to the container.",
											Optional:    true,
										},
										"type": &schema.Schema{
											Type:        schema.TypeString,
											Description: "Type is a SELinux type label that applies to the container.",
											Optional:    true,
										},
										"user": &schema.Schema{
											Type:        schema.TypeString,
											Description: "User is a SELinux user label that applies to the container.",
											Optional:    true,
										},
									},
								},
							},
						},
					},
				},
				"stdin": &schema.Schema{
					Type:        schema.TypeBool,
					Description: "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",
					Optional:    true,
				},
				"stdin_once": &schema.Schema{
					Type:        schema.TypeBool,
					Description: "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",
					Optional:    true,
				},
				"termination_message_path": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",
					Optional:    true,
					ForceNew:    true,
				},
				"termination_message_policy": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",
					Optional:    true,
					ForceNew:    true,
				},
				"tty": &schema.Schema{
					Type:        schema.TypeBool,
					Description: "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",
					Optional:    true,
				},
				"volume_mounts": &schema.Schema{
					Type:        schema.TypeSet,
					Description: "Pod volumes to mount into the container's filesystem. Cannot be updated.",
					Optional:    true,
					ForceNew:    true,
					Elem: &schema.Schema{Type: schema.TypeList,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"mount_path": &schema.Schema{
									Type:        schema.TypeString,
									Description: "Path within the container at which the volume should be mounted.  Must not contain ':'.",
									Optional:    true,
								},
								"name": &schema.Schema{
									Type:        schema.TypeString,
									Description: "This must match the Name of a Volume.",
									Optional:    true,
								},
								"read_only": &schema.Schema{
									Type:        schema.TypeBool,
									Description: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
									Optional:    true,
								},
								"sub_path": &schema.Schema{
									Type:        schema.TypeString,
									Description: "Path within the volume from which the container's volume should be mounted. Defaults to \"\" (volume's root).",
									Optional:    true,
								},
							},
						}},
				},
				"working_dir": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",
					Optional:    true,
					ForceNew:    true,
				},
			},
		},
	},
	"dns_policy": &schema.Schema{
		Type:        schema.TypeString,
		Description: "Set DNS policy for containers within the pod. One of 'ClusterFirst' or 'Default'. Defaults to \"ClusterFirst\".",
		Optional:    true,
	},
	"host_ipc": &schema.Schema{
		Type:        schema.TypeBool,
		Description: "Use the host's ipc namespace. Optional: Default to false.",
		Optional:    true,
	},
	"host_network": &schema.Schema{
		Type:        schema.TypeBool,
		Description: "Host networking requested for this pod. Use the host's network namespace. If this option is set, the ports that will be used must be specified. Default to false.",
		Optional:    true,
	},
	"host_pid": &schema.Schema{
		Type:        schema.TypeBool,
		Description: "Use the host's pid namespace. Optional: Default to false.",
		Optional:    true,
	},
	"hostname": &schema.Schema{
		Type:        schema.TypeString,
		Description: "Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a system-defined value.",
		Optional:    true,
	},
	"image_pull_secrets": &schema.Schema{
		Type:        schema.TypeSet,
		Description: "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. For example, in the case of docker, only DockerConfig type secrets are honored. More info: http://kubernetes.io/docs/user-guide/images#specifying-imagepullsecrets-on-a-pod",
		Optional:    true,
		Elem: &schema.Schema{Type: schema.TypeList,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": &schema.Schema{
						Type:        schema.TypeString,
						Description: "Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
						Optional:    true,
					},
				},
			}},
	},
	"node_name": &schema.Schema{
		Type:        schema.TypeString,
		Description: "NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.",
		Optional:    true,
	},
	"node_selector": &schema.Schema{
		Type:        schema.TypeMap,
		Description: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: http://kubernetes.io/docs/user-guide/node-selection/README",
		Optional:    true,
	},
	"restart_policy": &schema.Schema{
		Type:        schema.TypeString,
		Description: "Restart policy for all containers within the pod. One of Always, OnFailure, Never. Default to Always. More info: http://kubernetes.io/docs/user-guide/pod-states#restartpolicy",
		Optional:    true,
	},
	"scheduler_name": &schema.Schema{
		Type:        schema.TypeString,
		Description: "If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler.",
		Optional:    true,
	},
	"security_context": &schema.Schema{
		Type:        schema.TypeList,
		Description: "SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty.  See type description for default values of each field.",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"fs_group": &schema.Schema{
					Type:        schema.TypeInt,
					Description: "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:\n\n1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw ",
					Optional:    true,
				},
				"run_as_non_root": &schema.Schema{
					Type:        schema.TypeBool,
					Description: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
					Optional:    true,
				},
				"run_as_user": &schema.Schema{
					Type:        schema.TypeInt,
					Description: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.",
					Optional:    true,
				},
				"se_linux_options": &schema.Schema{
					Type:        schema.TypeList,
					Description: "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.",
					Optional:    true,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"level": &schema.Schema{
								Type:        schema.TypeString,
								Description: "Level is SELinux level label that applies to the container.",
								Optional:    true,
							},
							"role": &schema.Schema{
								Type:        schema.TypeString,
								Description: "Role is a SELinux role label that applies to the container.",
								Optional:    true,
							},
							"type": &schema.Schema{
								Type:        schema.TypeString,
								Description: "Type is a SELinux type label that applies to the container.",
								Optional:    true,
							},
							"user": &schema.Schema{
								Type:        schema.TypeString,
								Description: "User is a SELinux user label that applies to the container.",
								Optional:    true,
							},
						},
					},
				},
				"supplemental_groups": &schema.Schema{
					Type:        schema.TypeSet,
					Description: "A list of groups applied to the first process run in each container, in addition to the container's primary GID.  If unspecified, no groups will be added to any container.",
					Optional:    true,
					Elem:        &schema.Schema{Type: schema.TypeInt},
				},
			},
		},
	},
	"service_account_name": &schema.Schema{
		Type:        schema.TypeString,
		Description: "ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: http://releases.k8s.io/HEAD/docs/design/service_accounts.md",
		Optional:    true,
	},
	"subdomain": &schema.Schema{
		Type:        schema.TypeString,
		Description: "If specified, the fully qualified Pod hostname will be \"<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>\". If not specified, the pod will not have a domainname at all.",
		Optional:    true,
	},
	"termination_grace_period_seconds": &schema.Schema{
		Type:        schema.TypeInt,
		Description: "Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. Defaults to 30 seconds.",
		Optional:    true,
	},
	"volumes": &schema.Schema{
		Type:        schema.TypeSet,
		Description: "List of volumes that can be mounted by containers belonging to the pod. More info: http://kubernetes.io/docs/user-guide/volumes",
		Optional:    true,
		Elem: &schema.Schema{
			Type:     schema.TypeList,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": &schema.Schema{
						Type:        schema.TypeString,
						Description: "Volume's name. Must be a DNS_LABEL and unique within the pod. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
						Optional:    true,
					},
					// "volume_source": &schema.Schema{
					// 	Type:     schema.TypeList,
					// 	MaxItems: 1,
					// 	Elem: &schema.Resource{
					// 		Schema: volumeSourceSchema,
					// 	},
					// },
				},
			},
		},
	},
}
