package service

import (
	serv "github.com/Shin0kari/go_max"
	"github.com/Shin0kari/go_max/package/repository"
)

type DataListService struct {
	repo repository.DataList
}

func NewDataListService(repo repository.DataList) *DataListService {
	return &DataListService{repo: repo}
}

func (s *DataListService) Create(userId int, list serv.DataList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *DataListService) GetAll(userId int) ([]serv.DataList, error) {
	return s.repo.GetAll(userId)
}

func (s *DataListService) GetById(userId, listId int) (serv.DataList, error) {
	return s.repo.GetById(userId, listId)
}
