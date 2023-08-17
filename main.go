package main

import (
	"alghanim/mainAPI/api/router"
	"alghanim/mainAPI/database"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Connect to database
	db := database.Connect()
	defer db.Close()

	// Initialize routes
	router.Init(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
