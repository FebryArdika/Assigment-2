package database

import (
	"Assignment-2-Golang/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading .env file: ", er)
	}
	dbUser, dbPassword, dbPort, dbName, dbHost :=
		os.Getenv("postgres"),
		os.Getenv("123"),
		os.Getenv("5432"),
		os.Getenv("assigment-2"),
		os.Getenv("localhost")
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbPort, dbName)
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	db.Debug().AutoMigrate(models.Orders{}, models.Items{})
}

func GetDB() *gorm.DB {
	return db
}
