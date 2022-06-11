package repo

import (
	"github.com/jmoiron/sqlx"
)

type Core struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Core {
	return &Core{
		db: db,
	}
}

type Config struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}
