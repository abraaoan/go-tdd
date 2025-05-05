package usecase

import (
	"github.com/abraaoan/todo-list/internal/domain/entity"
	"github.com/abraaoan/todo-list/internal/repository"
)

type TaskUseCase struct {
	repo repository.TaskRepository
}

func NewTaskUseCase(repo repository.TaskRepository) *TaskUseCase {
	return &TaskUseCase{repo: repo}
}

func (uc *TaskUseCase) CreateTask(userID int, title string) (*entity.Task, error) {
	return uc.repo.Save(title, userID)
}

func (uc *TaskUseCase) ListTasks(userID int) ([]entity.Task, error) {
	return uc.repo.List(userID)
}

func (uc *TaskUseCase) CompleteTask(id int, userID int) (*entity.Task, error) {

	task, err := uc.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	task.Status = true

	return task, nil
}

func (uc *TaskUseCase) DeleteTask(id int, userID int) error {
	return uc.repo.Delete(id, userID)
}
