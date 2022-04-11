//заготовки интерфейсов для сущностей
package service

import (
	serv "github.com/Shin0kari/go_max"
	"github.com/Shin0kari/go_max/package/repository"
)

type Authorization interface {
	// возвращает id созданного пользователя
	CreateUser(user serv.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type DataList interface {
}

type DataItem interface {
}

// собираем все интерфейсы в 1-ом месте
type Service struct {
	Authorization
	DataList
	DataItem
}

// конструктор для сервисов
// сервисы обращаются к базе данных и поэтому объявляем указатель на репозиторий
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
