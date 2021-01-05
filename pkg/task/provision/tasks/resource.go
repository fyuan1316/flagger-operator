package tasks

import (
	"fmt"
	"gitlab-ce.alauda.cn/asm/flagger-operator/pkg/oprlib/manage/model"
	"gitlab-ce.alauda.cn/asm/flagger-operator/pkg/oprlib/processor/resource"
	resource2 "gitlab-ce.alauda.cn/asm/flagger-operator/pkg/oprlib/resource"
	"gitlab-ce.alauda.cn/asm/flagger-operator/pkg/task"
)

type ProvisionResourcesTask struct {
	*resource.ChartTask
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
		&resource.ChartTask{
			ChartDir: ClusterAsmResDir,
			//TemplateValues: data.GetDefaults(),
			// 增加自定义的mapping操作
			//ResourceMappings:

		},
	}
	ProvisionResources.Override(ProvisionResources)

	var (
		files map[string]string
		err   error
	)

	if files, err = resource2.GetChartResources(ClusterAsmResDir, nil); err != nil {
		panic(err)
	}
	if err = ProvisionResources.LoadResources(files); err != nil {
		panic(err)
	}
	fmt.Println()
}
