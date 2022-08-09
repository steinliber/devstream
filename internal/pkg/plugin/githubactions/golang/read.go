package golang

import (
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller"
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller/github"
)

func Read(options map[string]interface{}) (map[string]interface{}, error) {
	runner := &plugininstaller.Operator{
		PreExecuteOperations: plugininstaller.PreExecuteOperations{
			validate,
			github.BuildWorkFlowsWrapper(workflows),
		},
		GetStateOperation: github.GetActionState,
	}

	status, err := runner.Execute(plugininstaller.RawOptions(options))
	if err != nil {
		return nil, err
	}
	return status, nil
}
