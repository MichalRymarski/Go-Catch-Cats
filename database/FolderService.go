package database

import (
	"gorm.io/gorm"
)

type Folder struct {
	gorm.Model
	Name string `json:"name" gorm:"not null, unique"`
}

func AddFolder(folderName string) error {
	folder := Folder{Name: folderName}
	err := db.Create(&folder).Error

	return err
}

func DeleteFolder(folderName string) error {
	err := db.Where("name =?", folderName).Delete(&Folder{}).Error

	return err
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
