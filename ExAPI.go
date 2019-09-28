package exchange

// ExAPI exchange api
type ExAPI interface {
	LimitBuy(amount, price string, symbol string) (*Order, error)
	LimitSell(amount, price string, symbol string) (*Order, error)
	MarketBuy(amount, price string, symbol string) (*Order, error)
	MarketSell(amount, price string, symbol string) (*Order, error)
	CancelOrder(orderID string, symbol string) (bool, error)
	FetchOrder(orderID string, symbol string) (*Order, error)
	FetchOrders(symbol string) ([]Order, error)
	FetchOrderHistorys(symbol string, currentPage, pageSize int) ([]Order, error)
	FetchAccount() (*Account, error)
	FetchPersonalTrades(currencyPair Symbol, since int64) ([]Trade, error)
	Balance()
	AccountInfo()
	Market()

	// public method
	GetTicker(symbol string) (*Ticker, error)
	GetDepth(size int, symbol string) (*Depth, error)
	GetKLine(symbol string, period, size, since int) ([]KLine, error)

	GetExchangeName() string
}
