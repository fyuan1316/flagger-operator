package provision

import (
	"gitlab-ce.alauda.cn/asm/flagger-operator/pkg/oprlib/manage/model"
	"gitlab-ce.alauda.cn/asm/flagger-operator/pkg/task/provision/tasks"
)

func GetStages() [][]model.ExecuteItem {
	return [][]model.ExecuteItem{
		{
			tasks.ProvisionCrds,
		},
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
