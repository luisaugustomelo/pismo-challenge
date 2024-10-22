package migrations

import (
	"github.com/luisaugustomelo/pismo-challenge/db"
	"github.com/luisaugustomelo/pismo-challenge/models"
	"log"
)

func Migrate() {
	models.RegisterModels()

	err := db.DB.AutoMigrate(models.Models...)
	if err != nil {
		log.Fatalf("Error to do migration: %v", err)
	}
	log.Println("Migration was successful!")

	seedOperationTypes(db.DB)
}

func DropTable() {
	models.RegisterModels()

	err := db.DB.AutoMigrate(models.Models...)

	err = db.DB.Migrator().DropTable(models.Models...)
	if err != nil {
		log.Fatalf("Error to drop migration: %v", err)
	}
	log.Println("Drop table was successful!")
}
