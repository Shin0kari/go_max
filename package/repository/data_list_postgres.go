package repository

import (
	"fmt"

	serv "github.com/Shin0kari/go_max"
	"github.com/jmoiron/sqlx"
)

type DataListPostgres struct {
	db *sqlx.DB
}

func NewDataListPostgres(db *sqlx.DB) *DataListPostgres {
	return &DataListPostgres{db: db}
}

func (r *DataListPostgres) Create(userId int, list serv.DataList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", dataListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *DataListPostgres) GetAll(userId int) ([]serv.DataList, error) {
	var lists []serv.DataList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		dataListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *DataListPostgres) GetById(userId, listId int) (serv.DataList, error) {
	var list serv.DataList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2", dataListsTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}
