package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var DB *gorm.DB

func Connect() {
	dsn := "postgres://postgres:123456@localhost:5432/postgres"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
		Logger:  logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("[DATABASE]::CONNECTION_ERROR")
		panic(err)
	}

	DB = db

	log.Print("Successfully connected to database")

	DB.Logger = logger.Default.LogMode(logger.Info)
	fmt.Println("[DATABASE]::CONNECTED")
}

func Migrate(tables ...interface{}) error {
	return DB.AutoMigrate(tables...)
}
