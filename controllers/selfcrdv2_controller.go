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

	"fmt"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/go-logr/logr"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

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

	ctx := context.Background()
	log := r.Log.WithValues("selfcrdv2", req.NamespacedName)
	fmt.Printf("-------------req-1:%v-%v\n", req.Namespace, req.Name)

	// your logic here
	var selfcrdv2 corev1.SelfCRDV2
	if err := r.Get(ctx, req.NamespacedName, &selfcrdv2); err != nil {
		log.Error(err, "unable to fetch CronJob")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, ignoreNotFound(err)
	}
	fmt.Printf("-------------selfcrdv2-2:%v\n", selfcrdv2)

	var crds corev1.SelfCRDV2List
	if err := r.List(ctx, &crds, client.InNamespace(req.Namespace), client.MatchingFields{jobOwnerKey: req.Name}); err != nil {
		log.Error(err, "unable to list child Jobs")
		return ctrl.Result{}, err
	}
	fmt.Printf("-------------req-3:%v\n", req)

	for _, crd := range crds.Items {
		fmt.Printf("Username: %s, CustomID:%s\n", crd.Spec.Username, crd.Spec.CustomID)
	}

	return ctrl.Result{}, nil
}

var (
	jobOwnerKey = ".metadata.controller"
	apiGVStr    = corev1.GroupVersion.String()
)

func (r *SelfCRDV2Reconciler) SetupWithManager(mgr ctrl.Manager) error {

	if err := mgr.GetFieldIndexer().IndexField(&corev1.SelfCRDV2{}, jobOwnerKey, func(rawObj runtime.Object) []string {
		// grab the job object, extract the owner...
		job := rawObj.(*corev1.SelfCRDV2)
		owner := metav1.GetControllerOf(job)
		fmt.Println("----------------owner", owner)
		if owner == nil {
			return nil
		}
		// ...make sure it's a CronJob...
		if owner.APIVersion != apiGVStr || owner.Kind != "SelfCRDV2" {
			return nil
		}

		// ...and if so, return it
		return []string{owner.Name}
	}); err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.SelfCRDV2{}).
		Owns(&corev1.SelfCRDV2{}).
		Complete(r)
}

func ignoreNotFound(err error) error {
	if apierrs.IsNotFound(err) {
		return nil
	}
	return err
}
