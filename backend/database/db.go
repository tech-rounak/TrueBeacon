package database

import (
	"fmt"
	"log"

	"github.com/tech-rounak/true-beacon/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectToSQLite() error {
	d, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	db = d
	fmt.Println("-----------Database Connection Successfull-----------------")
	return nil
}

func InitDB() *gorm.DB {
	err := ConnectToSQLite()
	if err != nil {
		log.Fatal("Error while connecting db", err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Share{})
	return db
}
