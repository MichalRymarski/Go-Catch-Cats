package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func startDBConnection() {
	dsn := "host=localhost user=myuser password=mypassword dbname=mydb port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&CatUser{}, &Folder{})
	if err != nil {
		panic("failed to migrate database schema")
	}

	err = db.AutoMigrate(&CatUser{}, &Folder{})
	if err != nil {
		panic("failed to migrate database schema")
	}
}

func SetupDatabase() {
	if db == nil {
		startDBConnection()
	}
}
