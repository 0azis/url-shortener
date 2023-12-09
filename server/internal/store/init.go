package store

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type InterfaceStore interface {
	Open()
	Close()
	Users() interfaceUser
	Urls() interfaceUrl
}

type store struct {
	db *sql.DB
}

func (s *store) Open() {
	db, _ := sql.Open("sqlite3", "./url-shortener.db")

	s.db = db
}

func (s *store) Close() {
	s.db.Close()
}

func (s *store) Users() interfaceUser {
	return &user{db: s.db}
}

func (s *store) Urls() interfaceUrl {
	return &url{db: s.db}
}

func GetStoreInstance() InterfaceStore {
	return &store{}
}
