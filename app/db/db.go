package db

import (
	"fmt"
	"os"
	"user-api/app/zapgorm2"
	"user-api/app/zaplogger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DBError error

func InitDB() {
	dbString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable timezone=Asia/Bangkok",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)
	zaplog := zapgorm2.New(zaplogger.Log)
	zaplog.SetAsDefault()
	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
		Logger: zaplog,
	})

	DB = db
	DBError = err
}
