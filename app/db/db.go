package db

import (
	"book-simple-api/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	// settings
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbHost := "mysql" // db container name
	dbPort := "3306"

	// connect db
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	// migration
	err = AutoMigrate()
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func AutoMigrate() error {
	err := db.AutoMigrate(&models.Book{})
	if err != nil {
		return err
	}
	return nil
}
