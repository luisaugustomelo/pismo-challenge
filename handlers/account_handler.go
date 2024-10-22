package handlers

import (
	"github.com/luisaugustomelo/pismo-challenge/services"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	accountService services.AccountService
}

// AccountResponse is the response struct for account creation
type AccountResponse struct {
	AccountID      uint   `json:"account_id" example:"1"`
	DocumentNumber string `json:"document_number" example:"12345678900"`
}

func NewAccountHandler(accountService services.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

// CreateAccount creates a new account.
// @Summary Create a new account
// @Description Creates a new account with the provided document number.
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param document_number body string true "Document Number" example("12345678900")
// @Success 201 {object} AccountResponse "account_id and document_number of the created account"
// @Failure 400 {object} ErrorResponse "Invalid request or account already exists"
// @Router /accounts [post]
func (s *AccountHandler) CreateAccount(c *fiber.Ctx) error {
	type request struct {
		DocumentNumber string `json:"document_number"`
	}

	var req request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(ErrorResponse{Error: ErrInvalidRequestParams})
	}

	account, err := s.accountService.CreateAccount(req.DocumentNumber)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(ErrorResponse{Error: err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(AccountResponse{
		AccountID:      account.ID,
		DocumentNumber: account.DocumentNumber,
	})
}

// GetAccount retrieves an account by account ID.
// @Summary Get an account by ID
// @Description Retrieves the account details using the account ID.
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID" example(1)
// @Success 200 {object} AccountResponse "account_id and document_number of the retrieved account"
// @Failure 400 {object} ErrorResponse400 "Invalid accountId provided"
// @Failure 404 {object} ErrorResponse404 "Account not found"
// @Router /accounts/{id} [get]
func (s *AccountHandler) GetAccount(c *fiber.Ctx) error {
	accountIdParam := c.Params("id")
	accountId, err := strconv.ParseUint(accountIdParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{Error: ErrInvalidAccountIDParam})
	}

	account, err := s.accountService.GetAccount(uint(accountId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{Error: ErrAccountNotFound})
	}

	return c.Status(fiber.StatusOK).JSON(AccountResponse{
		AccountID:      account.ID,
		DocumentNumber: account.DocumentNumber,
	})
}
