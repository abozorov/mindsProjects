package main

import (
	"MindsWallet/internal/controller"
)

var nums []int

func main() {
	controller.RunCommands()

	///* 1. Получить от пользователя число,
	//2. проверить если число положительное
	//3. Сохранить число в базе данных
	//*/
	//
	//// 1 - транспортный слой (controllers, delivery, handlers)
	//fmt.Println("Enter a number: ")
	//var n int
	//fmt.Scan(&n)
	//
	//// 2 - бизнес логика (service, usecase)
	//if n < 0 {
	//	fmt.Println("Invalid number")
	//	return
	//}
	//
	//// 3 - работу с хранилищами / базами данных (repository)
	//nums = append(nums, n)

}

/*
-> controllers -> service -> repository -> DB
*/
