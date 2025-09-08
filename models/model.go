package models

type Tasks struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Status        string `json:"status"`
	Priority      string `json:"priority"`
	DueDate       string `json:"due_date"`
	CompletedDate string `json:"completed_date"`
}

// type TaskStatus string

// const (
// 	StatusPending    TaskStatus = "pending"
// 	StatusInProgress TaskStatus = "in_progress"
// 	StatusCompleted  TaskStatus = "completed"
// )
