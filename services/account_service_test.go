package services

import (
	"errors"
	"github.com/luisaugustomelo/pismo-challenge/models"
	"github.com/luisaugustomelo/pismo-challenge/services/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
)

func TestCreateAccountSuccess(t *testing.T) {
	mockDB := new(mocks.MockDB)
	accountService := NewAccountService(mockDB)

	documentNumber := "12345678900"

	mockDB.On("First", mock.AnythingOfType("*models.Account"), "document_number = ?", documentNumber).Return(&gorm.DB{Error: gorm.ErrRecordNotFound})

	mockDB.On("Create", mock.AnythingOfType("*models.Account")).Return(&gorm.DB{Error: nil})

	createdAccount, err := accountService.CreateAccount(documentNumber)

	assert.Nil(t, err)
	assert.NotNil(t, createdAccount)
	assert.Equal(t, documentNumber, createdAccount.DocumentNumber)

	mockDB.AssertCalled(t, "First", mock.AnythingOfType("*models.Account"), "document_number = ?", documentNumber)
	mockDB.AssertCalled(t, "Create", mock.AnythingOfType("*models.Account"))
}

func TestCreateAccountAlreadyExists(t *testing.T) {
	mockDB := new(mocks.MockDB)
	accountService := NewAccountService(mockDB)

	documentNumber := "12345678900"

	mockDB.On("First", mock.AnythingOfType("*models.Account"), "document_number = ?", documentNumber).Return(&gorm.DB{Error: nil})

	createdAccount, err := accountService.CreateAccount(documentNumber)

	assert.NotNil(t, err)
	assert.Nil(t, createdAccount)
	assert.Equal(t, "account with this document number already exists", err.Error())

	mockDB.AssertCalled(t, "First", mock.AnythingOfType("*models.Account"), "document_number = ?", documentNumber)
	mockDB.AssertNotCalled(t, "Create", mock.AnythingOfType("*models.Account"))
}

func TestCreateAccountDatabaseError(t *testing.T) {
	mockDB := new(mocks.MockDB)
	accountService := NewAccountService(mockDB)

	documentNumber := "12345678900"

	mockDB.On("First", mock.AnythingOfType("*models.Account"), "document_number = ?", documentNumber).Return(&gorm.DB{Error: errors.New("database error")})

	createdAccount, err := accountService.CreateAccount(documentNumber)

	assert.NotNil(t, err)
	assert.Nil(t, createdAccount)
	assert.Equal(t, "database error", err.Error())

	mockDB.AssertCalled(t, "First", mock.AnythingOfType("*models.Account"), "document_number = ?", documentNumber)
	mockDB.AssertNotCalled(t, "Create", mock.AnythingOfType("*models.Account"))
}

func TestGetAccountSuccess(t *testing.T) {
	mockDB := new(mocks.MockDB)
	accountService := NewAccountService(mockDB)

	accountID := uint(1)
	documentNumber := "12345678900"

	expectedAccount := &models.Account{ID: accountID, DocumentNumber: documentNumber}
	mockDB.On("First", mock.AnythingOfType("*models.Account"), accountID).Return(&gorm.DB{Error: nil}).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*models.Account)
		*arg = *expectedAccount
	})

	account, err := accountService.GetAccount(accountID)

	assert.Nil(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, accountID, account.ID)
	assert.Equal(t, documentNumber, account.DocumentNumber)

	mockDB.AssertCalled(t, "First", mock.AnythingOfType("*models.Account"), accountID)
}

func TestGetAccountNotFound(t *testing.T) {
	mockDB := new(mocks.MockDB)
	accountService := NewAccountService(mockDB)

	accountID := uint(1)

	mockDB.On("First", mock.AnythingOfType("*models.Account"), accountID).Return(&gorm.DB{Error: gorm.ErrRecordNotFound})

	account, err := accountService.GetAccount(accountID)

	assert.NotNil(t, err)
	assert.Nil(t, account)
	assert.Equal(t, "account not found", err.Error())

	mockDB.AssertCalled(t, "First", mock.AnythingOfType("*models.Account"), accountID)
}

func TestGetAccountError(t *testing.T) {
	mockDB := new(mocks.MockDB)
	accountService := NewAccountService(mockDB)

	accountID := uint(1)
	expectedError := errors.New("database error")

	mockDB.On("First", mock.AnythingOfType("*models.Account"), accountID).Return(&gorm.DB{Error: expectedError})

	account, err := accountService.GetAccount(accountID)

	// Asserts
	assert.NotNil(t, err)
	assert.Nil(t, account)
	assert.Equal(t, expectedError, err)

	mockDB.AssertCalled(t, "First", mock.AnythingOfType("*models.Account"), accountID)
}
