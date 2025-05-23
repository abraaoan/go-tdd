package terminal

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/abraaoan/todo-list/internal/usecase"
)

type UserHandler struct {
	uc     *usecase.UserUseCase
	reader *bufio.Reader
}

func NewUserHandler(uc *usecase.UserUseCase, reader *bufio.Reader) *UserHandler {
	return &UserHandler{uc: uc, reader: reader}
}

func (h *UserHandler) CreateUser() {
	fmt.Print("Email: ")
	email, _ := h.reader.ReadString('\n')

	fmt.Print("Name: ")
	name, _ := h.reader.ReadString('\n')

	fmt.Print("Password: ")
	password, _ := h.reader.ReadString('\n')

	fmt.Print("Role: ")
	role, _ := h.reader.ReadString('\n')

	email = strings.TrimSpace(email)
	name = strings.TrimSpace(name)
	password = strings.TrimSpace(password)
	role = strings.TrimSpace(role)

	token, err := h.uc.CreateUser(email, name, password, role)
	if err != nil {
		fmt.Println("❌ Error ao criar o usuário. ", err)
		fmt.Print("\n")
		return
	}

	fmt.Println("✅ Usuário criado com sucesso. Token: ", token)
	fmt.Print("\n")
}

func (h *UserHandler) HandleLogin() {
	fmt.Print("Email: ")
	email, _ := h.reader.ReadString('\n')

	fmt.Print("Password: ")
	password, _ := h.reader.ReadString('\n')

	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	token, err := h.uc.Login(email, password)
	if err != nil {
		fmt.Println("❌ Erro ao logar. ", err)
		fmt.Print("\n")
		return
	}

	fmt.Println("✅ Usuário logado com sucesso: ", token)
	fmt.Print("\n")
}
