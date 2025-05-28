package tasktracker

import "time"

type Task struct {
	id        int       `json:"id"`
	title     string    `json:"title"`
	status    string    `json:"status"`
	createdAt time.Time `json:"createdAt"`
	updatedAt time.Time `json:"updatedAt"`
}

type Tasks struct {
	tasks []Task `json:"tasks"`
}
