package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	database *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		database: db,
	}
}