package router

import (
	"alghanim/mainAPI/api/handler"
	"alghanim/mainAPI/api/middleware"
	"alghanim/mainAPI/repository"
	"alghanim/mainAPI/service"
	"database/sql"

	"github.com/Nerzal/gocloak/v8"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo, db *sql.DB) {
	// Initialize Keycloak client
	client := gocloak.NewClient("https://your-keycloak-server.com")

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)

	// Setup routes with Keycloak middleware
	e.GET("/users/:id", userHandler.Get, middleware.KeycloakMiddleware(client, "your-realm"))

}
