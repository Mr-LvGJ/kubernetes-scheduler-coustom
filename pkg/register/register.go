package register

import (
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
	"sample-scheduler-framework/pkg/plugins"
)

func Register() *cobra.Command {
	return app.NewSchedulerCommand(
		app.WithPlugin(plugins.Name, plugins.New),
	)
}