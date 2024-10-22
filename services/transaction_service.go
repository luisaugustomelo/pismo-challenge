package services

import (
	"github.com/luisaugustomelo/pismo-challenge/interfaces"
	"github.com/luisaugustomelo/pismo-challenge/models"
	"time"
)

type TransactionService interface {
	CreateTransaction(accountId, operationType uint, amount float64) (*models.Transaction, error)
}

type transactionService struct {
	db interfaces.Database
}

func NewTransactionService(db interfaces.Database) TransactionService {
	return &transactionService{
		db: db,
	}
}

func (s *transactionService) CreateTransaction(accountId, operationType uint, amount float64) (*models.Transaction, error) {
	transaction := models.Transaction{
		AccountID:       accountId,
		OperationTypeID: operationType,
		Amount:          amount,
		EventDate:       time.Now().UTC(),
	}

	if err := s.db.Create(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}
