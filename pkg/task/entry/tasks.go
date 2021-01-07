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

	// for deploy stages
	// 1 migrations
	//migrationtasks.SetUpMigShell()
	// 2 provisions
	provisiontasks.SetUpResource()
	//-----------------
	//for delete stages
	provisiontasks.SetUpDeletion()
	return nil
}
