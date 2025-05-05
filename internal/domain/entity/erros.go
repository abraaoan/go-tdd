package entity

import "errors"

var (
	ErrTaskNotFound         = errors.New("task not found")
	ErrTaskAlreadyCompleted = errors.New("task already completed")
	ErrUserNotFound         = errors.New("user not found")
	ErrSomethingWrong       = errors.New("somthing went wrong")
	ErrTitleRequired        = errors.New("title is required")
)
