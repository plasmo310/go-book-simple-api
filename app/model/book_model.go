package model

import "time"

type Book struct {
	Id        uint      `gorm:"primary_key"`
	Name      string    `gorm:"size:255"`
	Price     uint      `gorm:default:0`
	CreatedAt time.Time `gorm:default:CURRENT_TIMESTAMP`
	UpdatedAt time.Time `gorm:default:CURRENT_TIMESTAMP`
}
