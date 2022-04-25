package repository

// логика подключения бд + имена таблиц

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// названия таблиц из бд
const (
	usersTable      = "users"
	dataListsTable  = "data_lists"
	usersListsTable = "users_lists"
	dataItemsTable  = "data_items"
	listsItemsTable = "lists_items"
)

// параметры для подкл бд
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
