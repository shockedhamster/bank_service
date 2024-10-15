package service

import (
	"github.com/bank_service/internal/entities"
	"github.com/bank_service/internal/repository"
	"github.com/sirupsen/logrus"
)

type OperationService struct {
	repo repository.Operations
}

func NewOperationsService(repo repository.Operations) *OperationService {
	return &OperationService{repo: repo}
}

func (s *OperationService) GetUserBalanceById(id int) (int, error) {
	balance, err := s.repo.GetUserBalanceById(id)
	if err != nil {
		logrus.Errorf("Error getting a balance: %s", err.Error())
		return 0, err
	}
	return balance, nil
}

func (s *OperationService) GetTransactionHistoryById(id int, sortType string, limit, offset int) ([]entities.Operation, error) {
	resultTransactionList, err := s.repo.GetTransactionHistoryById(id, sortType, limit, offset)
	if err != nil {
		logrus.Errorf("Error getting a transaction list: %s", err.Error())
		return nil, err
	}

	return resultTransactionList, err

}
