package controllers

import (
	"context"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	blogv1 "stefanotorresi/kublog/api/v1"
)

// CommentUpvoteReconciler reconciles a CommentUpvote object
type CommentUpvoteReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=blog.torresi.io,resources=commentupvotes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=blog.torresi.io,resources=commentupvotes/status,verbs=get;update;patch

func (r *CommentUpvoteReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("commentupvote", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *CommentUpvoteReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&blogv1.CommentUpvote{}).
		Complete(r)
}
