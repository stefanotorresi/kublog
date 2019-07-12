package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BlogPostSpec defines the desired state of BlogPost
type BlogPostSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// BlogPostStatus defines the observed state of BlogPost
type BlogPostStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// BlogPost is the Schema for the blogposts API
type BlogPost struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BlogPostSpec   `json:"spec,omitempty"`
	Status BlogPostStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BlogPostList contains a list of BlogPost
type BlogPostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BlogPost `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BlogPost{}, &BlogPostList{})
}
