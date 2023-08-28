package postgres

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	db *sqlx.DB
}

func New(dsn string) (*Database, error) {
	//db, err := sqlx.Open("postgres", dsn)
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		fmt.Println("DataBase NOT WORK")
		return nil, err
	}
	fmt.Println("DataBase Work")
	return &Database{
		db: db,
	}, nil
}
