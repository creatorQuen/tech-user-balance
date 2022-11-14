package repository

import (
	"context"
	"database/sql"
	"tech-user-balance/domain"
	"tech-user-balance/pkg"
)

type userBalanceRepositoryDB struct {
	db  *sql.DB
	log *pkg.Logger
}

func NewUserBalanceRepositoryDB(dbClinet *sql.DB, log *pkg.Logger) *userBalanceRepositoryDB {
	return &userBalanceRepositoryDB{
		db:  dbClinet,
		log: log,
	}
}

func (u *userBalanceRepositoryDB) AddBalance(ctx context.Context, userBalance domain.UserBalance) (string, error) {

}
