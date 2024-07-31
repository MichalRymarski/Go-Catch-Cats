package database

import (
	"gorm.io/gorm"
)

type CatUser struct {
	gorm.Model
	NickName string `json:"nickname" gorm:"not null, unique"`
	Email    string `json:"email" gorm:"not null, unique"`
	Password string `json:"password" gorm:"not null"`
	Role     string `json:"role" gorm:"default:user"`
}

func AddExampleUser() error {
	user := CatUser{
		NickName: "asdaadadsd",
		Email:    "<EaMAIL>",
		Password: "<PaASSWORD>",
		Role:     "",
	}
	err := db.Create(&user).Error

	return err
}

func AddUser(user *CatUser) error {

	return db.Create(user).Error
}

func DeleteUser(userNickName string) error {
	err := db.Where("nick_name =?", userNickName).Delete(&CatUser{}).Error

	return err
}

func AddDefaultAdmin() error {
	user := CatUser{
		NickName: "admin",
		Email:    "<EMAIL>",
		Password: "<PASSWORD>",
		Role:     "admin",
	}
	err := db.Create(&user).Error

	return err
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
