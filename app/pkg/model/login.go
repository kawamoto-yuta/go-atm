package model

import (
	"gorm.io/gorm"
)

type LoginUser struct {
	gorm.Model
	AccountNum string `json:"account_num"`
	Password   string `json:"password"`
}

func NewLoginUser(account_num string, password string) *LoginUser {
	return &LoginUser{
		AccountNum: account_num,
		Password:   password,
	}
}

func GetUserByAccountNum(db *gorm.DB, AccountNum string) (*User, error) {
	user := User{}
	result := db.Where("account_num = ?", AccountNum).Find(&user)

	return &user, result.Error
}
