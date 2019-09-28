package exchange

import (
	"github.com/shopspring/decimal"
)

// Symbol 交易对
type Symbol struct {
	ID              string
	Name            string
	Base            Coin
	Quote           Coin
	PricePrecision  int
	AmountPrecision int
	Enabled         bool
	Desc            string
	MinTurnOver     decimal.Decimal
	MinAmount       decimal.Decimal
	FeeTaker        decimal.Decimal
	FeeMaker        decimal.Decimal
}
