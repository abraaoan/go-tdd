package usecase_test

import (
	"testing"

	"github.com/abraaoan/todo-list/internal/domain/entity"
	"github.com/abraaoan/todo-list/internal/usecase"
	"github.com/abraaoan/todo-list/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockTaskRepository(ctrl)
	mockRepo.EXPECT().
		Save("fake title", 1).
		Return(&entity.Task{
			ID:     1,
			Title:  "fake title",
			Status: false,
			UserID: 1,
		}, nil)

	uc := usecase.NewTaskUseCase(mockRepo)

	task, err := uc.CreateTask(1, "fake title")

	assert.NoError(t, err)
	assert.Equal(t, "fake title", task.Title)
	assert.Equal(t, false, task.Status)
	assert.Greater(t, task.ID, 0)
}

func TestListTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTaskRepository(ctrl)
	mockRepo.EXPECT().
		List(1).
		Return([]entity.Task{
			{
				ID:     1,
				Title:  "Title #1",
				Status: false,
				UserID: 1,
			},
			{
				ID:     2,
				Title:  "Title #2",
				Status: false,
				UserID: 1,
			},
			{
				ID:     3,
				Title:  "Title #3",
				Status: false,
				UserID: 1,
			},
		}, nil)

	uc := usecase.NewTaskUseCase(mockRepo)

	tasks, err := uc.ListTasks(1)

	assert.NoError(t, err)
	assert.Greater(t, len(tasks), 0)
	assert.Equal(t, "Title #1", tasks[0].Title)
}

func TestCompleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTaskRepository(ctrl)
	mockRepo.EXPECT().
		FindById(1).
		Return(&entity.Task{
			ID:     1,
			Title:  "Title #1",
			Status: false,
			UserID: 1,
		}, nil)

	uc := usecase.NewTaskUseCase(mockRepo)
	task, err := uc.CompleteTask(1, 1)

	assert.NoError(t, err)
	assert.Equal(t, true, task.Status)
}

func TestCompleteTaskFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTaskRepository(ctrl)
	mockRepo.EXPECT().
		FindById(1).
		Return(nil, entity.ErrTaskNotFound)

	uc := usecase.NewTaskUseCase(mockRepo)
	task, err := uc.CompleteTask(1, 1)

	assert.Error(t, err)
	assert.Nil(t, task)
	assert.ErrorIs(t, err, entity.ErrTaskNotFound)
}

func TestDeleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTaskRepository(ctrl)
	mockRepo.EXPECT().
		Delete(1, 1).
		Return(nil)

	uc := usecase.NewTaskUseCase(mockRepo)
	err := uc.DeleteTask(1, 1)

	assert.NoError(t, err)
}

func TestDeleteTaskFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTaskRepository(ctrl)
	mockRepo.EXPECT().Delete(1, 1).Return(entity.ErrTaskNotFound)

	uc := usecase.NewTaskUseCase(mockRepo)
	err := uc.DeleteTask(1, 1)

	assert.Error(t, err)
	assert.ErrorIs(t, err, entity.ErrTaskNotFound)
}
