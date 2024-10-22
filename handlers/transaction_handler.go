package handlers

import (
	"github.com/luisaugustomelo/pismo-challenge/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	transactionService services.TransactionService
}

// TransactionRequest is the request struct for transaction creation
type TransactionRequest struct {
	AccountID       uint    `json:"account_id" example:"1"`
	OperationTypeID uint    `json:"operation_type_id" example:"4"`
	Amount          float64 `json:"amount" example:"123.45"`
}

// TransactionResponse is the response struct for transaction creation
type TransactionResponse struct {
	TransactionID   uint    `json:"transaction_id" example:"1"`
	AccountID       uint    `json:"account_id" example:"1"`
	OperationTypeID uint    `json:"operation_type_id" example:"4"`
	Amount          float64 `json:"amount" example:"123.45"`
}

// NewTransactionHandler initializes a new transaction handler
func NewTransactionHandler(service services.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		transactionService: service,
	}
}

// CreateTransaction creates a new transaction.
// @Summary Create a new transaction
// @Description Creates a transaction with the provided account ID, operation type ID, and amount.
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param transaction body TransactionRequest true "Transaction request"
// @Success 201 {object} TransactionResponse "transaction_id, account_id, operation_type_id, and amount of the created transaction"
// @Failure 400 {object} ErrorResponse400 "Invalid request or failed transaction creation"
// @Router /transactions [post]
func (s *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	var req TransactionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(ErrorResponse{Error: ErrInvalidRequestParams})
	}

	transaction, err := s.transactionService.CreateTransaction(req.AccountID, req.OperationTypeID, req.Amount)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(ErrorResponse{Error: err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(TransactionResponse{
		TransactionID:   transaction.TransactionID,
		AccountID:       transaction.AccountID,
		OperationTypeID: transaction.OperationTypeID,
		Amount:          transaction.Amount,
	})
}
