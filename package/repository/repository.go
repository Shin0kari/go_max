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
	Create(userId int, list serv.DataList) (int, error)
	GetAll(userId int) ([]serv.DataList, error)
	GetById(userId, listId int) (serv.DataList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input serv.UpdateListInput) error
}

type DataItem interface {
	Create(listId int, item serv.DataItem) (int, error)
	GetAll(userId, listId int) ([]serv.DataItem, error)
	GetById(userId, itemId int) (serv.DataItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input serv.UpdateItemInput) error
}

type Repository struct {
	Authorization
	DataList
	DataItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		DataList:      NewDataListPostgres(db),
		DataItem:      NewDataItemPostgres(db),
	}
}
