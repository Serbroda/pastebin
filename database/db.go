package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"pastebin/models"
	"sync"
)

var (
	dbConnection *gorm.DB
	once         sync.Once
)

type ConnectionOptions struct {
	Name string
}

func Connect(options ConnectionOptions) *gorm.DB {
	once.Do(func() {
		db, err := gorm.Open(sqlite.Open(options.Name), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database %s: %v", options.Name, err)
			panic(err)
		}
		db.AutoMigrate(&models.Paste{})
		dbConnection = db
	})
	return dbConnection
}

func GetConnection() *gorm.DB {
	return dbConnection
}
