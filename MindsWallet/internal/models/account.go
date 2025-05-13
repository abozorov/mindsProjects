package models

import "time"

type Account struct {
	ID        int
	UserID    int
	Balance   float64
	CreatedAt time.Time
	DeletedAt *time.Time
}
