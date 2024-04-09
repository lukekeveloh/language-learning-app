// internal/app/app.go

package app

import (
	"github.com/lukekeveloh/language-learning-app/internal/app/handlers"
	"github.com/lukekeveloh/language-learning-app/internal/app/repository"
	"github.com/lukekeveloh/language-learning-app/internal/app/services"
)

// Application struct holds the dependencies for the application
type Application struct {
	UserService services.UserService
	DeckService services.DeckService
	GameService services.GameService
}

// Initialize initializes the application with its dependencies
func Initialize() *Application {
	// Initialize services
	userService := services.NewUserService(repository.NewUserRepository())
	deckService := services.NewDeckService(repository.NewDeckRepository())
	gameService := services.NewGameService() // Add game repository if needed

	// Initialize handlers with services
	userHandler := handlers.NewUserHandler(userService)
	deckHandler := handlers.NewDeckHandler(deckService)
	// Initialize game handler if needed

	return &Application{
		UserService: userService,
		DeckService: deckService,
		GameService: gameService,
	}
}
