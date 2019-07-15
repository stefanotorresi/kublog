package controllers

import (
	"context"

	"github.com/go-logr/logr"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
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

	var blogPost blogv1.BlogPost

	err := r.Get(ctx, req.NamespacedName, &blogPost)

	if apierrs.IsNotFound(err) {
		log.Info("resource gone")
		return ctrl.Result{}, nil
	}

	if err != nil {
		log.Error(err, "could not get resource")
		return ctrl.Result{}, err
	}

	var commentList blogv1.CommentList

	err = r.List(
		ctx,
		&commentList,
		client.InNamespace(req.Namespace),
		client.MatchingField(".metadata.controller", req.Name))
	if err != nil {
		log.Error(err, "unable to get comment list")
		return ctrl.Result{}, err
	}

	numComments := len(commentList.Items)

	if numComments == blogPost.Status.CommentCount {
		return ctrl.Result{}, nil
	}

	blogPost.Status.CommentCount = numComments
	err = r.Status().Update(ctx, &blogPost)
	if err != nil {
		log.Error(err, "unable to update status")
		return ctrl.Result{}, err
	}

	log.Info("comment count changed", "value", numComments)

	return ctrl.Result{}, nil
}

func (r *BlogPostReconciler) SetupWithManager(mgr ctrl.Manager) error {
	indexCommentsByBlogPostName := func(rawObj runtime.Object) []string {
		comment := rawObj.(*blogv1.Comment)
		owner := metav1.GetControllerOf(comment)
		if owner == nil {
			return nil
		}
		if owner.APIVersion != blogv1.GroupVersion.String() || owner.Kind != "BlogPost" {
			return nil
		}
		return []string{owner.Name}
	}

	err := mgr.GetFieldIndexer().IndexField(&blogv1.Comment{}, ".metadata.controller", indexCommentsByBlogPostName)
	if err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&blogv1.BlogPost{}).
		Owns(&blogv1.Comment{}).
		Complete(r)
}
