package tasks

/*
import (
	"fmt"
	"github.com/fyuan1316/flagger-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/flagger-operator/pkg/oprlib/processor/resource"
	resource2 "github.com/fyuan1316/flagger-operator/pkg/oprlib/resource"
	"github.com/fyuan1316/flagger-operator/pkg/task"
)

type ProvisionCrdsTask struct {
	resource.FileTask
}

func (p ProvisionCrdsTask) Name() string {
	return task.StageProvision
}

func (p ProvisionCrdsTask) Run(ctx *model.OperatorContext) error {
	fmt.Println("ProvisionCrdsTask Run")
	err := p.Sync(ctx)
	return err
}

var ProvisionCrds ProvisionCrdsTask

var ClusterAsmCrdDir = "files/provision/crds"

func SetUpCrds() {
	ProvisionCrds = ProvisionCrdsTask{
		FileTask: resource.FileTask{},
	}
	files, err := resource2.GetFilesInFolder(ClusterAsmCrdDir, resource2.Suffix(".yaml"))
	if err != nil {
		panic(err)
	}

	for path := range files {
		err := ProvisionCrds.LoadFile(
			path,
		)
		if err != nil {
			panic(err)
		}
	}
}

var _ model.ExecuteItem = ProvisionCrdsTask{}


*/
