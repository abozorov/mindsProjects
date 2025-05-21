package controller

import (
	"ConsoleAppToDo/internal/models"
	"ConsoleAppToDo/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func LoadToDo(w http.ResponseWriter, r *http.Request) {
	// достать из БД и отфильтровать если надо
	var (
		toDoList   []models.ToDo
		err        error
		allToDo    string = r.URL.Query().Get("all")
		isComplete string = r.URL.Query().Get("iscomplete")
	)

	if allToDo == "true" && len(r.URL.Query()) == 1 {
		toDoList, err = service.LoadAllToDo()

	} else if allToDo == "false" && isComplete == "true" {
		toDoList, err = service.Filter(true)

	} else if allToDo == "false" && isComplete == "false" {
		toDoList, err = service.Filter(false)

	} else {
		http.Error(w, "Invalid queries", http.StatusBadRequest)
		return
	}

	if err != nil {
		fmt.Fprintln(w, "Произошла ошибка по причине: ", err)
		return
	}

	// отправить все туду
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toDoList)
}

func CreateToDo(w http.ResponseWriter, r *http.Request) {
	// получить название задачи
	var todo models.ToDo

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}
	// (допы) получить приоритет задачи

	// (допы) срок в днях

	// добавить в бд
	service.CreateToDo(todo.TaskName, todo.Text)
	fmt.Fprintln(w, "Задача успешно создана")
}

func DoneToDo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// отправить запросс для завершения
	if id <= 0 {
		http.Error(w, "D не может быть отрицательным или равным 0", http.StatusBadRequest)
		return
	}
	err = service.DoneToDo(id)

	// ответ клиенту
	if err != nil {
		fmt.Fprintln(w, "Произошла ошибка по причине: ", err)
		return
	}
	fmt.Fprintf(w, "Задача %d успешно выполненв!\n", id)
}

func DeleteToDo(w http.ResponseWriter, r *http.Request) {
	// получить id для удаления
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// отправить запросс для удаления
	if id <= 0 {
		http.Error(w, "D не может быть отрицательным или равным 0", http.StatusBadRequest)
		return
	}
	err = service.DeleteToDo(id)

	// ответ клиенту
	if err != nil {
		fmt.Fprintln(w, "Произошла ошибка по причине: ", err)
		return
	}
	fmt.Fprintf(w, "Задача %d успешно удалено!\n", id)
}
