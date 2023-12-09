package store

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
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
	entries, _ := os.ReadDir("./")
	fmt.Println(entries)
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
