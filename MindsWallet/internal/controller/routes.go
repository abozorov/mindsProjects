package controller

import "fmt"

var Commands = `
Выберите нужную команду:
0. Выход
1. Получить список всех кошельков
2. Получить кошелек по ID
3. Пополнить кошелек по ID
4. Снять деньги с кошелька по ID`

func RunCommands() {
	for {
		fmt.Println(Commands)
		var command string
		fmt.Scanln(&command)
		switch command {
		case "0":
			fmt.Println("До скорого:)")
			return
		case "1":
			GetAllAccounts()
		case "2":
			GetAccountByID()
		case "3":
			TopUpAccount()
		case "4":
			WithdrawFromAccount()
		default:
			fmt.Println("Вы ввели несуществующую команду!")
		}
	}
}
