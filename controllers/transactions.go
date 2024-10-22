package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luisaugustomelo/pismo-challenge/db"
	"github.com/luisaugustomelo/pismo-challenge/handlers"
	"github.com/luisaugustomelo/pismo-challenge/services"
)

func SetupTransactionRoutes(app *fiber.App) {
	api := app.Group("/transactions")

	// Dependency Injection
	transactionService := services.NewTransactionService(db.DB)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	registerCreateTransactionHandler(api, transactionHandler)
}

func registerCreateTransactionHandler(api fiber.Router, transactionHandler *handlers.TransactionHandler) {
	api.Post("/", transactionHandler.CreateTransaction)
}
