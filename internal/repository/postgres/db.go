package postgres

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
