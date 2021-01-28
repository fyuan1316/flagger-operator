package tasks

import (
	"github.com/fyuan1316/operatorlib/manage/model"
	"github.com/fyuan1316/operatorlib/task/chart"
)

type DeleteResourcesTask struct {
	*chart.ChartTask
}

func (p DeleteResourcesTask) GetName() string {
	return "flagger-controlplane-delete"
}

var DeleteResources DeleteResourcesTask
var _ model.OverrideOperation = DeleteResourcesTask{}

func (p DeleteResourcesTask) GetOperation() model.OperationType {
	return model.Operations.Deletion
}

func SetUpDeletion() {
	DeleteResources = DeleteResourcesTask{
		&chart.ChartTask{
			Dir: ClusterAsmResDir,
		},
	}
	DeleteResources.Init()
	DeleteResources.Override(DeleteResources)

}
