package main

import (
	"encoding/json"
	"os"
)

const dataFile = "tasks.json"

func Load() (*TaskList, error) {
	data, err := os.ReadFile(dataFile)

	if os.IsNotExist(err) {
		return NewTaskList(), nil
	}

	if err != nil {
		return nil, err
	}

	var tl TaskList
	if err := json.Unmarshal(data, &tl); err != nil {
		return nil, err
	}

	maxId := 0

	for _, task := range tl.Tasks {
		if task.Id > maxId {
			maxId = task.Id
		}
	}

	tl.NextId = maxId + 1

	return &tl, nil
}

func save(taskList *TaskList) error {
	data, err := json.MarshalIndent(taskList, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(dataFile, data, 0644)
}
