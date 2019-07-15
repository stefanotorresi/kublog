package controllers

import (
	"context"

	logrtesting "github.com/go-logr/logr/testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"

	blogv1 "stefanotorresi/kublog/api/v1"
)

// These tests are written in BDD-style using Ginkgo framework. Refer to
// http://onsi.github.io/ginkgo to learn more.

var _ = Describe("BlogPostReconciler", func() {

	var ctx context.Context
	var SUT *BlogPostReconciler

	BeforeEach(func() {
		ctx = context.Background()

		SUT = &BlogPostReconciler{
			Client: k8sClient,
			Log:    &logrtesting.NullLogger{},
		}
	})

	It("should update the comment count on a BlogPost", func() {

		By("creating a blogpost")
		blogPost := &blogv1.BlogPost{
			Spec: blogv1.BlogPostSpec{
				Title: "title",
				Body:  "body",
				Date:  metav1.Unix(0, 0),
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "blogpost",
				Namespace: "default",
			},
		}
		Expect(k8sClient.Create(ctx, blogPost)).To(Succeed())

		By("creating a comment")
		comment := &blogv1.Comment{
			Spec: blogv1.CommentSpec{
				Text: "text",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "comment",
				Namespace: blogPost.Namespace,
				Labels:    map[string]string{"blogpost": blogPost.Name},
			},
		}
		Expect(k8sClient.Create(ctx, comment)).To(Succeed())

		By("reconciling")
		key := types.NamespacedName{
			Name:      blogPost.Name,
			Namespace: blogPost.Namespace,
		}
		req := ctrl.Request{
			NamespacedName: key,
		}
		_, err := SUT.Reconcile(req)
		Expect(err).ToNot(HaveOccurred())

		reconciled := &blogv1.BlogPost{}
		Expect(k8sClient.Get(ctx, key, reconciled)).To(Succeed())
		Expect(reconciled.Status.CommentCount).To(Equal(1))

	})

})
