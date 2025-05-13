package models

import "time"

// Структура наших задач
type ToDo struct {
	ID         int
	CreatedAt  time.Time
	CompleteAt time.Time
	Deadline   time.Time
	Priority   Priority
	TaskName   string
	Text       string
}

func NewToDo(id int, taskName, text string) *ToDo {
	return &ToDo {
		ID: id, // Тут можно пошуримурить с ID
		CreatedAt: time.Now(),
		// CompleteAt: time.Time{},
		// Deadline: time.Time{},
		// Priority: Priority{},
		TaskName: taskName,
		Text: text,
	}
}