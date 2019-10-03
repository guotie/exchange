package wtzex

import (
	"github.com/shopspring/decimal"
)

// PrecisionInfo precision info
type PrecisionInfo struct {
	Amount int `json:"amount"`
	Price  int `json:"price"`
}

// MaxMin max min
type MaxMin struct {
	Max decimal.Decimal `json:"max"`
	Min decimal.Decimal `json:"min"`
}

// LimitsInfo amount & price limits
type LimitsInfo struct {
	Amount MaxMin `json:"amount"`
	Price  MaxMin `json:"price"`
}

// MarketInfo market
type MarketInfo struct {
	Active     bool            `json:"active"`
	Precision  PrecisionInfo   `json:"precision"`
	Percentage bool            `json:"percentage"`
	Taker      decimal.Decimal `json:"taker"`
	Maker      decimal.Decimal `json:"maker"`
	ID         string          `json:"id"`
	Symbol     string          `json:"symbol"`
	Base       string          `json:"base"`
	Quote      string          `json:"quote"`
	BaseID     string          `json:"baseId"`
	QuoteID    string          `json:"quoteId"`
}

// Ticker ticker
type Ticker struct {
	Symbol     string          `json:"s"`
	Timestamp  long            `json:"t"`
	Open       decimal.Decimal `json:"o"`
	High       decimal.Decimal `json:"h"`
	Low        decimal.Decimal `json:"l"`
	Close      decimal.Decimal `json:"c"`
	Percentage decimal.Decimal `json:"p"`
	Changed    decimal.Decimal `json:"cg"`
	Volume     decimal.Decimal `json:"v"`
	Turnover   decimal.Decimal `json:"to"`
	Ask        decimal.Decimal `json:"a"`
	Bid        decimal.Decimal `json:"b"`
	AskAmount  decimal.Decimal `json:"aa"`
	BidAmount  decimal.Decimal `json:"ba"`
}

// Depth depth
type Depth struct {
	Asks [][]decimail.Decimal `json:"asks"`
	Bids [][]decimail.Decimal `json:"bids"`
}
