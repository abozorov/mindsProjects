package controller

import (
	"ConsoleAppToDo/internal/models"
	"ConsoleAppToDo/internal/service"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printToDo(t models.ToDo) {
	fmt.Printf("Task#%d - %s\nText: %s\nCreated: %s\n",
		t.ID,
		t.TaskName,
		t.Text,
		t.CreatedAt.Format("02-Jan-2006, 15:04"),
	)
	if t.CompleteAt.IsZero() {
		fmt.Println("Задача пока не выполнена!\n____________")
	} else {
		fmt.Printf("Complete time: %s\n____________\n",
			t.CompleteAt.Format("02-Jan-2006, 15:04"),
		)
	}
}

func CreateToDo() {
	// получить название задачи
	fmt.Println("Введите название задачи!")
	reader := bufio.NewReader(os.Stdin)
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSpace(taskName)

	// получитьт текст задачи
	fmt.Println("Введите описание к задаче!")
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)

	// (допы) получить приоритет задачи

	// (допы) срок в днях

	// добавить в бд
	service.CreateToDo(taskName, task)
	fmt.Println("Задача успешно создана")
}

func LoadAllToDo() {
	// достать из БД
	toDoList, err := service.LoadAllToDo()
	if err != nil {
		fmt.Println("Произошла ошибка по причине: ", err)
		return
	}

	// вывести все туду
	for _, v := range toDoList {
		printToDo(v)
	}
}

func Filter(isComplete bool) {
	// достать отсортированное
	filter, err := service.Filter(isComplete)
	if err != nil {
		fmt.Println("Произошла ошибка по причине: ", err)
		return
	}

	// вывести пользователю
	if isComplete {
		fmt.Println("Выполненные задачи!")
	} else {
		fmt.Println("Невыполненные задачи!")
	}
	for _, v := range filter {
		printToDo(v)
	}
}

func DoneToDo() {
	// достать и показать незвершенные
	filter, err := service.Filter(false)
	if err != nil {
		fmt.Println("Произошла ошибка по причине: ", err)
		return
	}
	fmt.Println("Список невыполненных задач!")
	for _, v := range filter {
		printToDo(v)
	}

	// получить id которую надо завершить
	fmt.Println("Введите ID задачи")
	var id int
	fmt.Scan(&id)

	// отправить запросс для завершения
	if id <= 0 {
		fmt.Println("ID не может быть отрицательным или равным 0.\nПосмотрите список повнимательнее")
		return
	}
	err = service.DoneToDo(id)

	// ответ клиенту
	if err != nil {
		fmt.Println("Произошла ошибка по причине: ", err)
		return
	}
	fmt.Printf("Задача %d успешно выполненв!\n", id)
}

func DeleteToDo() {
	// вывести список c ID клиенту
	toDoList, err := service.LoadAllToDo()
	if err != nil {
		fmt.Println("Произошла ошибка по причине: ", err)
		return
	}
	fmt.Println("Выберите ID для удаления из списка!")
	for _, v := range toDoList {
		printToDo(v)
	}

	// получить id для удаления
	var id int
	fmt.Scan(&id)

	// запросс для удаления
	if id <= 0 {
		fmt.Println("ID не может быть отрицательным или равным 0.\nПосмотрите список повнимательнее")
		return
	}
	err = service.DeleteToDo(id)

	// ответ клиенту
	if err != nil {
		fmt.Println("Произошла ошибка по причине: ", err)
		return
	}
	fmt.Printf("Задача %d успешно удалено!\n", id)
}
