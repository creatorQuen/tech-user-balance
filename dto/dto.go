package dto

import "github.com/shopspring/decimal"

type AddBalanceReq struct {
	UserId  string
	Balance decimal.Decimal
}
