package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
"github.com/joho/godotenv"
	"github.com/Reticent93/awesomeBanking/helpers"
	"github.com/Reticent93/awesomeBanking/internal/models"
	"github.com/Reticent93/awesomeBanking/internal/repository"
)

type PostgresDBRepo struct {
	//add a connection pool
	DB *sql.DB
}

type application struct {
	DSN	 string
	DB repository.Repository
	ErrorLog *log.Logger
	InfoLog *log.Logger
}


func(p *PostgresDBRepo) Connection() *sql.DB {
	return p.DB
}

var app application


func NewDatabase() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	  }

	PG_HOST := os.Getenv("PG_HOST")
	PG_PORT := os.Getenv("PG_PORT")
	PG_NAME := os.Getenv("PG_NAME")
	PG_USER := os.Getenv("PG_USER")
	PG_PASS := os.Getenv("PG_PASS")

	  
	log.Println("Connecting to database...")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", PG_HOST, PG_PORT,PG_USER, PG_PASS, PG_NAME  )

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err  
	}
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\r", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog
	
	return db, nil
}


func(p *PostgresDBRepo) connectToDB() (*sql.DB, error) {
	db, err := NewDatabase()
	if err != nil {
		return nil, err
	}

	return db, nil
}


const dbTimeOut = time.Second * 10

func(p *PostgresDBRepo) CreateUser(user *models.User) error {
	//create a new user account
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	//create a new user account
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
	_, err := p.DB.ExecContext(ctx, query, user.Username, user.Email, user.Password)
	helpers.HandleErr(err)

	return nil
}

func(p *PostgresDBRepo) CreateAccount(account *models.Account) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `INSERT INTO accounts (type, name, balance, user_id) VALUES ($1, $2, $3, $4)`
	_, err := p.DB.ExecContext(ctx, query, account.Type, account.Name, account.Balance, account.UserId)
	helpers.HandleErr(err)

	return nil
}

