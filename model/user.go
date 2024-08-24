package model

type UserModel struct {
	Id              uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name            string `json:"name"`
	IsAdministrator bool   `json:"isAdministrator"`
	DateField
}
