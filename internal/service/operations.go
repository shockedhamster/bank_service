package service

import (
	"fmt"

	"github.com/bank_service/internal/repository"
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
		fmt.Println(err.Error())
		return 0, err
	}
	return balance, nil
}
