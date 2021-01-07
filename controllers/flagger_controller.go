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
	entry "github.com/fyuan1316/flagger-operator/pkg/task/entry"
	"github.com/fyuan1316/operatorlib/manage"
	"github.com/fyuan1316/operatorlib/manage/model"
	pkgerrors "github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sync"

	flaagererrors "github.com/fyuan1316/flagger-operator/pkg/errors"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	operatorv1alpha1 "github.com/fyuan1316/flagger-operator/api/v1alpha1"
)

// FlaggerReconciler reconciles a Flagger object
type FlaggerReconciler struct {
	client.Client
	Log      logr.Logger
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

var once = sync.Once{}
var mgr *manage.OperatorManage
var (
	provisionTasks [][]model.ExecuteItem
	deletionTasks  [][]model.ExecuteItem
)

const finalizerID = "flagger.operator.alauda.io"

// +kubebuilder:rbac:groups=operator.alauda.io,resources=flaggers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=operator.alauda.io,resources=flaggers/status,verbs=get;update;patch

func (r *FlaggerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	var err error
	log := r.Log.WithValues("flagger", req.NamespacedName)
	log.Info(fmt.Sprintf("Starting reconcile loop for %v", req.NamespacedName))
	defer log.Info(fmt.Sprintf("Finish reconcile loop for %v", req.NamespacedName))
	instance := &operatorv1alpha1.Flagger{}
	err = r.Get(context.Background(), req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// CR not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}
	once.Do(func() {
		mgr = manage.NewOperatorManage(
			r.Client,
			manage.SetScheme(r.Scheme),
			manage.SetRecorder(r.Recorder),
			manage.SetFinalizer(finalizerID),
			manage.SetStatusUpdater(operatorStatusUpdater))

		provisionTasks, deletionTasks = entry.GetOperatorStages()
	})

	result, err := mgr.Reconcile(instance, provisionTasks, deletionTasks)
	if err != nil {
		log.Error(err, "Reconcile err")
		r.Recorder.Event(instance, flaagererrors.WarningEvent, flaagererrors.ReconcileError, err.Error())
		return result, err
	}

	return ctrl.Result{}, nil
}

var operatorStatusUpdater = func(obj runtime.Object, client client.Client) func(isReady, isHealthy bool) error {
	return func(isReady, isHealthy bool) error {
		var asm *operatorv1alpha1.Flagger
		var ok bool
		if asm, ok = obj.(*operatorv1alpha1.Flagger); !ok {
			return pkgerrors.New("operatorStatusUpdate cast model.Object to operatorv1alpha1.Flagger error")
		}
		asmCopy := asm.DeepCopy()
		asmCopy.Status.SetState(isReady, isHealthy)
		if asm.Status.State != asmCopy.Status.State {
			if updErr := client.Status().Update(context.Background(), asmCopy); updErr != nil {
				if errors.IsConflict(updErr) {
					cur := &operatorv1alpha1.Flagger{}
					if err := client.Get(
						context.Background(),
						types.NamespacedName{
							Namespace: asmCopy.GetNamespace(),
							Name:      asmCopy.GetName(),
						},
						cur,
					); err != nil {
						return err
					}
					retryObj := cur.DeepCopy()
					retryObj.Status.SetState(isReady, isHealthy)
					if updErr2 := client.Status().Update(context.Background(), retryObj); updErr2 != nil {
						return pkgerrors.Wrap(updErr2, "reUpdate FlaggerStatus error")
					}
					return nil
				}
				return pkgerrors.Wrap(updErr, "update FlaggerStatus error")
			}
		}
		return nil
	}

}

func (r *FlaggerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&operatorv1alpha1.Flagger{}).
		Complete(r)
}
