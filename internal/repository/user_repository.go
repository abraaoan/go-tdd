package repository

import "github.com/abraaoan/todo-list/internal/domain/entity"

type UserRepository interface {
	CreateUser(email, name, password, role string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Find(userId int) (*entity.User, error)
	DeleteUser(userId int) error
	ListUsers() ([]entity.User, error)
	UpdateUser(*entity.User) (*entity.User, error)
}
