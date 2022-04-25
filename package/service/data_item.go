package service

import (
	serv "github.com/Shin0kari/go_max"
	"github.com/Shin0kari/go_max/package/repository"
)

type DataItemService struct {
	repo     repository.DataItem
	ListRepo repository.DataList
}

func NewDataItemService(repo repository.DataItem, listRepo repository.DataList) *DataItemService {
	return &DataItemService{repo: repo, ListRepo: listRepo}
}

func (s *DataItemService) Create(userId, listId int, item serv.DataItem) (int, error) {
	_, err := s.ListRepo.GetById(userId, listId)
	if err != nil {
		// list does not exists or does not belongs to var
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *DataItemService) GetAll(userId, listId int) ([]serv.DataItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *DataItemService) GetById(userId, itemId int) (serv.DataItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *DataItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *DataItemService) Update(userId, itemId int, input serv.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}
