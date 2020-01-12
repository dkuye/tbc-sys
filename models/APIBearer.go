package models

import (
	"github.com/dkuye/database"
	"github.com/dkuye/helper"
	"os"
)

type APIBearer struct {
	ID     int    `gorm:"primary_key"`
	Name   string `gorm:"not null;unique"`
	Key    string `gorm:"not null;unique"`
	Status int    `gorm:"index"`
	Uuid   string `gorm:"unique"json:"uuid"`
	TimeStamp
}

func APIBearerSeeder() {
	db := database.Connect()
	defer db.Close()

	db.Create(&APIBearer{
		Name:   "Sys",
		Key:    os.Getenv("API_KEY"),
		Status: 1,
		Uuid:   helper.Uuid(),
	})
}
