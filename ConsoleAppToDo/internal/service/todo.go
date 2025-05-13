package service

import (
	"ConsoleAppToDo/internal/errs"
	"ConsoleAppToDo/internal/models"
	"ConsoleAppToDo/internal/repository"
	"sort"
)

func CreateToDo(taskName, text string) error {
	// добавить в бд
	err := repository.CreateToDo(taskName, text)

	// ответить
	if err != nil {
		return err
	}
	return nil
}

func LoadAllToDo() ([]models.ToDo, error) {
	// достать из БД
	toDos, err := repository.LoadAllToDo()
	if err != nil {
		return []models.ToDo{}, err
	}
	// Сортировка
	sort.Slice(toDos, func(i, j int) bool {
		return toDos[i].CreatedAt.Compare(toDos[j].CreatedAt) < 0
	})
	// отправить все туду
	return toDos, nil
}

func Filter(isComplete bool) ([]models.ToDo, error) {
	// достать все туду
	toDoList, err := repository.LoadAllToDo()
	if err != nil {
		return []models.ToDo{}, err
	}

	// Фильтрация и сортировка
	filter := make([]models.ToDo, 0, len(toDoList))
	for _, v := range toDoList {
		if v.CompleteAt.IsZero() != isComplete {
			filter = append(filter, v)
		}
	}

	if len(filter) == 0 {
		return []models.ToDo{}, errs.ErrEmptyToDoList
	}

	sort.Slice(filter, func(i, j int) bool {
		return (filter[i].CompleteAt.Compare(filter[j].CompleteAt) > 0 && isComplete) ||
			(filter[i].CreatedAt.Compare(filter[j].CreatedAt) < 0 && !isComplete)
	})

	// отправить пользователю
	return filter, nil
}

func DoneToDo(id int) error {
	// запросс для завершения туду
	err := repository.DoneToDo(id)
	if err != nil {
		return err
	}

	// ответ клиенту
	return nil
}

func DeleteToDo(id int) error {
	// запросс для удаления
	err := repository.DeleteToDo(id)
	if err != nil {
		return err
	}

	// ответ клиенту
	return nil
}