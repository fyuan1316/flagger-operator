package sync

import (
	"context"
	"gitlab-ce.alauda.cn/asm/flagger-operator/pkg/oprlib/manage/model"
	rbacv1beta1 "k8s.io/api/rbac/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var GeneratorClusterRoleBinding = func() model.Object {
	return &rbacv1beta1.ClusterRoleBinding{}
}
var FnClusterRoleBinding = func(client client.Client, object model.Object) error {
	deploy := rbacv1beta1.ClusterRoleBinding{}
	err := client.Get(context.Background(),
		types.NamespacedName{Name: object.GetName()},
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
	return nil
}
var GeneratorClusterRole = func() model.Object {
	return &rbacv1beta1.ClusterRole{}
}
var FnClusterRole = func(client client.Client, object model.Object) error {
	deploy := rbacv1beta1.ClusterRole{}
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
	return nil
}
var GeneratorRoleBinding = func() model.Object {
	return &rbacv1beta1.RoleBinding{}
}
var FnRoleBinding = func(client client.Client, object model.Object) error {
	deploy := rbacv1beta1.RoleBinding{}
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
	return nil
}
