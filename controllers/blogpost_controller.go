package controllers

import (
	"context"

	"github.com/go-logr/logr"
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

func (r *BlogPostReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("blogpost", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *BlogPostReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&blogv1.BlogPost{}).
		Complete(r)
}
