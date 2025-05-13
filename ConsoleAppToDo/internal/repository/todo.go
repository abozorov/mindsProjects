package repository

import (
	"ConsoleAppToDo/internal/errs"
	"ConsoleAppToDo/internal/models"
	"time"
)

func CreateToDo(taskName, text string) error {
	// добавить в бд
	id := len(data)+1
	data[id] = models.NewToDo(id, taskName, text)

	// ответить
	return nil
}

func LoadAllToDo() ([]models.ToDo, error) {
	// достать из БД и загрузить в массив

	if len(data) == 0 {
		return []models.ToDo{}, errs.ErrEmptyToDoList
	}

	toDo := make([]models.ToDo, 0, len(data))
	for _, v := range data {
		toDo = append(toDo, *v)
	}
	// отправить все туду
	return toDo, nil
}

func DoneToDo(id int) error {
	// завершение туду
	// ответ клиенту
	_, e := data[id]
	if !e {
		return errs.ErrToDoNotFound
	}

	if !data[id].CompleteAt.IsZero() {
		return errs.ErrToDoWasPreviouslyCompleted
	}

	data[id].CompleteAt = time.Now()
	return nil
}

func DeleteToDo(id int) error {
	// удаление туду
	// ответ клиенту
	_, e := data[id]
	if !e {
		return errs.ErrToDoNotFound
	}
	delete(data, id)
	return nil
}