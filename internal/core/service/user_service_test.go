package service

import (
	"awesomeProject/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// MockUserRepository is a mock implementation of ports.UserRepository
type MockUserRepository struct {
	mock.Mock
}

// Save saves the given user to the mock repository. It returns an error if any
// occurred during the save operation.
func (m *MockUserRepository) Save(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// Get retrieves a user from the mock repository by the given ID.
// It returns the user and an error if any occurred during the retrieval.
func (m *MockUserRepository) Get(id string) (domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(domain.User), args.Error(1)
}

// mockRepo is a new instance of the MockUserRepository, which is a mock implementation
// of the ports.UserRepository interface. This mock repository is used in tests to
// simulate the behavior of the real repository, allowing for controlled testing of
// the UserService.
func TestUserService_CreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	t.Run("successful user creation", func(t *testing.T) {
		userName := "John Doe"
		mockRepo.On("Save", mock.AnythingOfType("domain.User")).Return(nil)

		user, err := service.CreateUser(userName)

		assert.NoError(t, err)
		assert.NotEmpty(t, user.ID)
		assert.Equal(t, userName, user.Name)
		mockRepo.AssertExpectations(t)
	})
}

// TestUserService_GetUser tests the GetUser method of the UserService.
// It creates a mock UserRepository, sets up an expected user to be returned,
// and verifies that the service correctly retrieves the user.
func TestUserService_GetUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	t.Run("successful user retrieval", func(t *testing.T) {
		expectedUser := domain.User{
			ID:   "123e4567-e89b-12d3-a456-426614174000",
			Name: "John Doe",
		}
		mockRepo.On("Get", expectedUser.ID).Return(expectedUser, nil)

		user, err := service.GetUser(expectedUser.ID)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		mockRepo.AssertExpectations(t)
	})
}
