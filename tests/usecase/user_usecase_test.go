package usecase_test

import (
	"testing"

	"github.com/abraaoan/todo-list/internal/domain/entity"
	"github.com/abraaoan/todo-list/internal/usecase"
	"github.com/abraaoan/todo-list/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestLoginSuccessfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().FindByEmail("appleSeed@apple.com").Return(&entity.User{
		ID:       1,
		Email:    "appleSeed@apple.com",
		Password: "123456",
	}, nil)
	mockTokenProvider := mocks.NewMockTokenProvider(ctrl)
	mockTokenProvider.EXPECT().Generate("1").Return("fake-token-generate", nil)

	uc := usecase.NewUserUseCase(mockRepo, mockTokenProvider)
	token, err := uc.Login("appleSeed@apple.com", "123456")

	assert.NoError(t, err)
	assert.Equal(t, "fake-token-generate", token)
}

func TestLoginFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().FindByEmail("appleSeed@apple.com").Return(nil, entity.ErrUserNotFound)

	mockTokenProvider := mocks.NewMockTokenProvider(ctrl)

	uc := usecase.NewUserUseCase(mockRepo, mockTokenProvider)
	token, err := uc.Login("appleSeed@apple.com", "123456")

	assert.Error(t, err)
	assert.ErrorIs(t, entity.ErrUserNotFound, err)
	assert.Empty(t, token)
}

func TestLoginWrongPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().FindByEmail("appleSeed@apple.com").Return(&entity.User{
		ID:       1,
		Email:    "appleSeed@apple.com",
		Password: "Wrong pass",
	}, nil)
	mockTokenProvider := mocks.NewMockTokenProvider(ctrl)

	uc := usecase.NewUserUseCase(mockRepo, mockTokenProvider)
	token, err := uc.Login("appleSeed@apple.com", "123456")

	assert.Error(t, err)
	assert.ErrorIs(t, err, entity.ErrSomethingWrong)
	assert.Empty(t, token)
}

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().
		CreateUser("apple@seed.com", "password").
		Return(&entity.User{ID: 1, Email: "apple@seed.com", Password: "password"}, nil)

	mockTokenProvider := mocks.NewMockTokenProvider(ctrl)
	mockTokenProvider.EXPECT().Generate("1").Return("fake-token-generate", nil)

	uc := usecase.NewUserUseCase(mockRepo, mockTokenProvider)
	token, err := uc.CreateUser("apple@seed.com", "password")

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
