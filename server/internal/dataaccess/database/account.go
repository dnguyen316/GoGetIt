package database

import (
	"context"
	"log"

	"github.com/doug-martin/goqu/v9"
)

type Account struct {
	AccountId   uint64 `sql:"account_id"`
	AccountName string `sql:"accountName"`
}

type AccountDataAccessor interface {
	CreateAccount(ctx context.Context, account Account) (uint64, error)
	GetAccountByID(ctx context.Context, id uint64) (Account, error)
	GetAccountByAccountName(ctx context.Context, accountName string) (Account, error)
	WithDatabase(database Database) AccountDataAccessor
}

type accountDataAccessor struct {
	database Database
}

func NewAccountAccessor(
	database Database,
) AccountDataAccessor {
	return &accountDataAccessor{
		database: database,
	}
}

func (a accountDataAccessor) CreateAccount(ctx context.Context, account Account) (uint64, error) {
	// TODO: Implement create account logic
	result, err := a.database.Insert("accounts").Rows(goqu.Record{
		"accountName": account.AccountName,
	}).Executor().ExecContext(ctx)
	if err != nil {
		log.Printf("failed to create account, err=%+v", err)
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		log.Printf("Failed to get last inserted id, err=%+v\n", err)
		return 0, nil
	}

	return uint64(lastInsertId), nil
}

func (a *accountDataAccessor) GetAccountByID(ctx context.Context, id uint64) (Account, error) {
	// TODO: Implement get account by ID logic
	return Account{}, nil
}

func (a *accountDataAccessor) GetAccountByAccountName(ctx context.Context, accountName string) (Account, error) {
	// TODO: Implement get account by accountName logic
	return Account{}, nil
}

// WithDatabase implements AccountDataAccessor.
func (a *accountDataAccessor) WithDatabase(database Database) AccountDataAccessor {
	return &accountDataAccessor{
		database: database,
	}
}
