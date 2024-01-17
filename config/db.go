package config

import (
	"joe-l-mathew/Tambula-Ticket-Generator/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("tambula.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	DB = db

	DB.AutoMigrate(
		&models.Ticket{},
		&models.Game{},
	)
}
