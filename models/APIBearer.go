package models

import (
	"cg-pkg/database"
	"os"
)

type APIBearer struct {
	ID     int    `gorm:"primary_key"`
	Name   string `gorm:"not null;unique"`
	Key    string `gorm:"not null;unique"`
	Status int   `gorm:"index"`
	TimeStamp
}

func APIBearerSeeder() {
	db := database.Connect()
	defer db.Close()

	db.Create(&APIBearer{
		Name:      "Sys",
		Key:       os.Getenv("API_KEY"),
		Status:    1,
	})
}
