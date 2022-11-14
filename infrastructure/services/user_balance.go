package services

import (
	"context"
	"tech-user-balance/domain"
	"tech-user-balance/dto"
	"tech-user-balance/infrastructure/repository"
	"tech-user-balance/pkg"
)

type userBalanceService struct {
	repoUserBalance repository.UserBalanceRepo
	log             *pkg.Logger
}

func NewUserBalanceService(repoUserBalance repository.UserBalanceRepo, log *pkg.Logger) *userBalanceService {
	return &userBalanceService{
		repoUserBalance: repoUserBalance,
		log:             log,
	}
}

func (u *userBalanceService) AddBalance(ctx context.Context, req dto.AddBalanceReq) (string, error) {
	userBalance := domain.UserBalance{Balance: req.Balance, UserId: req.UserId}
	return u.repoUserBalance.AddBalance(ctx, userBalance)
}
