package main

import (
	"awesomeProject/internal/adapters/secondary"
	"awesomeProject/internal/core/service"
	"fmt"
	"log"
)

// main is the entry point of the application. It initializes the necessary dependencies,
// creates a new user, and then retrieves the created user.
func main() {
	// Initialize dependencies
	repo := secondary.NewInMemoryUserRepository()
	userService := service.NewUserService(repo)

	// Create a new user
	user, err := userService.CreateUser("John Doe")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created user: %+v\n", user)

	foundUser, err := userService.GetUser(user.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found user: %+v\n", foundUser)
} // Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
