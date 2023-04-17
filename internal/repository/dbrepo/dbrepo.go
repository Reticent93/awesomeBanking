package dbrepo

import (
	"awesomeBanking/cmd/utils"
	"database/sql"
	"time"
	
)

type PostgresDBRepo struct {
	//add a connection pool
	DB *sql.DB
}

func(p *PostgresDBRepo) ConnectionDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5433/banking?sslmode=disable")
	utils.HandleErr(err)
	return db
}

const dbTimeOut = time.Second * 10

