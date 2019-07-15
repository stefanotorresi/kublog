package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
// Important: Run "make" to regenerate code after modifying this file

// CommentSpec defines the desired state of Comment
type CommentSpec struct {
	BlogPost corev1.ObjectReference `json:"blogPost"`
	Text     string                 `json:"text"`
}

// CommentStatus defines the observed state of Comment
type CommentStatus struct {
	Upvotes     []corev1.ObjectReference `json:"upvotes"`
	UpvoteCount int32                    `json:"commentCount"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

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
