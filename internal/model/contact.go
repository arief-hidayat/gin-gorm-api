package model

type Contact struct {
	ID          uint   `gorm:"primary_key:auto_increment" json:"id"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Email       string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	MobileNo    string `gorm:"type:varchar(255)" json:"mobile_no"`
	Institution string `gorm:"type:varchar(255)" json:"institution"`
}
