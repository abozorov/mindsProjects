package repository

import (
	"ConsoleAppToDo/internal/models"
	"time"
)

// var data map[int]*models.ToDo = make(map[int]*models.ToDo)
var data map[int]*models.ToDo = map[int]*models.ToDo {
	1: {
		ID: 1,
		CreatedAt: time.Date(2025, time.May, 12, 18, 13, 0, 0, time.Local),
		TaskName: "Выбросить мусор",
		Text: "Попросили выбросить мусор",
	}, 
	2: {
		ID:         2,
		CreatedAt:  time.Now().AddDate(0, 0, -9),
		// CompleteAt: time.Now().AddDate(0, 0, -7),
		TaskName:   "Clean apartment",
		Text:       "Vacuum, mop floors, clean windows.",
	},
	3: {
		ID:         3,
		CreatedAt:  time.Now().AddDate(0, 0, -8),
		CompleteAt: time.Now().AddDate(0, 0, -6),
		TaskName:   "Pay bills",
		Text:       "Electricity, Internet, and phone bills.",
	},
	4: {
		ID:         4,
		CreatedAt:  time.Now().AddDate(0, 0, -7),
		CompleteAt: time.Now().AddDate(0, 0, -5),
		TaskName:   "Call plumber",
		Text:       "Fix the leaking kitchen tap.",
	},
	5: {
		ID:         5,
		CreatedAt:  time.Now().AddDate(0, 0, -6),
		// CompleteAt: time.Now().AddDate(0, 0, -4),
		TaskName:   "Write blog post",
		Text:       "Finish article about Go routines.",
	},
	6: {
		ID:         6,
		CreatedAt:  time.Now().AddDate(0, 0, -5),
		CompleteAt: time.Now().AddDate(0, 0, -3),
		TaskName:   "Team meeting",
		Text:       "Weekly sync with the dev team.",
	},
	7: {
		ID:         7,
		CreatedAt:  time.Now().AddDate(0, 0, -4),
		CompleteAt: time.Now().AddDate(0, 0, -2),
		TaskName:   "Doctor appointment",
		Text:       "Routine check-up at 10 AM.",
	},
	8: {
		ID:         8,
		CreatedAt:  time.Now().AddDate(0, 0, -3),
		CompleteAt: time.Now().AddDate(0, 0, -1),
		TaskName:   "Car service",
		Text:       "Oil change and tire rotation.",
	},
	9: {
		ID:         9,
		CreatedAt:  time.Now().AddDate(0, 0, -2),
		CompleteAt: time.Now().AddDate(0, 0, 0),
		TaskName:   "Book tickets",
		Text:       "Reserve seats for summer vacation.",
	},
	10: {
		ID:         10,
		CreatedAt:  time.Now().AddDate(0, 0, -1),
		CompleteAt: time.Now().AddDate(0, 0, 1),
		TaskName:   "Grocery restock",
		Text:       "Buy fruits, vegetables, and snacks.",
	},
	11: {
		ID:         11,
		CreatedAt:  time.Now().AddDate(0, 0, -10),
		// CompleteAt: time.Now().AddDate(0, 0, -8),
		TaskName:   "Buy groceries",
		Text:       "Milk, eggs, bread, and coffee from the supermarket.",
	},
}
