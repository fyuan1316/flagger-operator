package entry

import (
	"github.com/fyuan1316/flagger-operator/pkg/task/provision"
	provisiontasks "github.com/fyuan1316/flagger-operator/pkg/task/provision/tasks"
	"github.com/fyuan1316/operatorlib/manage/model"
)

func GetOperatorStages() ([][]model.ExecuteItem, [][]model.ExecuteItem) {
	return getDeployStages(), getDeleteStages()
}

func getDeployStages() [][]model.ExecuteItem {
	//stages contains migration and deploy
	//stages := append(migration.GetStages(), provision.GetStages()...)

	stages := provision.GetStages()

	return stages
}

func getDeleteStages() [][]model.ExecuteItem {
	return provision.GetDeleteStages()
}

// 初始化任务数据
func SetUp() error {
	//  在这里定义部署任务和删除任务
	// 一般情况下部署任务可以包含两类，即migration 和 provision，而删除任务只有一类。（也可以根据自己的需要定义）

	// 1 migrations
	// migrationtasks.SetUpMigShell()
	// 2 provisions
	provisiontasks.SetUpResource()

	// 3 deletion
	provisiontasks.SetUpDeletion()
	return nil
}
