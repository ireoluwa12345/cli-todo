package cmd

import (
	"cli/todo/models"
	_ "embed"
	"encoding/json"
	"errors"
	"strconv"
)

//go:embed tasks.json
var tasksJson []byte

func getTasks() []models.Tasks {
	var tasks []models.Tasks
	err := json.Unmarshal(tasksJson, &tasks)
	if err != nil {
		panic(err)
	}

	return tasks
}

func getTask(arg any) (models.Tasks, error) {
	var tasks []models.Tasks
	err := json.Unmarshal(tasksJson, &tasks)

	if err != nil {
		panic(err)
	}

	taskMap := make(map[string]models.Tasks, len(tasks)*2)
	for _, task := range tasks {
		taskMap[task.Id] = task
		taskMap[task.Name] = task
	}
	if t, ok := taskMap[arg.(string)]; ok {
		return t, nil
	}

	return models.Tasks{}, errors.New("no task found")
}

func createTask(task models.Tasks) {
	var tasks []models.Tasks
	err := json.Unmarshal(tasksJson, &tasks)

	if err != nil {
		panic(err)
	}

	task_length := len(tasks)

	lastTask := tasks[task_length-1]
	id, _ := strconv.Atoi(lastTask.Id)

	task.Id = strconv.Itoa(id + 1)

	tasks = append(tasks, task)

	updatedTasksJson, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		panic(err)
	}

	tasksJson = updatedTasksJson
}
