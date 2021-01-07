package provision

import (
	"github.com/fyuan1316/flagger-operator/pkg/task/provision/tasks"
	"github.com/fyuan1316/operatorlib/manage/model"
)

func GetStages() [][]model.ExecuteItem {
	return [][]model.ExecuteItem{
		{
			tasks.ProvisionResources,
		},
	}
}
func GetDeleteStages() [][]model.ExecuteItem {
	return [][]model.ExecuteItem{
		{
			tasks.DeleteResources,
		},
	}
}
