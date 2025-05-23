package usecase

import (
	"log"
	"strconv"

	"github.com/abraaoan/todo-list/internal/domain/entity"
	"github.com/abraaoan/todo-list/internal/provider"
	"github.com/abraaoan/todo-list/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repo          repository.UserRepository
	tokenProvider provider.TokenProvider
}

func NewUserUseCase(repo repository.UserRepository, tokenProvider provider.TokenProvider) *UserUseCase {
	return &UserUseCase{repo: repo, tokenProvider: tokenProvider}
}

func (uc *UserUseCase) CreateUser(email, name, password, role string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	user, err := uc.repo.CreateUser(email, name, string(hash), role)

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

	if !isValidUser(password, user) {
		print("-----> AQUI")
		return "", entity.ErrSomethingWrong
	}

	token, errToken := uc.tokenProvider.Generate(strconv.Itoa(user.ID))
	if errToken != nil {
		return "", errToken
	}

	return token, nil
}

func isValidUser(password string, user *entity.User) bool {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return false
	}

	err = bcrypt.CompareHashAndPassword(hash, []byte(user.Password))
	if err != nil {
		return false
	} else {
		return true
	}
}
