package database

import (
	"context"
	"log"

	"github.com/doug-martin/goqu/v9"
)

type Account struct {
	UserId   uint64 `sql:"user_id"`
	Username string `sql:"username"`
}

type AccountDataAccessor interface {
	CreateAccount(ctx context.Context, account Account) (uint64, error)
	GetAccountByID(ctx context.Context, id uint64) (Account, error)
	GetAccountByUsername(ctx context.Context, username string) (Account, error)
}

type accountDataAccessor struct {
	database *goqu.Database
}

func (a accountDataAccessor) CreateAccount(ctx context.Context, account Account) (uint64, error) {
	// TODO: Implement create account logic
	result, err := a.database.Insert("accounts").Rows(goqu.Record{
		"username": account.Username,
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

func (a *accountDataAccessor) GetAccountByUsername(ctx context.Context, username string) (Account, error) {
	// TODO: Implement get account by username logic
	return Account{}, nil
}

func NewAccountDatabaseAccessor(
	database *goqu.Database,
) AccountDataAccessor {
	return &accountDataAccessor{
		database: database,
	}
}
