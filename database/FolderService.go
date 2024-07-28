package database

import (
	"gorm.io/gorm"
)

type Folder struct {
	gorm.Model
	Id   uint64 `json:"id" gorm:"primarykey, autoincrement"`
	Name string `json:"name" gorm:"not null"`
}

func AddFolder(folderName string) error {
	folder := Folder{Name: folderName}

	return db.Create(&folder).Error
}

func GetFolderByName(folderName string) (Folder, error) {
	var folder Folder
	err := db.Where("name =?", folderName).First(&folder).Error

	return folder, err
}

func GetAllFolders() []Folder {
	var folders []Folder
	db.Find(&folders)

	return folders
}
