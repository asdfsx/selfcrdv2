/*

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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "github.com/asdfsx/selfcrdv2/api/v1"
)

// SelfCRDV2Reconciler reconciles a SelfCRDV2 object
type SelfCRDV2Reconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=core.clustar.ai,resources=selfcrdv2s,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core.clustar.ai,resources=selfcrdv2s/status,verbs=get;update;patch

func (r *SelfCRDV2Reconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("selfcrdv2", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *SelfCRDV2Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.SelfCRDV2{}).
		Complete(r)
}
