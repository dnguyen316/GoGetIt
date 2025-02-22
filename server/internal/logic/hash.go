package logic

import (
	"context"
	"errors"

	"github.com/dnguyen316/GoGetIt/server/internal/config"
	"golang.org/x/crypto/bcrypt"
)

type Hash interface {
	Hash(ctx context.Context, data string) (string, error)
	IsHashEqual(ctx context.Context, data string, hashed string) (bool, error)
}

type hash struct {
	accountConfig config.Account
}

func NewHash(accountConfig config.Account) Hash {
	return &hash{
		accountConfig: accountConfig,
	}
}

// Hash implements Hash.
func (h *hash) Hash(ctx context.Context, data string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(data), h.accountConfig.HashCost)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

// IsHashEqual implements Hash.
func (h *hash) IsHashEqual(ctx context.Context, data string, hashed string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(data)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
