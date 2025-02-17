/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"io"
	"net/http"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	o2imsv1alpha1 "github.com/example/memcached-operator/api/v1alpha1"
)

// O2ImsReconciler reconciles a O2Ims object
type O2ImsReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=o2ims.example.com,resources=o2ims,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=o2ims.example.com,resources=o2ims/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=o2ims.example.com,resources=o2ims/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the O2Ims object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.18.4/pkg/reconcile
func (r *O2ImsReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	o2ims := &o2imsv1alpha1.O2Ims{}
	err := r.Get(ctx, req.NamespacedName, o2ims)
	if err != nil {
		if apierrors.IsNotFound(err) {
			log.Info("O2ims resource not found")
			return ctrl.Result{}, nil

		}

		log.Error(err, "Failed to get o2ims")
		return ctrl.Result{}, err
	}

	log.Info(o2ims.Spec.Endpoint)

	resp, err := http.Get(o2ims.Spec.Endpoint)
	if err != nil {
		log.Error(err, "Failed to GET endpoint "+o2ims.Spec.Endpoint)
		return ctrl.Result{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err, "Failed to read body")
		return ctrl.Result{}, err
	}

	log.Info(string(body))

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *O2ImsReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&o2imsv1alpha1.O2Ims{}).
		Complete(r)
}
