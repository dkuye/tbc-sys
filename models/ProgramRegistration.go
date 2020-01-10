package models

type ProgramRegistration struct {
	ID        int `gorm:"primary_key"`
	ProgramID int `gorm:"not null"`
	UserID    int `gorm:"not null"`
	TimeStamp
}
