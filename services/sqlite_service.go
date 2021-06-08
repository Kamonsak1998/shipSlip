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

func (sqlite *Sqlite) Insert(data interface{}) error {
	res := sqlite.db.Create(data)
	if res.Error != nil {
		return res.Error
	}
	log.Printf("%+v", res.RowsAffected)
	return nil
}

func (sqlite *Sqlite) QueryAll(data interface{}) error {
	res, err := sqlite.db.Find(data).Rows()
	if err != nil {
		return err
	}
	for res.Next() {
		sqlite.db.ScanRows(res, &data)
	}
	return nil
}
