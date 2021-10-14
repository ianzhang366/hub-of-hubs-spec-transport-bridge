package bundle

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	targetNamespace = "hub-of-hubs.open-cluster-management.io/remoteNamespace"
)

// NewPlacementBindingBundle creates a new placement binding bundle with no data in it.
func NewClusterLifecycleBundle() Bundle {
	return newBaseBundle(WithManipulate(putObjectInClcNamespace))
}

// make sure the object is put to the target namespace, specified in the annotation
func putObjectInClcNamespace(object metav1.Object) {
	a := object.GetAnnotations()
	if len(a) == 0 || a[targetNamespace] == "" {
		return
	}

	object.SetNamespace(a[targetNamespace])
}
