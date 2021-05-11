package model

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         string `json:"id"`
	Name       string `json:"name"`
	AccountNum string `json:"account_num"`
	Password   string `json:"password"`
}

func NewUser(id string, name string, account_num string, password string) *User {
	return &User{
		ID:         id,
		Name:       name,
		AccountNum: account_num,
		Password:   password,
	}
}

func CreateUser(db *gorm.DB, user *User) (*User, error) {
	result := db.Create(&user)

	return user, result.Error
}

func GetUsers(db *gorm.DB) ([]*User, error) {
	users := []*User{}
	result := db.Find(&users)
	fmt.Println("aa")

	return users, result.Error
}

func GetUserById(db *gorm.DB, ID string) (*User, error) {
	user := User{}
	result := db.First(&user, ID)

	return &user, result.Error
}

func (user *User) Update(db *gorm.DB, param map[string]interface{}) (*User, error) {
	result := db.Model(&user).Updates(param)

	return user, result.Error
}

func (user *User) Delete(db *gorm.DB) (*User, error) {
	result := db.Delete(&user)

	return user, result.Error
}
