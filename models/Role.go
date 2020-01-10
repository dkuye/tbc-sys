package models

type Role struct {
	ID        	int 	`gorm:"primary_key"`
	Name        string	`gorm:"not null"`
	Slug        string	`gorm:"not null;unique"`
	LongName    string
	Description string
	Status      bool	`gorm:"index"`
	TimeStamp

	Users []User
}

