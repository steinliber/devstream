package kubectl

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/devstream-io/devstream/pkg/util/k8s"
)

func writeContentToTmpFile(file string, content string, opts *Options) error {
	t, err := template.New("app").Option("missingkey=error").Parse(content)
	if err != nil {
		return err
	}

	output, err := os.Create(file)
	if err != nil {
		return err
	}

	err = t.Execute(output, opts)
	if err != nil {
		if strings.Contains(err.Error(), "can't evaluate field name") {
			msg := err.Error()
			start := strings.Index(msg, "<")
			end := strings.Index(msg, ">")
			return fmt.Errorf("plugin argocdapp needs options%s but it's missing from the config file", msg[start+1:end])
		} else {
			return fmt.Errorf("executing tpl: %s", err)
		}
	}
	return nil
}

func getArgoCDAppFromK8sAndSetState(state map[string]interface{}, name, namespace string) error {
	kubeClient, err := k8s.NewClient()
	if err != nil {
		return err
	}

	app, err := kubeClient.GetArgocdApplication(namespace, name)
	if err != nil {
		return err
	}

	d := kubeClient.DescribeArgocdApp(app)
	state["app"] = d["app"]
	state["src"] = d["src"]
	state["dest"] = d["dest"]

	return nil
}
