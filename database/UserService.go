package database

import (
	"gorm.io/gorm"
)

type CatUser struct {
	gorm.Model
	Id       uint64 `json:"id" gorm:"primarykey,autoincrement"`
	NickName string `json:"nickname" gorm:"not null"`
	Email    string `json:"email" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}

func AddUser(user *CatUser) error {

	return db.Create(user).Error
}

func GetUserById(id uint64) (CatUser, error) {
	var user CatUser
	err := db.First(&user, id).Error

	return user, err
}

func GetAllUsers() []CatUser {
	var users []CatUser
	db.Find(&users)

	return users
}
