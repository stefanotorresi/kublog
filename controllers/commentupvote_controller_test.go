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

var _ = Describe("CommentReconciler", func() {

	var ctx context.Context
	var SUT *CommentUpvoteReconciler

	BeforeEach(func() {
		ctx = context.Background()

		SUT = &CommentUpvoteReconciler{
			Client: k8sClient,
			Log:    &logrtesting.NullLogger{},
			Scheme: testScheme,
		}
	})

	It("should register a Comment as owner reference on a CommentUpvote", func() {

		By("creating a blogpost")
		blogPost := &blogv1.BlogPost{
			Spec: blogv1.BlogPostSpec{
				Title: "title",
				Body:  "body",
				Date:  metav1.Unix(0, 0),
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "blogpost",
				Namespace: namespace,
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
				Namespace: namespace,
				Labels:    map[string]string{"blogpost": blogPost.Name},
			},
		}
		Expect(k8sClient.Create(ctx, comment)).To(Succeed())

		By("creating a comment upvote")
		commentUpvote := &blogv1.CommentUpvote{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "upvote",
				Namespace: namespace,
				Labels:    map[string]string{"comment": comment.Name},
			},
		}
		Expect(k8sClient.Create(ctx, commentUpvote)).To(Succeed())

		By("reconciling")
		key := types.NamespacedName{
			Name:      commentUpvote.Name,
			Namespace: namespace,
		}
		req := ctrl.Request{
			NamespacedName: key,
		}
		_, err := SUT.Reconcile(req)
		Expect(err).ToNot(HaveOccurred())

		reconciled := &blogv1.CommentUpvote{}
		Expect(k8sClient.Get(ctx, key, reconciled)).To(Succeed())
		Expect(reconciled.ObjectMeta.OwnerReferences[0].UID).To(Equal(comment.UID))

	})

})
