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

	for {
		fmt.Println("O que você quer fazer?")
		fmt.Println("1. Criar usuário")
		fmt.Println("2. Login")
		fmt.Println("3. Criar tarefa")
		fmt.Println("4. Listar tarefas")
		fmt.Println("5. Completar tarefas")
		fmt.Println("0. Sair")
		fmt.Print("> ")

		opt, _ := reader.ReadString('\n')
		opt = strings.TrimSpace(opt)

		switch opt {
		case "1":
			userHandler.CreateUser()
		case "2":
			userHandler.HandleLogin()
		case "3":
			taskHandler.CreateTask()
		case "4":
			taskHandler.ListTask()
		case "5":
			taskHandler.CompleteTask()
		case "0":
			fmt.Println("Até mais!")
			return
		default:
			fmt.Println("Opção inválida.")
		}
	}
}
