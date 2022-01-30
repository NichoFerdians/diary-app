package db

import (
	"fmt"
	"log"
	"os"

	"diary-app-service/entity"
	"diary-app-service/helper/auth"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("No .env file found")
	// }

	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_SERVICE")
	PORT := os.Getenv("DB_PORT")
	NAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, NAME)

	fmt.Println(URL)
	database, err := gorm.Open(mysql.Open(URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	database.AutoMigrate(&entity.User{})
	database.AutoMigrate(&entity.Diary{})

	hashedPassword, err := auth.HashPassword("123456")
	if err != nil {
		log.Fatal(err.Error())
	}

	newUser := &entity.User{
		Username: "admin",
		Password: hashedPassword,
		Name:     "Admin",
		Email:    "admin@admin.com",
	}
	database.Create(newUser)

	return database
}
