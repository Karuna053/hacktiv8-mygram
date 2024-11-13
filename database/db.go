package database

import (
	"fmt"
	"mygram/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBHost = "localhost"
	DBUser = "postgres"
	DBPass = "password"
	DBName = "hacktivate-mygram"
	DBPort = "5432"

	DB *gorm.DB // Declare data type.
)

func InitDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		DBHost, DBUser, DBPass, DBName, DBPort,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// DB.AutoMigrate(&models.Item{}, &models.Order{})
	DB.Debug().AutoMigrate(
		&models.Comment{},
		&models.Photo{},
		&models.SocialMedia{},
		&models.User{},
	)
}

func GetDB() *gorm.DB {
	return DB
}
