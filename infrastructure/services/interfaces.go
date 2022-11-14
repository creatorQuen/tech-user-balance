package services

import (
	"context"
	"tech-user-balance/dto"
)

type UserBalanceService interface {
	AddBalance(ctx context.Context, req dto.AddBalanceReq) (string, error)
}
