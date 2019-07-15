package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
// Important: Run "make" to regenerate code after modifying this file

// CommentSpec defines the desired state of Comment
type CommentSpec struct {
	Text string `json:"text"`
}

// CommentStatus defines the observed state of Comment
type CommentStatus struct {
	UpvoteCount int `json:"upvoteCount"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Comment is the Schema for the comments API
// +kubebuilder:printcolumn:name="BlogPost",type="string",JSONPath=".metadata.owner", description="The name of the blog post"
// +kubebuilder:printcolumn:name="UpvotesCount",type="integer",JSONPath=".status.upvoteCount", description="The number of upvotes for this comment"
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
