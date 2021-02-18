package database

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3" //driver
)

var (
	//ErrEmtpyField if message or name is empty
	ErrEmtpyField error = errors.New("Empty field")
)
//SQLite struct that allows to work with sqlite methods
type SQLite struct {
	DSN string
	DB  *sql.DB
}

//MakeSQlite returns sqlite
func MakeSQlite(DSN string) (*SQLite, error) {
	db, err := sql.Open("sqlite3", DSN)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &SQLite{DSN, db}, nil
}

//TestHandler writes data from test/topic
func (s *SQLite) TestHandler(message, name string) error {
	_, err := s.DB.Exec("CREATE TABLE IF NOT EXISTS Test (id INTEGER PRIMARY KEY AUTOINCREMENT, message TEXT NOT NULL, name TEXT NOT NULL)")
	if err != nil {
		return err
	}
	if message == "" || name == "" {
		return ErrEmtpyField
	}
	_, err = s.DB.Exec("INSERT INTO Test(message, name) VALUES(?, ?)", message, name)
	if err != nil {
		return err
	}
	return nil
}
