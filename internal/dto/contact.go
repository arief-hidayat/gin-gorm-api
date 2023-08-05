package dto

type Contact struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	MobileNo    string `json:"mobile_no" binding:"required,e164"`
	Institution string `json:"institution" binding:"required"`
}
