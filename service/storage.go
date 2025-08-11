package service

import (
	"encoding/json"
	"fmt"
	"os"
	"todo/data"
)

const Filename = "store/tasks.json"

func SaveTasks(tasks []data.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal tasks: %w", err)
	}

	err = os.WriteFile(Filename, data, 0644)
	if err != nil {
		return fmt.Errorf("could not write file: %w", err)
	}

	return nil
}

func LoadTasks() (data.TodoList, error) {
	var tasks data.TodoList

	data, err := os.ReadFile(Filename)
	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil // no file yet, return empty list
		}
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	return tasks, nil
}
