package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

type Config struct {
	Host     string
	Port     string
	UserName string
	Password string
	DbName   string
	SslMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.UserName, cfg.Password, cfg.DbName, cfg.SslMode))

	log.Default().Printf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.UserName, cfg.Password, cfg.DbName, cfg.SslMode)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
