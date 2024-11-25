package secondary

import (
	"awesomeProject/internal/core/domain"
	"errors"
)

// InMemoryUserRepository is an in-memory implementation of the UserRepository interface.
// It stores user data in a map, keyed by the user's ID.
type InMemoryUserRepository struct {
	users map[string]domain.User
}

// NewInMemoryUserRepository creates a new instance of the InMemoryUserRepository struct,
// which implements the UserRepository interface. It initializes the users map to store
// user data in memory.
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]domain.User),
	}
}

// Save stores the provided user in the in-memory repository, keyed by the user's ID.
// If the user already exists, it will be overwritten.
func (r *InMemoryUserRepository) Save(user domain.User) error {
	r.users[user.ID] = user
	return nil
}

// Get retrieves a user from the in-memory repository by the provided ID. If the user is not found,
// it returns an error.
func (r *InMemoryUserRepository) Get(id string) (domain.User, error) {
	user, exists := r.users[id]
	if !exists {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}
