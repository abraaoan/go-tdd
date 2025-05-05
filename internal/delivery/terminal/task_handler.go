package terminal

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/abraaoan/todo-list/internal/domain/entity"
	"github.com/abraaoan/todo-list/internal/usecase"
)

type TaskHandler struct {
	uc     *usecase.TaskUseCase
	reader *bufio.Reader
}

func NewTaskHandler(uc *usecase.TaskUseCase, reader *bufio.Reader) *TaskHandler {
	return &TaskHandler{uc: uc, reader: reader}
}

func (h *TaskHandler) CreateTask() {
	fmt.Print("Qual o user id: ")
	userID, _ := h.reader.ReadString('\n')
	fmt.Print("Qual o nome da task: ")
	task, _ := h.reader.ReadString('\n')

	userID = strings.TrimSpace(userID)
	task = strings.TrimSpace(task)
	userIDInt, _ := strconv.Atoi(userID)

	taskCreated, err := h.uc.CreateTask(userIDInt, task)
	if err != nil {
		fmt.Println("Erro ao criar usuário:", err)
		return
	}

	fmt.Println("✅ 4Tarefa criada com sucesso", taskCreated.Title)
	fmt.Print("\n")
}

func (h *TaskHandler) ListTask() {
	fmt.Print("Qual o user id: ")
	userID, _ := h.reader.ReadString('\n')

	userID = strings.TrimSpace(userID)
	userIDInt, _ := strconv.Atoi(userID)

	tasks, err := h.uc.ListTasks(userIDInt)
	if err != nil {
		fmt.Println("❌ Erro ao listar tarefas:", err)
		return
	}

	printTasks(tasks)
}

func (h *TaskHandler) CompleteTask() {
	fmt.Print("Qual o user id: ")
	userID, _ := h.reader.ReadString('\n')

	fmt.Print("Qual o task id: ")
	taskID, _ := h.reader.ReadString('\n')

	userID = strings.TrimSpace(userID)
	userIDInt, _ := strconv.Atoi(userID)
	taskID = strings.TrimSpace(taskID)
	taskIDInt, _ := strconv.Atoi(taskID)

	task, err := h.uc.CompleteTask(taskIDInt, userIDInt)
	if err != nil {
		fmt.Println("❌ Erro ao completar tarefas:", err)
		return
	}

	printTask(task)
}

func printTasks(tasks []entity.Task) {
	for _, t := range tasks {
		printTask(&t)
	}

	fmt.Print("\n")
}

func printTask(t *entity.Task) {
	status := "❌ Pendente"
	if t.Status {
		status = "✅ Concluída"
	}

	fmt.Printf("ID:     %d\n", t.ID)
	fmt.Printf("Título: %s\n", t.Title)
	fmt.Printf("Status: %s\n", status)
	fmt.Printf("Dono:   %d\n", t.UserID)
	fmt.Println(strings.Repeat("-", 30))
}
