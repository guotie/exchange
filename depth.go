package exchange

import "github.com/shopspring/decimal"

// Depth 深度
type Depth struct {
	Symbol string      `json:"symbol"`
	Asks   []DepthItem `json:"asks"`
	Bids   []DepthItem `json:"bids"`
}

// DepthItem price, amount
type DepthItem [2]decimal.Decimal

// DepthEntry 有些交易所返回的depth 数组是一个类似于DepthEntry的结构体
type DepthEntry struct {
	Price  decimal.Decimal
	Amount decimal.Decimal
}
