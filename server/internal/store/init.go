package store

import (
	"database/sql"
	_ "github.com/jackc/pgx/stdlib"
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
	db, _ := sql.Open("pgx", "postgres://notes_db_sebp_user:mGA7RJ0scHFknI60vZ6RGkQmaEovnTNU@dpg-clgdq9ef27hc739jfplg-a.frankfurt-postgres.render.com/notes_db_sebp")

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
