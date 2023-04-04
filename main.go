package main

import (
	"netflix_conductor/task"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/conductor-sdk/conductor-go/sdk/worker"
	"github.com/sirupsen/logrus"
)

var (
	apiClient = client.NewAPIClient(nil,
		settings.NewHttpSettings("http://localhost:8080/api/"))
	taskRunner = worker.NewTaskRunnerWithApiClient(apiClient)
)

func main() {

	taskRunner.StartWorker("task_11", task.Task11, 1, time.Millisecond*100)

	logrus.Info("Started Workers")

	taskRunner.WaitWorkers()

}
