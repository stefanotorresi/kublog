package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CommentUpvoteSpec defines the desired state of CommentUpvote
type CommentUpvoteSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// CommentUpvoteStatus defines the observed state of CommentUpvote
type CommentUpvoteStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// CommentUpvote is the Schema for the commentupvotes API
type CommentUpvote struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CommentUpvoteSpec   `json:"spec,omitempty"`
	Status CommentUpvoteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CommentUpvoteList contains a list of CommentUpvote
type CommentUpvoteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CommentUpvote `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CommentUpvote{}, &CommentUpvoteList{})
}
