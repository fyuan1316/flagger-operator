package tasks

import (
	"fmt"
	"gitlab-ce.alauda.cn/asm/flagger-operator/pkg/oprlib/manage/model"
	"gitlab-ce.alauda.cn/asm/flagger-operator/pkg/oprlib/processor/resource"
	resource2 "gitlab-ce.alauda.cn/asm/flagger-operator/pkg/oprlib/resource"
	"gitlab-ce.alauda.cn/asm/flagger-operator/pkg/task"
)

type DeleteResourcesTask struct {
	*resource.ChartTask
}

var DeleteResources DeleteResourcesTask
var _ model.OverrideOperation = DeleteResourcesTask{}

func (p DeleteResourcesTask) GetOperation() model.OperationType {
	return model.Operations.Deletion
}

func (p DeleteResourcesTask) Name() string {
	return task.StageDeletion
}

//func (p DeleteResourcesTask) Run(om *model.OperatorManage) error {
//	fmt.Println("DeleteResourcesTask Run")
//	err := p.Delete(om)
//	return err
//}

func SetUpDeletion() {
	DeleteResources = DeleteResourcesTask{
		&resource.ChartTask{
			ChartDir: ClusterAsmResDir,
			FileTask: resource.FileTask{
				//加载任务values
				//TemplateValues: data.GetDefaults(),

				// 增加自定义的mapping操作
				//ResourceMappings:

				//设置任务对应k8s资源的生命周期
				KeepResourceAfterOperatorDeleted: resource.PointerFalse(),
			},
		},
	}
	//DeleteResources.implementor = DeleteResources
	DeleteResources.Override(DeleteResources)

	var (
		files map[string]string
		err   error
	)
	if files, err = resource2.GetChartResources(ClusterAsmResDir, nil); err != nil {
		panic(err)
	}
	if err = DeleteResources.LoadResources(files); err != nil {
		panic(err)
	}
	fmt.Println()

}
