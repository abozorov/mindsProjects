package errs

import "errors"

var (
	ErrSomethingWentWrong         = errors.New("something went wrong")
	ErrToDoNotFound               = errors.New("To Do not found")
	ErrToDoWasPreviouslyCompleted = errors.New("the task was previously completed")
	ErrEmptyToDoList              = errors.New("Empty To-Do List")
)
