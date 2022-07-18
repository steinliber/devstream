package argocdapp

import (
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller"
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller/kubectl"
)

func Read(options map[string]interface{}) (map[string]interface{}, error) {
	// 1. config install operations
	runner := &plugininstaller.Runner{
		PreExecuteOperations: []plugininstaller.MutableOperation{
			kubectl.Validate,
		},
		GetStatusOperation: kubectl.GetPluginDynamicState,
	}

	// 2. execute installer get status and error
	status, err := runner.Execute(plugininstaller.RawOptions(options))
	if err != nil {
		return nil, err
	}

	return status, nil

}
