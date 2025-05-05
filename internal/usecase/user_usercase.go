package usecase

import (
	"strconv"

	"github.com/abraaoan/todo-list/internal/domain/entity"
	"github.com/abraaoan/todo-list/internal/provider"
	"github.com/abraaoan/todo-list/internal/repository"
)

type UserUseCase struct {
	repo          repository.UserRepository
	tokenProvider provider.TokenProvider
}

func NewUserUseCase(repo repository.UserRepository, tokenProvider provider.TokenProvider) *UserUseCase {
	return &UserUseCase{repo: repo, tokenProvider: tokenProvider}
}

func (uc *UserUseCase) CreateUser(email, password string) (string, error) {
	user, err := uc.repo.CreateUser(email, password)

	if err != nil {
		return "", entity.ErrUserNotFound
	}

	token, errToken := uc.tokenProvider.Generate(strconv.Itoa(user.ID))
	if errToken != nil {
		return "", errToken
	}

	return token, nil
}

func (uc *UserUseCase) Login(email, password string) (string, error) {
	user, err := uc.repo.FindByEmail(email)
	if err != nil {
		return "", entity.ErrUserNotFound
	}

	if user.Password != password {
		return "", entity.ErrSomethingWrong
	}

	token, errToken := uc.tokenProvider.Generate(strconv.Itoa(user.ID))
	if errToken != nil {
		return "", errToken
	}

	return token, nil
}
