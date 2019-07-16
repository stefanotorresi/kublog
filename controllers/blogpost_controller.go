package controllers

import (
	"context"

	"github.com/go-logr/logr"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	blogv1 "stefanotorresi/kublog/api/v1"
)

// BlogPostReconciler reconciles a BlogPost object
type BlogPostReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=blog.torresi.io,resources=blogposts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=blog.torresi.io,resources=blogposts/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=blog.torresi.io,resources=comments,verbs=get;list;watch;create;update;patch;delete

func (r *BlogPostReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("blogpost", req.NamespacedName)

	blogPost := &blogv1.BlogPost{}

	err := r.Get(ctx, req.NamespacedName, blogPost)
	if apierrs.IsNotFound(err) {
		log.Info("resource gone")
		return ctrl.Result{}, nil
	}
	if err != nil {
		log.Error(err, "could not get resource")
		return ctrl.Result{}, err
	}

	numComments, err := r.countComments(ctx, log, blogPost)
	if err != nil {
		log.Error(err, "unable to count comments")
		return ctrl.Result{}, err
	}
	if numComments == blogPost.Status.CommentCount {
		return ctrl.Result{}, nil
	}

	blogPost.Status.CommentCount = numComments

	err = r.Status().Update(ctx, blogPost)
	if err != nil {
		log.Error(err, "unable to update status")
		return ctrl.Result{}, err
	}

	log.Info("comment count changed", "value", numComments)

	return ctrl.Result{}, nil
}

func (r *BlogPostReconciler) countComments(ctx context.Context, log logr.Logger, blogPost *blogv1.BlogPost) (numComments int, err error) {
	var commentList blogv1.CommentList

	listOptions := []client.ListOptionFunc{
		client.InNamespace(blogPost.Namespace),
		client.MatchingLabels(map[string]string{"blogpost": blogPost.Name}),
	}
	err = r.List(ctx, &commentList, listOptions...)
	if err != nil {
		log.Error(err, "unable to get comment list")
		return numComments, err
	}

	numComments = len(commentList.Items)

	return numComments, err
}

func (r *BlogPostReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&blogv1.BlogPost{}).
		Owns(&blogv1.Comment{}).
		Complete(r)
}
