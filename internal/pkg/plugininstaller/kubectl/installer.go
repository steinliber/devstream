package kubectl

import (
	"os"

	"github.com/devstream-io/devstream/internal/pkg/plugininstaller"
	"github.com/devstream-io/devstream/pkg/util/kubectl"
	"github.com/devstream-io/devstream/pkg/util/log"
)

const kubectlTempYAMLFile = "./app.yaml"

// InstallOrUpdate will install or update service by input options
func InstallOrUpdateWrapper(templateStr string) plugininstaller.BaseOperation {
	installWithTempFunc := func(options plugininstaller.RawOptions) error {
		opts, err := NewOptions(options)
		if err != nil {
			return err
		}

		// render an ArgoCD App YAML file based on inputs and template
		if err = writeContentToTmpFile(kubectlTempYAMLFile, templateStr, &opts); err != nil {
			return err
		}

		// kubectl apply -f
		if err = kubectl.KubeApply(kubectlTempYAMLFile); err != nil {
			return err
		}

		// remove temporary YAML file used for kubectl apply
		if err = os.Remove(kubectlTempYAMLFile); err != nil {
			log.Warnf("Temporary YAML file %s can't be deleted, but the installation is successful.", kubectlTempYAMLFile)
		}
		return err
	}
	return installWithTempFunc
}

// Delete will delete kubectl app from options
func DeleteWrapper(templateStr string) plugininstaller.BaseOperation {
	deleteWithTempFunc := func(options plugininstaller.RawOptions) error {
		opts, err := NewOptions(options)
		if err != nil {
			return err
		}

		// render an ArgoCD App YAML file based on inputs and template
		if err = writeContentToTmpFile(kubectlTempYAMLFile, templateStr, &opts); err != nil {
			return err
		}

		// kubectl delete -f
		if err = kubectl.KubeDelete(kubectlTempYAMLFile); err != nil {
			return err
		}

		// remove temporary YAML file used for kubectl apply
		if err = os.Remove(kubectlTempYAMLFile); err != nil {
			log.Warnf("Temporary YAML file %s can't be deleted, but the installation is successful.", kubectlTempYAMLFile)
		}
		return nil
	}
	return deleteWithTempFunc
}
