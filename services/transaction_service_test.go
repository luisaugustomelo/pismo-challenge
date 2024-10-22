package services

import (
	"github.com/luisaugustomelo/pismo-challenge/services/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
)

func TestCreateTransactionSuccess(t *testing.T) {
	mockDB := new(mocks.MockDB)
	transactionService := NewTransactionService(mockDB)

	accountID := uint(1)
	operationTypeID := uint(1)
	amount := 100.50

	mockDB.On("Create", mock.AnythingOfType("*models.Transaction")).Return(&gorm.DB{Error: nil})

	createdTransaction, err := transactionService.CreateTransaction(accountID, operationTypeID, amount)

	assert.Nil(t, err)
	assert.NotNil(t, createdTransaction)
	assert.Equal(t, accountID, createdTransaction.AccountID)
	assert.Equal(t, operationTypeID, createdTransaction.OperationTypeID)
	assert.Equal(t, amount, createdTransaction.Amount)

	mockDB.AssertCalled(t, "Create", mock.AnythingOfType("*models.Transaction"))
}

func TestCreateTransactionFailure(t *testing.T) {
	mockDB := new(mocks.MockDB)
	transactionService := NewTransactionService(mockDB)

	accountID := uint(1)
	operationTypeID := uint(1)
	amount := 100.50

	mockDB.On("Create", mock.AnythingOfType("*models.Transaction")).Return(&gorm.DB{Error: gorm.ErrInvalidData})

	createdTransaction, err := transactionService.CreateTransaction(accountID, operationTypeID, amount)

	assert.NotNil(t, err)
	assert.Nil(t, createdTransaction)

	mockDB.AssertCalled(t, "Create", mock.AnythingOfType("*models.Transaction"))
}
