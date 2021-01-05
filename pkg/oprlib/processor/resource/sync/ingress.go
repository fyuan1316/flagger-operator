package sync

import (
	"context"
	"gitlab-ce.alauda.cn/asm/flagger-operator/pkg/oprlib/manage/model"
	extv1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var GeneratorIngress = func() model.Object {
	return &extv1beta1.Ingress{}
}
var FnIngress = func(client client.Client, object model.Object) error {
	deploy := extv1beta1.Ingress{}
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
	wanted := object.(*extv1beta1.Ingress)
	if !equality.Semantic.DeepDerivative(wanted.Spec, deploy.Spec) {
		deploy.Spec = wanted.Spec
		if errUpd := client.Update(context.Background(), &deploy); errUpd != nil {
			return errUpd
		}
	}
	return nil
}
