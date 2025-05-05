package repository

import "github.com/abraaoan/todo-list/internal/domain/entity"

type TaskRepository interface {
	Save(title string, userID int) (*entity.Task, error)
	List(userID int) ([]entity.Task, error)
	FindById(id int) (*entity.Task, error)
	Delete(id int, userID int) error
	Update(*entity.Task) error
}
