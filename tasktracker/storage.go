package tasktracker

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const tasksFile = "tasks.json"

func SaveTasks() error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(tasksFile, data, 0644)
}

func LoadTasks() error {
	if _, err := os.Stat(tasksFile); os.IsNotExist(err) {
		tasks = []Task{}
		return nil // Arquivo não existe, lista vazia
	}
	data, err := ioutil.ReadFile(tasksFile)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &tasks)
}
