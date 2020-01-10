package models

type PasswordReset struct {
	ID          int 	`gorm:"primary_key"`
	Email       string  `gorm:"index;not null"`
	Token       string 	`gorm:"unique;not null"`
	Status       bool	`gorm:"index;DEFAULT:0"`
	TimeStamp
}
