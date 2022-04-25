package repository

import (
	"fmt"

	serv "github.com/Shin0kari/go_max"
	"github.com/jmoiron/sqlx"
)

// имплиминтирует интерфейс бд с репоз
type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// запрос создания пользователя
func (r *AuthPostgres) CreateUser(user serv.User) (int, error) {
	var id int
	// места с $ - это метса в которые будут переданы значения для выполнения запроса к бд
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id",
		usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (serv.User, error) {
	var user serv.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2",
		usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
