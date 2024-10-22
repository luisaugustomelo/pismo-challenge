package migrations

import (
	"github.com/luisaugustomelo/pismo-challenge/models"
	"gorm.io/gorm"
)

func seedOperationTypes(db *gorm.DB) {
	operationTypes := []models.OperationType{
		{Description: "PURCHASE"},
		{Description: "INSTALLMENT PURCHASE"},
		{Description: "WITHDRAWAL"},
		{Description: "PAYMENT"},
	}

	for _, opType := range operationTypes {
		db.FirstOrCreate(&opType, models.OperationType{Description: opType.Description})
	}
}
