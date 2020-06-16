package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
type Datastore interface {
	AllDataInterface() ([]*Test, error)
}

type DB struct {
	*sql.DB
}

func MakeDb(dataSourceName string) (*DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}