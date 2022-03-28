package repository

import (
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
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
	return &Repository{}
}
