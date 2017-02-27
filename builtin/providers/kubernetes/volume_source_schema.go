package kubernetes

import (
	"github.com/hashicorp/terraform/helper/schema"
)

var volumeSourceSchema = map[string]*schema.Schema{
	"aws_elastic_block_store": &schema.Schema{
		Type:        schema.TypeList,
		Description: "AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"fs_type": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore",
					Optional:    true,
				},
				"partition": &schema.Schema{
					Type:        schema.TypeInt,
					Description: "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as \"1\". Similarly, the volume partition for /dev/sda is \"0\" (or you can leave the property empty).",
					Optional:    true,
				},
				"read_only": &schema.Schema{
					Type:        schema.TypeBool,
					Description: "Specify \"true\" to force and set the ReadOnly property in VolumeMounts to \"true\". If omitted, the default is \"false\". More info: http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore",
					Optional:    true,
				},
				"volume_id": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore",
					Optional:    true,
				},
			},
		},
	},
	"azure_disk": &schema.Schema{
		Type:        schema.TypeList,
		Description: "AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"caching_mode": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Host Caching mode: None, Read Only, Read Write.",
					Optional:    true,
				},
				"data_disk_uri": &schema.Schema{
					Type:        schema.TypeString,
					Description: "The URI the data disk in the blob storage",
					Optional:    true,
				},
				"disk_name": &schema.Schema{
					Type:        schema.TypeString,
					Description: "The Name of the data disk in the blob storage",
					Optional:    true,
				},
				"fs_type": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.",
					Optional:    true,
				},
				"read_only": &schema.Schema{
					Type:        schema.TypeBool,
					Description: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
					Optional:    true,
				},
			},
		},
	},
	"azure_file": &schema.Schema{
		Type:        schema.TypeList,
		Description: "AzureFile represents an Azure File Service mount on the host and bind mount to the pod.",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"read_only": &schema.Schema{
					Type:        schema.TypeBool,
					Description: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
					Optional:    true,
				},
				"secret_name": &schema.Schema{
					Type:        schema.TypeString,
					Description: "the name of secret that contains Azure Storage Account Name and Key",
					Optional:    true,
				},
				"share_name": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Share Name",
					Optional:    true,
				},
			},
		},
	},
	"ceph_fs": &schema.Schema{
		Type:        schema.TypeList,
		Description: "CephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"monitors": &schema.Schema{
					Type:        schema.TypeSet,
					Description: "Required: Monitors is a collection of Ceph monitors More info: http://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it",
					Optional:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
					Set:         schema.HashString,
				},
				"path": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
					Optional:    true,
				},
				"read_only": &schema.Schema{
					Type:        schema.TypeBool,
					Description: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: http://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it",
					Optional:    true,
				},
				"secret_file": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: http://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it",
					Optional:    true,
				},
				"secret_ref": &schema.Schema{
					Type:        schema.TypeList,
					Description: "Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: http://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it",
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
				"user": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Optional: User is the rados user name, default is admin More info: http://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it",
					Optional:    true,
				},
			},
		},
	},
	"cinder": &schema.Schema{
		Type:        schema.TypeList,
		Description: "Cinder represents a cinder volume attached and mounted on kubelets host machine More info: http://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"fs_type": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: http://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md",
					Optional:    true,
				},
				"read_only": &schema.Schema{
					Type:        schema.TypeBool,
					Description: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: http://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md",
					Optional:    true,
				},
				"volume_id": &schema.Schema{
					Type:        schema.TypeString,
					Description: "volume id used to identify the volume in cinder More info: http://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md",
					Optional:    true,
				},
			},
		},
	},
	"config_map": &schema.Schema{
		Type:        schema.TypeList,
		Description: "ConfigMap represents a configMap that should populate this volume",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"default_mode": &schema.Schema{
					Type:        schema.TypeInt,
					Description: "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
					Optional:    true,
				},
				"items": &schema.Schema{
					Type:        schema.TypeSet,
					Description: "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
					Optional:    true,
					Elem: &schema.Schema{Type: schema.TypeList,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"key": &schema.Schema{
									Type:        schema.TypeString,
									Description: "The key to project.",
									Optional:    true,
								},
								"mode": &schema.Schema{
									Type:        schema.TypeInt,
									Description: "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
									Optional:    true,
								},
								"path": &schema.Schema{
									Type:        schema.TypeString,
									Description: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
									Optional:    true,
								},
							},
						}},
				},
				"local_object_reference": &schema.Schema{
					Type:        schema.TypeList,
					Description: "Adapts a ConfigMap into a volume.\n\nThe contents of the target ConfigMap's Data field will be presented in a volume as files using the keys in the Data field as the file names, unless the items element is populated with specific mappings of keys to paths. ConfigMap volumes support ownership management and SELinux relabeling.",
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
					Description: "Specify whether the ConfigMap or it's keys must be defined",
					Optional:    true,
				},
			},
		},
	},
	"downward_api": &schema.Schema{
		Type:        schema.TypeList,
		Description: "DownwardAPI represents downward API about the pod that should populate this volume",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"default_mode": &schema.Schema{
					Type:        schema.TypeInt,
					Description: "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
					Optional:    true,
				},
				"items": &schema.Schema{
					Type:        schema.TypeSet,
					Description: "Items is a list of downward API volume file",
					Optional:    true,
					Elem: &schema.Schema{Type: schema.TypeList,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"field_ref": &schema.Schema{
									Type:        schema.TypeList,
									Description: "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
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
								"mode": &schema.Schema{
									Type:        schema.TypeInt,
									Description: "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
									Optional:    true,
								},
								"path": &schema.Schema{
									Type:        schema.TypeString,
									Description: "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
									Optional:    true,
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
												Type:        schema.TypeList,
												Description: "Specifies the output format of the exposed resources, defaults to \"1\"",
												Optional:    true,
												MaxItems:    1,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"d": &schema.Schema{
															Type:     schema.TypeList,
															MaxItems: 1,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"dec": &schema.Schema{
																		Type:     schema.TypeList,
																		MaxItems: 1,
																		Elem: &schema.Resource{
																			Schema: map[string]*schema.Schema{
																				"scale": &schema.Schema{
																					Type: schema.TypeInt,
																				},
																				"unscaled": &schema.Schema{
																					Type:     schema.TypeList,
																					MaxItems: 1,
																					Elem: &schema.Resource{
																						Schema: map[string]*schema.Schema{
																							"neg": &schema.Schema{
																								Type: schema.TypeBool,
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
														"format": &schema.Schema{
															Type: schema.TypeString,
														},
														"i": &schema.Schema{
															Type:     schema.TypeList,
															MaxItems: 1,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"scale": &schema.Schema{
																		Type: schema.TypeInt,
																	},
																	"value": &schema.Schema{
																		Type: schema.TypeInt,
																	},
																},
															},
														},
														"s": &schema.Schema{
															Type: schema.TypeString,
														},
													},
												},
											},
											"resource": &schema.Schema{
												Type:        schema.TypeString,
												Description: "Required: resource to select",
												Optional:    true,
											},
										},
									},
								},
							},
						}},
				},
			},
		},
	},
	"empty_dir": &schema.Schema{
		Type:        schema.TypeList,
		Description: "EmptyDir represents a temporary directory that shares a pod's lifetime. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"medium": &schema.Schema{
					Type:        schema.TypeString,
					Description: "What type of storage medium should back this directory. The default is \"\" which means to use the node's default medium. Must be an empty string (default) or Memory. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",
					Optional:    true,
				},
			},
		},
	},
	"fc": &schema.Schema{
		Type:        schema.TypeList,
		Description: "FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"fs_type": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.",
					Optional:    true,
				},
				"lun": &schema.Schema{
					Type:        schema.TypeInt,
					Description: "Required: FC target lun number",
					Optional:    true,
				},
				"read_only": &schema.Schema{
					Type:        schema.TypeBool,
					Description: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
					Optional:    true,
				},
				"target_ww_ns": &schema.Schema{
					Type:        schema.TypeSet,
					Description: "Required: FC target worldwide names (WWNs)",
					Optional:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
					Set:         schema.HashString,
				},
			},
		},
	},
	"flex_volume": &schema.Schema{
		Type:        schema.TypeList,
		Description: "FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin. This is an alpha feature and may change in future.",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"driver": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Driver is the name of the driver to use for this volume.",
					Optional:    true,
				},
				"fs_type": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". The default filesystem depends on FlexVolume script.",
					Optional:    true,
				},
				"options": &schema.Schema{
					Type:        schema.TypeMap,
					Description: "Optional: Extra command options if any.",
					Optional:    true,
				},
				"read_only": &schema.Schema{
					Type:        schema.TypeBool,
					Description: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
					Optional:    true,
				},
				"secret_ref": &schema.Schema{
					Type:        schema.TypeList,
					Description: "Optional: SecretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
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
			},
		},
	},
	"flocker": &schema.Schema{
		Type:        schema.TypeList,
		Description: "Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"dataset_name": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
					Optional:    true,
				},
				"dataset_uuid": &schema.Schema{
					Type:        schema.TypeString,
					Description: "UUID of the dataset. This is unique identifier of a Flocker dataset",
					Optional:    true,
				},
			},
		},
	},
	"gce_persistent_disk": &schema.Schema{
		Type:        schema.TypeList,
		Description: "GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"fs_type": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk",
					Optional:    true,
				},
				"partition": &schema.Schema{
					Type:        schema.TypeInt,
					Description: "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as \"1\". Similarly, the volume partition for /dev/sda is \"0\" (or you can leave the property empty). More info: http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk",
					Optional:    true,
				},
				"pd_name": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk",
					Optional:    true,
				},
				"read_only": &schema.Schema{
					Type:        schema.TypeBool,
					Description: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk",
					Optional:    true,
				},
			},
		},
	},
	"git_repo": &schema.Schema{
		Type:        schema.TypeList,
		Description: "GitRepo represents a git repository at a particular revision.",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"directory": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
					Optional:    true,
				},
				"repository": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Repository URL",
					Optional:    true,
				},
				"revision": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Commit hash for the specified revision.",
					Optional:    true,
				},
			},
		},
	},
	"glusterfs": &schema.Schema{
		Type:        schema.TypeList,
		Description: "Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: http://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"endpoints_name": &schema.Schema{
					Type:        schema.TypeString,
					Description: "EndpointsName is the endpoint name that details Glusterfs topology. More info: http://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod",
					Optional:    true,
				},
				"path": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Path is the Glusterfs volume path. More info: http://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod",
					Optional:    true,
				},
				"read_only": &schema.Schema{
					Type:        schema.TypeBool,
					Description: "ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: http://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod",
					Optional:    true,
				},
			},
		},
	},
	"host_path": &schema.Schema{
		Type:        schema.TypeList,
		Description: "HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: http://kubernetes.io/docs/user-guide/volumes#hostpath",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"path": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Path of the directory on the host. More info: http://kubernetes.io/docs/user-guide/volumes#hostpath",
					Optional:    true,
				},
			},
		},
	},
	"iscsi": &schema.Schema{
		Type:        schema.TypeList,
		Description: "ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: http://releases.k8s.io/HEAD/examples/volumes/iscsi/README.md",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"fs_type": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: http://kubernetes.io/docs/user-guide/volumes#iscsi",
					Optional:    true,
				},
				"iqn": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Target iSCSI Qualified Name.",
					Optional:    true,
				},
				"iscsi_interface": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Optional: Defaults to 'default' (tcp). iSCSI interface name that uses an iSCSI transport.",
					Optional:    true,
				},
				"lun": &schema.Schema{
					Type:        schema.TypeInt,
					Description: "iSCSI target lun number.",
					Optional:    true,
				},
				"read_only": &schema.Schema{
					Type:        schema.TypeBool,
					Description: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
					Optional:    true,
				},
				"target_portal": &schema.Schema{
					Type:        schema.TypeString,
					Description: "iSCSI target portal. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
					Optional:    true,
				},
			},
		},
	},
	"nfs": &schema.Schema{
		Type:        schema.TypeList,
		Description: "NFS represents an NFS mount on the host that shares a pod's lifetime More info: http://kubernetes.io/docs/user-guide/volumes#nfs",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"path": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Path that is exported by the NFS server. More info: http://kubernetes.io/docs/user-guide/volumes#nfs",
					Optional:    true,
				},
				"read_only": &schema.Schema{
					Type:        schema.TypeBool,
					Description: "ReadOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: http://kubernetes.io/docs/user-guide/volumes#nfs",
					Optional:    true,
				},
				"server": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Server is the hostname or IP address of the NFS server. More info: http://kubernetes.io/docs/user-guide/volumes#nfs",
					Optional:    true,
				},
			},
		},
	},
	"persistent_volume_claim": &schema.Schema{
		Type:        schema.TypeList,
		Description: "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#persistentvolumeclaims",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"claim_name": &schema.Schema{
					Type:        schema.TypeString,
					Description: "ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#persistentvolumeclaims",
					Optional:    true,
				},
				"read_only": &schema.Schema{
					Type:        schema.TypeBool,
					Description: "Will force the ReadOnly setting in VolumeMounts. Default false.",
					Optional:    true,
				},
			},
		},
	},
	"photon_persistent_disk": &schema.Schema{
		Type:        schema.TypeList,
		Description: "PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"fs_type": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.",
					Optional:    true,
				},
				"pd_id": &schema.Schema{
					Type:        schema.TypeString,
					Description: "ID that identifies Photon Controller persistent disk",
					Optional:    true,
				},
			},
		},
	},
	"quobyte": &schema.Schema{
		Type:        schema.TypeList,
		Description: "Quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"group": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Group to map volume access to Default is no group",
					Optional:    true,
				},
				"read_only": &schema.Schema{
					Type:        schema.TypeBool,
					Description: "ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
					Optional:    true,
				},
				"registry": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
					Optional:    true,
				},
				"user": &schema.Schema{
					Type:        schema.TypeString,
					Description: "User to map volume access to Defaults to serivceaccount user",
					Optional:    true,
				},
				"volume": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Volume is a string that references an already created Quobyte volume by name.",
					Optional:    true,
				},
			},
		},
	},
	"rbd": &schema.Schema{
		Type:        schema.TypeList,
		Description: "RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: http://releases.k8s.io/HEAD/examples/volumes/rbd/README.md",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"ceph_monitors": &schema.Schema{
					Type:        schema.TypeSet,
					Description: "A collection of Ceph monitors. More info: http://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it",
					Optional:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
					Set:         schema.HashString,
				},
				"fs_type": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: http://kubernetes.io/docs/user-guide/volumes#rbd",
					Optional:    true,
				},
				"keyring": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: http://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it",
					Optional:    true,
				},
				"rados_user": &schema.Schema{
					Type:        schema.TypeString,
					Description: "The rados user name. Default is admin. More info: http://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it",
					Optional:    true,
				},
				"rbd_image": &schema.Schema{
					Type:        schema.TypeString,
					Description: "The rados image name. More info: http://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it",
					Optional:    true,
				},
				"rbd_pool": &schema.Schema{
					Type:        schema.TypeString,
					Description: "The rados pool name. Default is rbd. More info: http://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it.",
					Optional:    true,
				},
				"read_only": &schema.Schema{
					Type:        schema.TypeBool,
					Description: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: http://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it",
					Optional:    true,
				},
				"secret_ref": &schema.Schema{
					Type:        schema.TypeList,
					Description: "SecretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: http://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it",
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
			},
		},
	},
	"secret": &schema.Schema{
		Type:        schema.TypeList,
		Description: "Secret represents a secret that should populate this volume. More info: http://kubernetes.io/docs/user-guide/volumes#secrets",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"default_mode": &schema.Schema{
					Type:        schema.TypeInt,
					Description: "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
					Optional:    true,
				},
				"items": &schema.Schema{
					Type:        schema.TypeSet,
					Description: "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
					Optional:    true,
					Elem: &schema.Schema{Type: schema.TypeList,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"key": &schema.Schema{
									Type:        schema.TypeString,
									Description: "The key to project.",
									Optional:    true,
								},
								"mode": &schema.Schema{
									Type:        schema.TypeInt,
									Description: "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
									Optional:    true,
								},
								"path": &schema.Schema{
									Type:        schema.TypeString,
									Description: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
									Optional:    true,
								},
							},
						}},
				},
				"optional": &schema.Schema{
					Type:        schema.TypeBool,
					Description: "Specify whether the Secret or it's keys must be defined",
					Optional:    true,
				},
				"secret_name": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Name of the secret in the pod's namespace to use. More info: http://kubernetes.io/docs/user-guide/volumes#secrets",
					Optional:    true,
				},
			},
		},
	},
	"vsphere_volume": &schema.Schema{
		Type:        schema.TypeList,
		Description: "VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
		Optional:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"fs_type": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.",
					Optional:    true,
				},
				"volume_path": &schema.Schema{
					Type:        schema.TypeString,
					Description: "Path that identifies vSphere volume vmdk",
					Optional:    true,
				},
			},
		},
	},
}
