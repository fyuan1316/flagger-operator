package tasks

import (
	"github.com/fyuan1316/flagger-operator/pkg/task"
	"github.com/fyuan1316/operatorlib/manage/model"
	"github.com/fyuan1316/operatorlib/task/chart"
)

type ProvisionResourcesTask struct {
	*chart.ChartTask
}

func (p ProvisionResourcesTask) GetName() string {
	return "flagger-controlplane-install"
}

func (p ProvisionResourcesTask) IsReady(oCtx *model.OperatorContext) bool {
	return true
}

func (p ProvisionResourcesTask) IsHealthy(oCtx *model.OperatorContext) bool {
	return false
}

var ProvisionResources ProvisionResourcesTask
var _ model.OverrideOperation = ProvisionResourcesTask{}
var _ model.HealthCheck = ProvisionResourcesTask{}

// 子类需要实现的接口，可以统一合并为一个大的接口定义。
func (p ProvisionResourcesTask) GetOperation() model.OperationType {
	return model.Operations.Provision
}

func (p ProvisionResourcesTask) Name() string {
	return task.StageProvision
}

var ClusterAsmResDir = "files/provision/flagger"

func SetUpResource() {
	ProvisionResources = ProvisionResourcesTask{
		&chart.ChartTask{
			Dir: ClusterAsmResDir,
		},
	}
	ProvisionResources.Init()
	ProvisionResources.Override(ProvisionResources)
}
