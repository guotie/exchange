package exchange

import "github.com/shopspring/decimal"

// KLine k-line
type KLine struct {
	Timestamp uint64
	Open      decimal.Decimal
	High      decimal.Decimal
	Low       decimal.Decimal
	Close     decimal.Decimal
}
