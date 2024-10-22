package models

import (
	"time"
)

type Transaction struct {
	TransactionID   uint      `gorm:"primaryKey"`
	AccountID       uint      `gorm:"not null"`
	OperationTypeID uint      `gorm:"not null"`
	Amount          float64   `gorm:"not null"`
	EventDate       time.Time `gorm:"type:timestamp;not null"`

	Account       Account       `gorm:"foreignKey:AccountID;references:ID"`
	OperationType OperationType `gorm:"foreignKey:OperationTypeID;references:ID"`
}
