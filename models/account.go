package models

type Account struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	DocumentNumber string `gorm:"unique;not null" json:"document_number"`
}
