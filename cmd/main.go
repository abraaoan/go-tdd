package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"

	"github.com/abraaoan/todo-list/internal/adapter/postgres"
	"github.com/abraaoan/todo-list/internal/delivery/terminal"
	"github.com/abraaoan/todo-list/internal/provider"
	"github.com/abraaoan/todo-list/internal/usecase"
)

func main() {
	// Database
	connStr := "postgresql://bob:sponge@localhost:5432/todo_list?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Connect database error %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Erro ao pingar banco: %v", err)
	}

	fmt.Println("✅ Banco conectado!")

	// Postgres repo
	taskRepo := postgres.NewPostgresTaskRepository(db)
	userRepo := postgres.NewPostgresUserRepository(db)

	// Token provider
	tokenProvider := provider.NewSimpleTokenProvider()

	// Use case
	taskUseCase := usecase.NewTaskUseCase(taskRepo)
	userUseCase := usecase.NewUserUseCase(userRepo, tokenProvider)

	reader := bufio.NewReader(os.Stdin)

	// Handler
	userHandler := terminal.NewUserHandler(userUseCase, reader)
	taskHandler := terminal.NewTaskHandler(taskUseCase, reader)

	fmt.Print("\n")
	for {
		fmt.Println("O que você quer fazer?")
		fmt.Println("1. Criar usuário")
		fmt.Println("2. Listar usuários")
		fmt.Println("3. Login")
		fmt.Println("4. Criar tarefa")
		fmt.Println("5. Listar tarefas")
		fmt.Println("6. Completar tarefas")
		fmt.Println("0. Sair")
		fmt.Print("> ")

		opt, _ := reader.ReadString('\n')
		opt = strings.TrimSpace(opt)

		switch opt {
		case "1":
			userHandler.CreateUser()
		case "2":
			userHandler.ListUsers()
		case "3":
			userHandler.HandleLogin()
		case "4":
			taskHandler.CreateTask()
		case "5":
			taskHandler.ListTask()
		case "6":
			taskHandler.CompleteTask()
		case "0":
			fmt.Print("\n 👋 Até mais! \n\n")
			return
		default:
			fmt.Println("Opção inválida.")
		}
	}
}
