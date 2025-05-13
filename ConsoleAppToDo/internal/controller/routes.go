package controller

import "fmt"

func Commands() {
	commands := `
Выерите команду:
	0. Выход
	1. Добавить задачу
	2. Вывести все задачи
	3. Вывести выполненные задачи
	4. Вывести невыполненные задачи
	5. Выполнить задачу из списка невыполненных c ID
	6. Удалить задачу с ID`

	for {
		fmt.Println(commands)
		var c string
		fmt.Scanln(&c)
		switch c {
		case "0":
			fmt.Println("До скорых встреч :D")
			return
		case "1":
			CreateToDo()
		case "2":
			LoadAllToDo()
		case "3":
			Filter(true)
		case "4":
			Filter(false)
		case "5":
			DoneToDo()
		case "6":
			DeleteToDo()
		default:
			fmt.Println("Такой команды нет")
		}
	}
}