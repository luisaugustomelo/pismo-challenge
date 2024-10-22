package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luisaugustomelo/pismo-challenge/db"
	"github.com/luisaugustomelo/pismo-challenge/handlers"
	"github.com/luisaugustomelo/pismo-challenge/services"
)

func SetupAccountRoutes(app *fiber.App) {
	api := app.Group("/accounts")

	// Dependency Injection
	accountService := services.NewAccountService(db.DB)
	accountHandler := handlers.NewAccountHandler(accountService)

	registerCreateAccountHandler(api, accountHandler)
}

func registerCreateAccountHandler(api fiber.Router, accountHandler *handlers.AccountHandler) {
	api.Post("/", accountHandler.CreateAccount)
	api.Get("/:id", accountHandler.GetAccount)
}
