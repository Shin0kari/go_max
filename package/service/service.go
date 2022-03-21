//заготовки интерфейсов для сущностей
package service

import "github.com/Shin0kari/go_max/package/repository"

type Authorization interface {
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
func NewService(rep *repository.Repository) *Service {
	return &Service{}
}
