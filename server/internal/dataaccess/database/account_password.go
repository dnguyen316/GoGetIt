package database

import (
	"context"

	"github.com/doug-martin/goqu/v9"
)

type AccountPassword struct {
	OfUserId uint64 `sql:"of_user_id"`
	Hash     string `sql:"hash"`
}

type AccountPasswordDataAccessor interface {
	CreateUserPassword(ctx context.Context, accountPassword AccountPassword) error
	UpdateUserPassword(ctx context.Context, accountPassword AccountPassword) error
}

type accountPasswordDataAccessor struct {
	database *goqu.Database
}

func NewAccountPasswordDataAccessor(
	database *goqu.Database,
) AccountPasswordDataAccessor {
	return &accountPasswordDataAccessor{
		database: database,
	}
}

// CreateUserPassword implements AccountPasswordDataAccessor.
func (a *accountPasswordDataAccessor) CreateUserPassword(ctx context.Context, accountPassword AccountPassword) error {
	panic("unimplemented")
}

// UpdateUserPassword implements AccountPasswordDataAccessor.
func (a *accountPasswordDataAccessor) UpdateUserPassword(ctx context.Context, accountPassword AccountPassword) error {
	panic("unimplemented")
}
