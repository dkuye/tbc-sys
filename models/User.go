package models

import (
	"cg-pkg/database"
	"cg-pkg/helper"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID               int       `gorm:"primary_key"`
	FirstName        string    `gorm:"not null"`
	LastName         string    `gorm:"not null"`
	Email            string    `gorm:"not null;unique"`
	Phone            string    `gorm:"not null"`
	Password         string    `sql:"DEFAULT:null"`
	Gender           string    `gorm:"not null"`
	MaritalStatus    string    `gorm:"not null"`
	DateOfBirth      string    `sql:"DEFAULT:null"`
	MembershipStatus string    `sql:"DEFAULT:null"`
	MembershipDate   time.Time `sql:"DEFAULT:null"`
	VisitorChurch    string    `sql:"DEFAULT:null"`
	VisitationDate   time.Time `sql:"DEFAULT:null"`
	RoleID           int       `gorm:"index;DEFAULT:0"`
	ActiveEmail      bool      `sql:"DEFAULT:null"`
	ActivePhone      bool      `sql:"DEFAULT:null"`
	Address          string    `sql:"DEFAULT:null"`
	Lga              string    `sql:"DEFAULT:null"`
	State            string    `sql:"DEFAULT:null"`
	DisplayPicture   string    `sql:"DEFAULT:null"`
	Occupation       string    `sql:"DEFAULT:null"`
	Education        string    `sql:"DEFAULT:null"`
	ApiToken         string    `sql:"DEFAULT:null"`
	RememberToken    string    `sql:"DEFAULT:null"`
	Uuid             string    `gorm:"unique"json:"uuid"`
	TimeStamp
}

type TimeStamp struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

const password = "$2a$14$qRiTUe/jeBYhey.LsKaXE.M5LHqdK2PDkeMKZ8hqWVD2roKhupkji"

func UserSeeder() {
	db := database.Connect()
	defer db.Close()
	db.Create(&User{
		FirstName:        "Sys",
		LastName:         "User",
		Email:            "sys@tbcoutofzion.com",
		Phone:            "08011223344",
		Password:         password,
		Gender:           "male",
		MaritalStatus:    "married",
		MembershipStatus: "Member",
		Uuid:             helper.Unid(),
	})

}

func (u User) Create(db *gorm.DB) User {
	var existingUser User
	db.Where(User{Email: u.Email}).Find(&existingUser)
	if existingUser.Email != "" {
		return existingUser
	}
	newUser:= User{
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		Email:         u.Email,
		Phone:         u.Phone,
		Gender:        u.Gender,
		MaritalStatus: u.MaritalStatus,
		Uuid:          helper.Unid(),
	}
	db.Create(&newUser)
	return newUser
}
