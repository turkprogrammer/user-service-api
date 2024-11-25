package ports

import "awesomeProject/internal/core/domain"

// UserRepository is the interface for managing user-related data persistence operations.
// It provides methods to save a user to the underlying storage and retrieve an existing user.
type UserRepository interface {
	// Save persists the given user to the underlying storage.
	Save(user domain.User) error
	Get(id string) (domain.User, error)
}

// UserService is the interface for managing user-related operations.
// It provides methods to create a new user and retrieve an existing user.
type UserService interface {
	CreateUser(name string) (domain.User, error)
	GetUser(id string) (domain.User, error)
}
