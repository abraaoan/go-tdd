package repository

import "github.com/abraaoan/todo-list/internal/domain/entity"

type UserRepository interface {
	CreateUser(email, password string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}
