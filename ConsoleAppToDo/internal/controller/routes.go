package controller

import (
	"fmt"
	"net/http"
)

func router(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		LoadToDo(w, r)
	case http.MethodPost:
		CreateToDo(w, r)
	case http.MethodPut:
		DoneToDo(w, r)
	case http.MethodDelete:
		DeleteToDo(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

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
	fmt.Println(commands)

	mux := http.NewServeMux()
	mux.HandleFunc("/todo", router)
	mux.HandleFunc("/", router)

	http.ListenAndServe(":8080", mux)

}
