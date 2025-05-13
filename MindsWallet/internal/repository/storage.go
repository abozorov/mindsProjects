package repository

import (
	"MindsWallet/internal/models"
	"time"
)

var accountsList = []models.Account{
	{
		ID:        1,
		UserID:    1,
		Balance:   1000,
		CreatedAt: time.Now(),
		DeletedAt: nil,
	},
	{
		ID:        2,
		UserID:    2,
		Balance:   400,
		CreatedAt: time.Now(),
		DeletedAt: nil,
	},
	{
		ID:        3,
		UserID:    5,
		Balance:   800,
		CreatedAt: time.Now(),
		DeletedAt: nil,
	},
	{
		ID:        4,
		UserID:    8,
		Balance:   800,
		CreatedAt: time.Now(),
		DeletedAt: nil,
	},
	{
		ID:        5,
		UserID:    9,
		Balance:   100,
		CreatedAt: time.Now(),
		DeletedAt: &time.Time{},
	},
}
