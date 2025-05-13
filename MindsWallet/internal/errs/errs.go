package errs

import "errors"

var (
	ErrSomethingWentWrong = errors.New("something went wrong")
	ErrAccountNotFound    = errors.New("account not found")
	ErrNotEnoughBalance   = errors.New("not enough balance")
)
