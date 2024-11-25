# User Service API Project Overview

## Architecture & Implementation

### Core Components
- Domain: User entity with UUID and Name
- Ports: Repository interface definitions
- Service: Business logic implementation
- Tests: Mock-based testing with testify


### Code Structure
```go
// Domain
type User struct {
    ID   string
    Name string
}

// Ports
type UserRepository interface {
    Save(user User) error
    Get(id string) (User, error)
}

// Service
type UserService struct {
    repo UserRepository
}

func (s *UserService) CreateUser(name string) (User, error)
func (s *UserService) GetUser(id string) (User, error)

// Mock Repository
type MockUserRepository struct {
    mock.Mock
}

// Test Cases
func TestUserService_CreateUser(t *testing.T)
func TestUserService_GetUser(t *testing.T)

