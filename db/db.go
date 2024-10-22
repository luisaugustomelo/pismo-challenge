package db

import (
	"database/sql"
	"fmt"
	"github.com/luisaugustomelo/pismo-challenge/utils/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func Connect() {
	// Load environment variables
	config.LoadEnv()

	// Set up DSN using environment variables
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USER,
		config.DB_PASS,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	fmt.Println("Database connection established")
}

func CloseConnection() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			log.Fatalf("Close database connection error: %v", err)
		}
	}(sqlDB)
}
