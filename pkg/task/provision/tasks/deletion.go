package tasks

import (
	"github.com/fyuan1316/flagger-operator/pkg/task"
	"github.com/fyuan1316/operatorlib/manage/model"
	"github.com/fyuan1316/operatorlib/task/chart"
)

type DeleteResourcesTask struct {
	*chart.ChartTask
}

var DeleteResources DeleteResourcesTask
var _ model.OverrideOperation = DeleteResourcesTask{}

func (p DeleteResourcesTask) GetOperation() model.OperationType {
	return model.Operations.Deletion
}

func (p DeleteResourcesTask) Name() string {
	return task.StageDeletion
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
