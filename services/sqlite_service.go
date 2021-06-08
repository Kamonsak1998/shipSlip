package services

import (
	"log"
	"shipSlip/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Sqlite struct {
	db *gorm.DB
}

func Connect(dbFile string) (Sqlite, error) {
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		return Sqlite{}, err
	}
	db.AutoMigrate(&models.Customers{})
	return Sqlite{db: db}, nil
}

func (sqlite *Sqlite) Insert(data *models.Customers) error {
	res := sqlite.db.Create(&data)
	if res.Error != nil {
		return res.Error
	}
	log.Printf("%+v", res.RowsAffected)
	return nil
}
