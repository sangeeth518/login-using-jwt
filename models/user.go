package models

type User struct {
	ID       int    `json:"id" gorm:"primarykey"`
	Email    string `json:"email" gorm:"column:email;unique"`
	Password string `json:"_" gorm:"column:password"`
}
