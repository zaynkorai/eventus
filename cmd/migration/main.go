package main

import (
	"log"
	"os"

	"github.com/zaynkorai/eventus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(mysql.Open("eventus:eventus123@tcp(localhost:3306)/eventus?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	db.AutoMigrate(&eventus.Event{})

	log.Println("🚀 Connected Successfully to the Database")
}
