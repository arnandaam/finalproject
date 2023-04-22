package app

import (
	"finalMygram/model"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (

	host     string
	user     string
	password string
	dbPort   string
	dbname   string
	DB       *gorm.DB
	err      error
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	host = os.Getenv("DB_HOST")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbPort = os.Getenv("DB_PORT")
	dbname = os.Getenv("DB_NAME")
}

func StartDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}

	DB = db

	db.AutoMigrate(&model.User{}, &model.Comment{}, &model.Photo{}, &model.SocialMedia{})
	fmt.Println("DB connected!!!")
}

func GetDB() *gorm.DB {
	return DB
}
