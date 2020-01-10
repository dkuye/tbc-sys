package models

type Payment struct {
	ID             int     `gorm:"primary_key"`
	Ref            string  `gorm:"unique"json:"ref"`
	UserID         int     `gorm:"index;DEFAULT:0"json:"user_id"`
	ProgramID      int     `gorm:"index;DEFAULT:0"json:"program_id"`
	Amount         float64 `gorm:"index;not null"json:"amount"`
	Status         int     `gorm:"index;DEFAULT:0"json:"status"`
	GatewayStatus  string  `gorm:"index"sql:"DEFAULT:null"json:"gateway_status"`
	GatewayMessage string  `sql:"DEFAULT:null"json:"gateway_message"`
	GatewayChannel string  `sql:"DEFAULT:null"json:"gateway_channel"`
	Uuid           string  `gorm:"unique"json:"uuid"`
	TimeStamp
}
