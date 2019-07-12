package controllers

import (
	"context"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	blogv1 "stefanotorresi/kublog/api/v1"
)

// CommentReconciler reconciles a Comment object
type CommentReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=blog.torresi.io,resources=comments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=blog.torresi.io,resources=comments/status,verbs=get;update;patch

func (r *CommentReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("comment", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *CommentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&blogv1.Comment{}).
		Complete(r)
}
