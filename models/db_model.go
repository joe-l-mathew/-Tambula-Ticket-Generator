package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	Tickets []Ticket `gorm:"foreignKey:GameID"`
}

type Ticket struct {
	gorm.Model
	GameID     uint `gorm:"index"`
	MatrixData string
}
