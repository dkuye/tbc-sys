package models

import (
	"time"
)

type User struct {
	ID               int       `gorm:"primary_key"`
	FirstName        string    `gorm:"not null"`
	LastName         string    `gorm:"not null"`
	Email            string    `gorm:"not null;unique"`
	Phone            string    `gorm:"not null;unique"`
	Password         string    `gorm:"not null"`
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
	TimeStamp
}

type TimeStamp struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
