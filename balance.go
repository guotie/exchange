package exchange

import "github.com/shopspring/decimal"

// BalanceItem balance item
type BalanceItem struct {
	Avail     decimal.Decimal
	Frozen    decimal.Decimal
	OTCFrozen decimal.Decimal
}

// AccountBalance account balance
type AccountBalance struct {
	AccountTyp int
	Balances   map[Coin]BalanceItem
}
