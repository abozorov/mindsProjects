package errs

import (
	"errors"
	"net/http"
)

var (
	ErrSomethingWentWrong         = errors.New("something went wrong")
	ErrToDoNotFound               = errors.New("to Do not found")
	ErrToDoWasPreviouslyCompleted = errors.New("the task was previously completed")
	ErrEmptyToDoList              = errors.New("empty To-Do List")

	HTTPErr = map[error]int{
		ErrSomethingWentWrong:         http.StatusBadGateway,
		ErrToDoNotFound:               http.StatusNotFound,
		ErrToDoWasPreviouslyCompleted: http.StatusOK,
		ErrEmptyToDoList:              http.StatusNoContent,
	}
)
