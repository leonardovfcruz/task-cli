package tasktracker

import (
	"errors"
	"time"
)

// Global variable to store tasks
var tasks []Task

// init loads tasks from storage when the package is initialized
func init() {
	if err := LoadTasks(); err != nil {
		// If there's an error loading tasks, initialize with an empty slice
		tasks = []Task{}
	}
}

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

func AddTask(description string) Task {
	id := getNextID()
	now := time.Now()
	task := Task{
		ID:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	tasks = append(tasks, task)
	// Save tasks to storage
	_ = SaveTasks()
	return task
}

func UpdateTask(id int, newDescription string) error {
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Description = newDescription
			tasks[i].UpdatedAt = time.Now()
			// Save tasks to storage
			_ = SaveTasks()
			return nil
		}
	}
	return errors.New("tarefa não encontrada")
}

func DeleteTask(id int) error {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("tarefa não encontrada")
}

func ListTasks() []Task {
	return tasks
}

func getNextID() int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func MarkInProgress(id int) error {
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = "in-progress"
			tasks[i].UpdatedAt = time.Now()
			return nil
		}
	}
	return errors.New("tarefa não encontrada")
}

func MarkDone(id int) error {
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = "done"
			tasks[i].UpdatedAt = time.Now()
			return nil
		}
	}
	return errors.New("tarefa não encontrada")
}

func ListTasksByStatus(status string) []Task {
	var filtered []Task
	for _, t := range tasks {
		if t.Status == status {
			filtered = append(filtered, t)
		}
	}
	return filtered
}
