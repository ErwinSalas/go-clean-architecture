package pkg

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Datastore interface {
	CreateAccount(context.Context, Account) error
	DeleteAccount(context.Context, int) (error, *Account)
	UpdateAccount(context.Context, Account) error
	GetAccountByID(context.Context, int) (error, *Account)
}

type PostgresDatastore struct {
	database *sql.DB
}

func NewPostgresDatastore(config *Config) (*PostgresDatastore, error) {
	conStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.DBUser, config.DBPassword, config.DBName, config.DBHost, config.DBPort,
	)
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, err
	}
	return &PostgresDatastore{
		database: db,
	}, nil
}

func (pds *PostgresDatastore) CreateAccount(context.Context, Account) error {
	return nil
}
func (pds *PostgresDatastore) DeleteAccount(context.Context, int) (error, *Account) {
	return nil, nil
}
func (pds *PostgresDatastore) UpdateAccount(context.Context, Account) error {
	return nil
}
func (pds *PostgresDatastore) GetAccountByID(context.Context, int) (error, *Account) {
	return nil, nil
}
