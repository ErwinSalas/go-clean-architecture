package pkg

import (
	"context"
)

type Service interface {
	GetAccountByID(context.Context, int) (error, *Account)
}

type AccountService struct {
	datastore Datastore
}

func NewAccountService(datastore Datastore) *AccountService {
	return &AccountService{
		datastore: datastore,
	}
}

func (s *AccountService) GetAccountByID(context.Context, int) (error, *Account) {
	return nil, nil
}
