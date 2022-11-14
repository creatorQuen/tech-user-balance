package domain

import "github.com/shopspring/decimal"

type UserBalance struct {
	Id      string
	UserId  string
	Balance decimal.Decimal
}
