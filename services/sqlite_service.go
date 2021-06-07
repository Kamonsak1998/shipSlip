package services

import (
	"database/sql"
	// _ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	db *sql.DB
}

// func Connect(dbFile string) (Sqlite, error) {
// 	con, err := sql.Open("sqlite3", dbFile)
// 	if err != nil {
// 		return Sqlite{}, err
// 	}
// 	return Sqlite{db: con}, nil
// }

// func (sqlite *Sqlite) CreateTable() error {
// 	statement, err := sqlite.db.Prepare("CREATE TABLE IF NOT EXISTS customer (id INT NOT  NULL AUTO_INCREMENT, name TEXT, district TEXT, province TEXT, sender TEXT, PRIMARY KEY (id))")
// 	if err != nil {
// 		return err
// 	}
// 	statement.Exec()
// 	return nil
// }

// func (sqlite *Sqlite) Query() error {
// 	return nil
// }
