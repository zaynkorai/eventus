package mysql

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// New creates new database connection to a mysql database
func New(dsn string, timeout int) (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		return nil, err
	}

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("🚀 Connected Successfully to the Database")

	return db, nil
}
