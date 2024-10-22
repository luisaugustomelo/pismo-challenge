package services

import (
	"errors"
	"github.com/luisaugustomelo/pismo-challenge/interfaces"
	"github.com/luisaugustomelo/pismo-challenge/models"
	"gorm.io/gorm"
)

type AccountService interface {
	CreateAccount(documentNumber string) (*models.Account, error)
	GetAccount(accountId uint) (*models.Account, error)
}

type accountService struct {
	db interfaces.Database
}

func NewAccountService(db interfaces.Database) AccountService {
	return &accountService{
		db: db,
	}
}

func (s *accountService) CreateAccount(documentNumber string) (*models.Account, error) {
	var account models.Account

	err := s.db.First(&account, "document_number = ?", documentNumber).Error
	if err == nil {
		return nil, errors.New("account with this document number already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	account = models.Account{
		DocumentNumber: documentNumber,
	}
	if err := s.db.Create(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}

func (s *accountService) GetAccount(accountId uint) (*models.Account, error) {
	var account models.Account

	result := s.db.First(&account, accountId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("account not found")
		}
		return nil, result.Error
	}

	return &account, nil
}
