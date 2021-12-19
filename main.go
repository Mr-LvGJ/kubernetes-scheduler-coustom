package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"sample-scheduler-framework/pkg/info"
	"sample-scheduler-framework/pkg/plugins"

	"k8s.io/component-base/logs"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

var i info.Result

func main() {
	res, err := i.Init()
	if err != nil {
		return
	}
	for nodeName := range res.Info {
		fmt.Println(nodeName, res.Info[nodeName])
	}

	rand.Seed(time.Now().UTC().UnixNano())

	command := app.NewSchedulerCommand(
		app.WithPlugin(plugins.Name, plugins.New),
	)
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
