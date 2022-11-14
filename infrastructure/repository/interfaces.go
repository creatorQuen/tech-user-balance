package repository

import (
	"context"
	"tech-user-balance/domain"
)

type UserBalanceRepo interface {
	AddBalance(ctx context.Context, userBalance domain.UserBalance) (string, error)
}
