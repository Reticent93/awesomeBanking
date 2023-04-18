package repository

import (
	"database/sql"

	"github.com/Reticent93/awesomeBanking/internal/models"
)

type Repository interface {
	Connection() *sql.DB
	CreateUser(user *models.User) error
	CreateAccount(account *models.Account) error
}