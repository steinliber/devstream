package kubectl

import (
	"time"

	"github.com/cenkalti/backoff"

	"github.com/devstream-io/devstream/internal/pkg/plugininstaller"
)

// GetPluginStaticState will return status by static deploy config
func GetPluginStaticState(options plugininstaller.RawOptions) (map[string]interface{}, error) {
	opts, err := NewOptions(options)
	if err != nil {
		return nil, err
	}
	res := make(map[string]interface{})

	res["app"] = map[string]interface{}{
		"name":      opts.App.Name,
		"namespace": opts.App.Namespace,
	}

	res["src"] = map[string]interface{}{
		"repoURL":   opts.Source.RepoURL,
		"path":      opts.Source.Path,
		"valueFile": opts.Source.Valuefile,
	}

	res["dest"] = map[string]interface{}{
		"server":    opts.Destination.Server,
		"namespace": opts.Destination.Namespace,
	}

	return res, nil
}

// GetPluginDynamicState will return status by query kubernetes deployment
func GetPluginDynamicState(options plugininstaller.RawOptions) (map[string]interface{}, error) {
	opts, err := NewOptions(options)
	if err != nil {
		return nil, err
	}

	state := make(map[string]interface{})
	operation := func() error {
		err := getArgoCDAppFromK8sAndSetState(state, opts.App.Name, opts.App.Namespace)
		if err != nil {
			return err
		}
		return nil
	}
	bkoff := backoff.NewExponentialBackOff()
	bkoff.MaxElapsedTime = 3 * time.Minute
	err = backoff.Retry(operation, bkoff)
	if err != nil {
		return nil, err
	}
	return state, nil
}
