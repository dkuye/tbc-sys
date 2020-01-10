package models

import (
	"cg-pkg/database"
	"cg-pkg/helper"
	"github.com/gosimple/slug"
	"time"
)

type Program struct {
	ID        int       `gorm:"primary_key"`
	Name      string    `gorm:"not null"`
	Slug      string    `gorm:"not null"`
	Year      string    `gorm:"not null"`
	StartDate time.Time `sql:"DEFAULT:null"`
	EndDate   time.Time `sql:"DEFAULT:null"`
	Uuid      string    `gorm:"unique"json:"uuid"`
	TimeStamp
}

func ProgramSeeder() {
	db := database.Connect()
	defer db.Close()
	db.Create(&Program{
		Name: "Voltage Plug",
		Slug: slug.Make("Voltage Plug"),
		Year: "2019",
		Uuid: helper.Unid(),
	})

}
