package controllers

import (
	"context"

	"github.com/go-logr/logr"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	blogv1 "stefanotorresi/kublog/api/v1"
)

// CommentReconciler reconciles a Comment object
type CommentReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=blog.torresi.io,resources=comments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=blog.torresi.io,resources=comments/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=blog.torresi.io,resources=commentupvotes,verbs=get;list;watch;create;update;patch;delete

func (r *CommentReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("comment", req.NamespacedName)

	var comment blogv1.Comment

	err := r.Get(ctx, req.NamespacedName, &comment)

	if apierrs.IsNotFound(err) {
		log.Info("resource gone")
		return ctrl.Result{}, nil
	}

	if err != nil {
		log.Error(err, "could not get resource")
		return ctrl.Result{}, err
	}

	var blogPost blogv1.BlogPost
	key := types.NamespacedName{Namespace: req.Namespace, Name: comment.Spec.BlogPostName}
	err = r.Get(ctx, key, &blogPost)
	if err != nil {
		log.Error(err, "could not find blog post")
		return ctrl.Result{}, err
	}

	err = ctrl.SetControllerReference(&blogPost, &comment, r.Scheme)
	if err != nil {
		log.Error(err, "could set owner reference")
		return ctrl.Result{}, err
	}

	err = r.Update(ctx, &comment)
	if err != nil {
		log.Error(err, "unable to update comment")
		return ctrl.Result{}, err
	}

	var upvoteList blogv1.CommentUpvoteList

	err = r.List(
		ctx,
		&upvoteList,
		client.InNamespace(req.Namespace),
		client.MatchingField(".metadata.controller", req.Name))
	if err != nil {
		log.Error(err, "unable to get upvote list")
		return ctrl.Result{}, err
	}

	numUpvotes := len(upvoteList.Items)

	if numUpvotes == comment.Status.UpvoteCount {
		return ctrl.Result{}, nil
	}

	comment.Status.UpvoteCount = numUpvotes

	err = r.Status().Update(ctx, &comment)
	if err != nil {
		log.Error(err, "unable to update status")
		return ctrl.Result{}, err
	}

	log.Info("upvote count changed", "value", numUpvotes)

	return ctrl.Result{}, nil
}

func (r *CommentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	indexUpvotesByCommentName := func(rawObj runtime.Object) []string {
		comment := rawObj.(*blogv1.CommentUpvote)
		owner := metav1.GetControllerOf(comment)
		if owner == nil {
			return nil
		}
		if owner.APIVersion != blogv1.GroupVersion.String() || owner.Kind != "Comment" {
			return nil
		}
		return []string{owner.Name}
	}

	err := mgr.GetFieldIndexer().IndexField(&blogv1.CommentUpvote{}, ".metadata.controller", indexUpvotesByCommentName)
	if err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&blogv1.Comment{}).
		Owns(&blogv1.CommentUpvote{}).
		Complete(r)
}
