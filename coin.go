package exchange

import "github.com/shopspring/decimal"

// Coin 加密货币
type Coin struct {
	Name            string
	Desc            string
	Precision       int
	CanWithdraw     bool
	CanDeposit      bool
	MaxWithdraw     decimal.Decimal
	MinWithdraw     decimal.Decimal
	WithdrawFee     decimal.Decimal
	ConfirmWithdraw int
	ConfirmDeposit  int
	Stable          bool
	BlockChain      string // 可用于区分不同链的 usdt
}
