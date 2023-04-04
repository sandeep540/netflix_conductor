package task

import (
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

func Task11(task *model.Task) (interface{}, error) {

	taskResult := model.NewTaskResultFromTask(task)
	taskResult.OutputData = map[string]interface{}{
		"key":  "value",
		"key1": "value1",
		"key2": 42,
	}
	taskResult.Status = model.CompletedTask
	return taskResult, nil
}
