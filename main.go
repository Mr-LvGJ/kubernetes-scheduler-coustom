package main

import (
	"fmt"
	"math/rand"
	"os"
	"sample-scheduler-framework/pkg/register"
	"time"

	"k8s.io/component-base/logs"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	command := register.Register()
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
