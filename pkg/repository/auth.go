package repository

import "github.com/jmoiron/sqlx"

type AuthRepository struct {
	db *sqlx.DB
}
