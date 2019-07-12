package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
// Important: Run "make" to regenerate code after modifying this file

// BlogPostSpec defines the desired state of BlogPost
type BlogPostSpec struct {
	Title string `json:"title"`
	Body  string `json:"body"`

	// +optional
	Date *metav1.Time `json:"date,omitempty"`
}

// BlogPostStatus defines the observed state of BlogPost
type BlogPostStatus struct {
	Comments     []corev1.ObjectReference `json:"comments"`
	CommentCount int32                    `json:"commentCount"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// BlogPost is the Schema for the blogposts API
// +kubebuilder:printcolumn:name="Title",type="string",JSONPath=".spec.title",description="The title of the blog post"
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
