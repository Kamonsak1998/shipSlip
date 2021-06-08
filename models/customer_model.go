package models

type Customers struct {
	Name     string `gorm:"name"`
	District string `gorm:"district"`
	Province string `gorm:"province"`
	Sender   string `gorm:"sender"`
}
