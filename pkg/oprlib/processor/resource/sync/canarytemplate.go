package sync

/*
import (
	"context"
	depv1alpha1 "github.com/fyuan1316/asm-operator/api/dep/v1alpha1"
	"gitlab-ce.alauda.cn/asm/flagger-operator/pkg/oprlib/manage/model"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var GeneratorCanaryTemplate = func() model.Object {
	return &depv1alpha1.CanaryTemplate{}
}
var FnCanaryTemplate = func(client client.Client, object model.Object) error {
	deploy := depv1alpha1.CanaryTemplate{}
	err := client.Get(context.Background(),
		types.NamespacedName{Namespace: object.GetNamespace(), Name: object.GetName()},
		&deploy,
	)
	if err != nil {
		if errors.IsNotFound(err) {
			errCreate := client.Create(context.Background(), object)
			if errCreate != nil {
				return errCreate
			}
			return nil
		}
		return err
	}
	//update
	wanted := object.(*depv1alpha1.CanaryTemplate)
	if !equality.Semantic.DeepDerivative(wanted.Spec, deploy.Spec) {
		deploy.Spec = wanted.Spec
		if errUpd := client.Update(context.Background(), &deploy); errUpd != nil {
			return errUpd
		}
	}
	return nil
}
*/
