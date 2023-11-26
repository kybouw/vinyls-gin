package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type album struct {
	gorm.Model
	Title        string
	Artist       string
	YearReleased uint
}

func ConnectDatabase() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&album{})
	if err != nil {
		return nil
	}

	return database
}
