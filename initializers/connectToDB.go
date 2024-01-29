package initializers

import (
	"log"
	"os"

	"github.com/sangeeth/jwt-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := os.Getenv("dsn")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to Db")
	}
	log.Println("database connected succesfully")

	DB.AutoMigrate(&models.User{})
}
