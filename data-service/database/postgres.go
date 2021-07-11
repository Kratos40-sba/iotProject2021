package database

import (
	"fmt"
	"github.com/Kratos40-sba/data-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

const (
	dbName = "POSTGRES_DB_NAME"
	dbUser = "POSTGRES_USER_NAME"
	dbPort = "POSTGRES_PORT"
)

func DbConnection() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: fmt.Sprintf("user=%v dbname=%v port=%v sslmode=disable", os.Getenv(dbUser), os.Getenv(dbName), os.Getenv(dbPort))}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatalln("Failed to connect to database!")
	}
	db.AutoMigrate(&models.DhtEvent{})
	return db
}
