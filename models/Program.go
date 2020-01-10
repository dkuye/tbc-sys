package models

import "time"

type Program struct {
	ID        int       `gorm:"primary_key"`
	Name      string    `gorm:"not null"`
	Slug      string    `gorm:"not null"`
	Year      string    `gorm:"not null"`
	StartDate time.Time `sql:"DEFAULT:null"`
	EndDate   time.Time `sql:"DEFAULT:null"`
	TimeStamp
}
