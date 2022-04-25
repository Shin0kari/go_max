package repository

import (
	"fmt"

	serv "github.com/Shin0kari/go_max"
	"github.com/jmoiron/sqlx"
)

type DataItemPostgres struct {
	db *sqlx.DB
}

func NewDataItemPostgres(db *sqlx.DB) *DataItemPostgres {
	return &DataItemPostgres{db: db}
}

func (r *DataItemPostgres) Create(listId int, item serv.DataItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id",
		dataItemsTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)",
		listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

// sql запрос для Select
func (r *DataItemPostgres) GetAll(userId, listId int) ([]serv.DataItem, error) {
	var items []serv.DataItem
	query := fmt.Sprintf("SELECT * ti.id, ti.title, ti.description, ti.done FROM %s INNER JOIN %s li on li.item_id = ti.id INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2",
		dataItemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}
