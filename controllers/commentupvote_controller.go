package controllers

import (
	"context"
	"errors"

	"github.com/go-logr/logr"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	blogv1 "stefanotorresi/kublog/api/v1"
)

// CommentUpvoteReconciler reconciles a CommentUpvote object
type CommentUpvoteReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=blog.torresi.io,resources=commentupvotes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=blog.torresi.io,resources=commentupvotes/status,verbs=get;update;patch

func (r *CommentUpvoteReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("comment", req.NamespacedName)

	var upvote blogv1.CommentUpvote

	err := r.Get(ctx, req.NamespacedName, &upvote)

	if apierrs.IsNotFound(err) {
		log.Info("resource gone")
		return ctrl.Result{}, nil
	}

	if err != nil {
		log.Error(err, "could not get resource")
		return ctrl.Result{}, err
	}

	err = r.setOwner(req, ctx, &upvote)
	if err != nil {
		log.Error(err, "unable to set comment owner")
		return ctrl.Result{}, err
	}

	err = r.Update(ctx, &upvote)
	if err != nil {
		log.Error(err, "unable to update upvote")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *CommentUpvoteReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&blogv1.CommentUpvote{}).
		Complete(r)
}

func (r *CommentUpvoteReconciler) setOwner(req ctrl.Request, ctx context.Context, upvote *blogv1.CommentUpvote) error {
	if _, ok := upvote.Labels["comment"]; ok != true {
		return errors.New("missing 'comment' label")
	}

	var comment blogv1.Comment
	key := types.NamespacedName{Namespace: req.Namespace, Name: upvote.Labels["comment"]}
	err := r.Get(ctx, key, &comment)
	if err != nil {
		return err
	}

	err = ctrl.SetControllerReference(&comment, upvote, r.Scheme)
	if err != nil {
		return err
	}

	return nil
}
