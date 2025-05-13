package controller

import (
	"MindsWallet/internal/service"
	"fmt"
)

func GetAllAccounts() {
	// получить список кошельков
	accounts, err := service.GetAllAccounts()
	if err != nil {
		fmt.Println("Не удалось получить список кошельков, причина:", err.Error())
		return
	}

	// вывести клиенту
	for _, account := range accounts {
		fmt.Printf("ID: %d, Balance: %.2f\n", account.ID, account.Balance)
	}
}

func GetAccountByID() {
	// получить от клиента ID кошелька
	fmt.Println("Введите ID кошелька: ")
	var accountID int
	fmt.Scan(&accountID)

	if accountID <= 0 {
		fmt.Println("ID кошелька невалидный")
		return
	}

	// достать из бд кошелек
	account, err := service.GetAccountByID(accountID)
	if err != nil {
		fmt.Println("Не удалось получить кошелек, причина:", err.Error())
		return
	}

	// вывести клиенту
	fmt.Printf("ID: %d, Balance: %f\n", account.ID, account.Balance)
	fmt.Printf("CreatedAt: %s, UserID: %d\n", account.CreatedAt, account.UserID)
}

func TopUpAccount() {
	// получить от клиента ID кошелька
	fmt.Println("Введите ID кошелька: ")
	var accountID int
	fmt.Scan(&accountID)

	if accountID <= 0 {
		fmt.Println("ID кошелька невалидный")
		return
	}

	// получить сумму пополнения
	fmt.Println("Введите сумму пополнения: ")
	var amount float64
	fmt.Scan(&amount)
	if amount <= 0 {
		fmt.Println("Сумма должна быть больше 0")
		return
	}

	// отправить запрос на пополнение
	err := service.TopUpAccount(accountID, amount)
	if err != nil {
		fmt.Println("Не удалось пополнить кошелек, причина:", err.Error())
		return
	}

	// ответить клиенту
	fmt.Println("Ваш кошелек успешно пополнен!")
}

func WithdrawFromAccount() {
	// получить от клиента ID кошелька
	fmt.Println("Введите ID кошелька: ")
	var accountID int
	fmt.Scan(&accountID)

	if accountID <= 0 {
		fmt.Println("ID кошелька невалидный")
		return
	}

	// получить сумму снятия
	fmt.Println("Введите сумму снятия: ")
	var amount float64
	fmt.Scan(&amount)
	if amount <= 0 {
		fmt.Println("Сумма должна быть больше 0")
		return
	}

	// отправить запрос на снятие
	err := service.WithdrawFromAccount(accountID, amount)
	if err != nil {
		fmt.Println("Не удалось снять деньги с кошелька, причина:", err.Error())
		return
	}

	// ответить клиенту
	fmt.Println("Успешная транзакция")
}
