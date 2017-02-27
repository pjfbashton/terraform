package kubernetes

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api/errors"
	api "k8s.io/kubernetes/pkg/api/v1"
	kubernetes "k8s.io/kubernetes/pkg/client/clientset_generated/release_1_5"
)

func resourceKubernetesPod() *schema.Resource {
	return &schema.Resource{
		Create: resourceKubernetesPodCreate,
		Read:   resourceKubernetesPodRead,
		Update: resourceKubernetesPodUpdate,
		Delete: resourceKubernetesPodDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: podSchema,
	}
}

func resourceKubernetesPodCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	metadata := expandMetadata(d.Get("metadata").([]interface{}))
	// spec := expandPodSpec(d.Get("spec").([]interface{}))
	pod := api.Pod{
		ObjectMeta: metadata,
		//Spec:       spec,
	}
	log.Printf("[INFO] Creating new pod: %#v", pod)
	out, err := conn.CoreV1().Pods(metadata.Namespace).Create(&pod)
	if err != nil {
		return fmt.Errorf("Pod creation failed: %s", err)
	}
	log.Printf("[INFO] Submitted new pod: %#v", out)
	d.SetId(out.Name)

	stateConf := &resource.StateChangeConf{
		Target:  []string{"Running", "Succeeded"},
		Pending: []string{"Unknown", "Failed", "Pending"},
		Timeout: 15 * time.Minute,
		Refresh: func() (interface{}, string, error) {
			out, err := conn.CoreV1().Pods(metadata.Namespace).Get(d.Id())
			if err != nil {
				return out, "Error", err
			}

			statusPhase := fmt.Sprintf("%v", out.Status.Phase)
			log.Printf("[DEBUG] Pod %s status received: %#v", out.Name, statusPhase)
			return out, statusPhase, nil
		},
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return err
	}

	return resourceKubernetesPodRead(d, meta)
}

func resourceKubernetesPodRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	metadata := expandMetadata(d.Get("metadata").([]interface{}))
	name := d.Id()
	log.Printf("[INFO] Reading pod: %#v", name)
	pod, err := conn.CoreV1().Pods(metadata.Namespace).Get(name)
	if err != nil {
		log.Printf("Received error: %#v", err)
		if statusErr, ok := err.(*errors.StatusError); ok && statusErr.ErrStatus.Code == 404 {
			log.Printf("[WARN] Removing pod %q (it is gone)", name)
			d.SetId("")
			return nil
		}
		return err
	}
	log.Printf("[INFO] Received pod: %#v", pod)
	d.Set("metadata", flattenMetadata(pod.ObjectMeta))

	return nil
}

func resourceKubernetesPodUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	metadata := expandMetadata(d.Get("metadata").([]interface{}))
	// This is necessary in case the name is generated
	metadata.Name = d.Id()
	//spec := expandPodSpec(d.Get("spec").([]interface{}))
	pod := api.Pod{
		ObjectMeta: metadata,
		//Spec:       spec,
	}
	log.Printf("[INFO] Updating pod: %#v", pod)
	out, err := conn.CoreV1().Pods(metadata.Namespace).Update(&pod)
	if err != nil {
		return fmt.Errorf("Pod update failed: %s", err)
	}
	log.Printf("[INFO] Submitted updated pod: %#v", out)

	stateConf := &resource.StateChangeConf{
		Target:  []string{"Running", "Succeeded"},
		Pending: []string{"Unknown", "Failed", "Pending"},
		Timeout: 15 * time.Minute,
		Refresh: func() (interface{}, string, error) {
			out, err := conn.CoreV1().Pods(metadata.Namespace).Get(d.Id())
			if err != nil {
				return out, "Error", err
			}

			statusPhase := fmt.Sprintf("%v", out.Status.Phase)
			log.Printf("[DEBUG] Pod %s status received: %#v", out.Name, statusPhase)
			return out, statusPhase, nil
		},
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return err
	}

	return resourceKubernetesPodRead(d, meta)
}

func resourceKubernetesPodDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*kubernetes.Clientset)

	metadata := expandMetadata(d.Get("metadata").([]interface{}))
	name := d.Id()

	log.Printf("[INFO] Deleting pod: %#v", name)
	err := conn.CoreV1().Pods(metadata.Namespace).Delete(name, &api.DeleteOptions{})
	if err != nil {
		return err
	}

	stateConf := &resource.StateChangeConf{
		Target:  []string{},
		Pending: []string{"Unknown", "Failed", "Running", "Succeeded"},
		Timeout: 15 * time.Minute,
		Refresh: func() (interface{}, string, error) {
			out, err := conn.CoreV1().Pods(metadata.Namespace).Get(name)
			if err != nil {
				if statusErr, ok := err.(*errors.StatusError); ok && statusErr.ErrStatus.Code == 404 {
					return nil, "", nil
				}
				log.Printf("[ERROR] Received error: %#v", err)
				return out, "Error", err
			}

			statusPhase := fmt.Sprintf("%v", out.Status.Phase)
			log.Printf("[DEBUG] Pod %s status received: %#v", out.Name, statusPhase)
			return out, statusPhase, nil
		},
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
