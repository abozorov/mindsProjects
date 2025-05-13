package repository

import (
	"MindsWallet/internal/errs"
	"MindsWallet/internal/models"
)

func GetAllAccounts() ([]models.Account, error) {
	if accountsList == nil {
		return nil, errs.ErrSomethingWentWrong
	}

	return accountsList, nil
}

func GetAccountByID(accountID int) (models.Account, error) {
	for _, account := range accountsList {
		if account.ID == accountID && account.DeletedAt != nil {
			return account, nil
		}
	}

	return models.Account{}, errs.ErrAccountNotFound
}

func TopUpAccount(accountID int, amount float64) error {
	for i, account := range accountsList {
		if account.ID == accountID {
			accountsList[i].Balance += amount
			return nil
		}
	}

	return errs.ErrAccountNotFound
}

func WithdrawFromAccount(accountID int, amount float64) error {
	for i, account := range accountsList {
		if account.ID == accountID {
			accountsList[i].Balance -= amount
			return nil
		}
	}

	return errs.ErrAccountNotFound
}
