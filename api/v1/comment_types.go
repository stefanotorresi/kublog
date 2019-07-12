package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CommentSpec defines the desired state of Comment
type CommentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// CommentStatus defines the observed state of Comment
type CommentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// Comment is the Schema for the comments API
type Comment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CommentSpec   `json:"spec,omitempty"`
	Status CommentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CommentList contains a list of Comment
type CommentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Comment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Comment{}, &CommentList{})
}
