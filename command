mkdir -p internal/core/domain internal/core/ports internal/core/service

go mod tidy
go test ./...

userService := service.NewUserService(repository)
user, err := userService.CreateUser("John Doe")
retrievedUser, err := userService.GetUser(user.ID)

