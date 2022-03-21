package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
