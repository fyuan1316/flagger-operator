package shell

import (
	"fmt"
	pkgerrors "github.com/pkg/errors"
	"gitlab-ce.alauda.cn/asm/flagger-operator/pkg/oprlib/manage/model"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Scripts struct {
	preCheck  []string
	postCheck []string

	preRun  []string
	postRun []string

	run []string
}

const NotAExitCode = -1

func (s Scripts) Execute(filePath string) (int, error) {
	cmd := exec.Command(filePath)
	outputs, outErr := cmd.CombinedOutput()
	if outErr != nil {
		wrappedErr := pkgerrors.Wrap(outErr, fmt.Sprintf("%s", outputs))
		if exitError, ok := outErr.(*exec.ExitError); ok {
			return exitError.ExitCode(), wrappedErr
		}
		return NotAExitCode, wrappedErr
	}

	return 0, nil
}
func (s *ScriptManager) ensureExecutable(filePath string) error {
	return os.Chmod(filePath, 0755)
}
func (s *ScriptManager) LoadFiles(files map[string]string) error {
	for path := range files {
		if err := s.Load(path); err != nil {
			return err
		}
	}
	return nil
}
func (s *ScriptManager) Load(filePath string) error {
	if err := s.ensureExecutable(filePath); err != nil {
		return err
	}
	lowercaseFilePath := strings.ToLower(filePath)
	_, fileName := filepath.Split(lowercaseFilePath)

	if strings.HasPrefix(fileName, "precheck") {
		s.preCheck = append(s.preCheck, filePath)
	} else if strings.HasPrefix(fileName, "postcheck") {
		s.postCheck = append(s.postCheck, filePath)
	} else if strings.HasPrefix(fileName, "prerun") {
		s.preRun = append(s.preRun, filePath)
	} else if strings.HasPrefix(fileName, "postrun") {
		s.postRun = append(s.postRun, filePath)
	} else {
		s.run = append(s.run, filePath)
	}
	return nil
}

type ScriptManager struct {
	Scripts
}

var _ model.ExecuteItem = ScriptManager{}
var _ model.PreCheck = ScriptManager{}
var _ model.PostCheck = ScriptManager{}
var _ model.PreRun = ScriptManager{}
var _ model.PostRun = ScriptManager{}

func (s ScriptManager) Name() string {
	panic("implement me")
}
func (s ScriptManager) runScripts(filePaths []string) error {
	for _, shFilePath := range filePaths {
		if exitCode, err := s.Execute(shFilePath); err != nil || (exitCode > 0 && exitCode <= 255) {
			return pkgerrors.Wrap(err, fmt.Sprintf("execute shell task error,file=%s", shFilePath))
		}
	}
	return nil
}
func (s ScriptManager) PreCheck(oCtx *model.OperatorContext) (bool, error) {
	if len(s.preCheck) == 0 {
		return true, nil
	}
	if err := s.runScripts(s.preCheck); err != nil {
		return false, err
	}
	return true, nil
}
func (s ScriptManager) PostCheck(oCtx *model.OperatorContext) (bool, error) {
	if len(s.postCheck) == 0 {
		return true, nil
	}
	if err := s.runScripts(s.postCheck); err != nil {
		return false, err
	}
	return true, nil
}

func (s ScriptManager) Run(oCtx *model.OperatorContext) error {
	if len(s.run) == 0 {
		return nil
	}
	return s.runScripts(s.run)
}
func (s ScriptManager) PreRun(oCtx *model.OperatorContext) error {
	if len(s.preRun) == 0 {
		return nil
	}
	return s.runScripts(s.preRun)
}
func (s ScriptManager) PostRun(oCtx *model.OperatorContext) error {
	if len(s.postRun) == 0 {
		return nil
	}
	return s.runScripts(s.postRun)
}
