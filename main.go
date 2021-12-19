package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"time"

// 	"sample-scheduler-framework/pkg/plugins"

// 	"k8s.io/component-base/logs"
// 	"k8s.io/kubernetes/cmd/kube-scheduler/app"
// )

// func main() {
// 	rand.Seed(time.Now().UTC().UnixNano())

// 	command := app.NewSchedulerCommand(
// 		app.WithPlugin(plugins.Name, plugins.New),
// 	)
// 	logs.InitLogs()
// 	defer logs.FlushLogs()

// 	if err := command.Execute(); err != nil {
// 		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
// 		os.Exit(1)
// 	}
// }

import (
	"fmt"
	"sample-scheduler-framework/pkg/info"
)

func main() {
	var i info.Result
	r, err := i.Init()
	if err != nil {
		return
	}
	for instance := range r.Info {
		fmt.Println(instance, r.Info[instance].Cpu)
	}
}
