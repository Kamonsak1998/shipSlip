package models

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	Name     string `gorm:"name"`
	District string `gorm:"district"`
	Province string `gorm:"province"`
	Sender   string `gorm:"sender"`
}
