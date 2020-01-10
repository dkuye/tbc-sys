package models

import (
	"cg-pkg/database"
	"cg-pkg/helper"
	"github.com/gosimple/slug"
)

type Role struct {
	ID          int    `gorm:"primary_key"`
	Name        string `gorm:"not null"`
	Slug        string `gorm:"not null;unique"`
	LongName    string `sql:"DEFAULT:null"`
	Description string `sql:"DEFAULT:null"`
	Uuid        string `gorm:"unique"json:"uuid"`
	TimeStamp
	Users []User
}

func RoleSeeder() {
	db := database.Connect()
	defer db.Close()
	db.Create(&Role{
		Name:     "Admin",
		Slug:     slug.Make("admin"),
		LongName: "Administrators",
		Uuid:     helper.Unid(),
	})
	db.Create(&Role{
		Name:     "Member",
		Slug:     slug.Make("member"),
		LongName: "Members",
		Uuid:     helper.Unid(),
	})
}
