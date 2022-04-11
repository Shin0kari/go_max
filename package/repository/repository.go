package repository

import (
	serv "github.com/Shin0kari/go_max"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user serv.User) (int, error)
	GetUser(username, password string) (serv.User, error)
}

type DataList interface {
}

type DataItem interface {
}

type Repository struct {
	Authorization
	DataList
	DataItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
