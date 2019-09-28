package huobi

const (
	AccountTypeSpot   = "spot"
	AccountTypeMargin = "margin"
	AccountTypeOTC    = "OTC"
	AccountTypePoint  = "point"
)

// AccountInfo account info
// 参数名称	是否必须	数据类型	描述	取值范围
// id	true	long	account-id
// state	true	string	账户状态	working：正常, lock：账户被锁定
// type	true	string	账户类型	spot：现货账户， margin：杠杆账户，otc：OTC 账户，point：点卡账户
//
type AccountInfo struct {
	ID    uint64 `json:"id"`
	State string `json:"state"`
	Type  string `json:"type"`
}

// Balance account balance
// 参数名称	是否必须	数据类型	描述	取值范围
// id	true	long	账户 ID
// state	true	string	账户状态	working：正常 lock：账户被锁定
// type	true	string	账户类型	spot：现货账户， margin：杠杆账户，otc：OTC 账户，point：点卡账户
// list	false	Array	子账号数组
// list字段说明
type Balance struct {
	ID    uint64        `json:"id"`
	State string        `json:"state"`
	Type  string        `json:"type"`
	Items []BalanceItem `json:"list"`
}

const (
	// BlanceItemTypeTrade trade: 交易余额
	BlanceItemTypeTrade = "trade"
	// BlanceItemTypeFrozen frozen: 冻结余额
	BlanceItemTypeFrozen = "frozen"
)

// BalanceItem balance item
// 参数名称	是否必须	数据类型	描述	取值范围
// balance	true	string	余额
// currency	true	string	币种
// type	true	string	类型	trade: 交易余额，frozen: 冻结余额
type BalanceItem struct {
	Balance  string `json:"balance"`
	Currency string `json:"currency"`
	Type     string `json:"type"`
}

const (
	// OrderPriceTypeBuyMarket "buy-market"
	OrderPriceTypeBuyMarket = "buy-market"
	// OrderPriceTypeSellMarket "sell-market"
	OrderPriceTypeSellMarket = "sell-market"
	// OrderPriceTypeBuyLimit "buy-limit"
	OrderPriceTypeBuyLimit = "buy-limit"
	// OrderPriceTypeSellLimit "sell-limit"
	OrderPriceTypeSellLimit = "sell-limit"
	// OrderPriceTypeBuyIOC buy-ioc
	OrderPriceTypeBuyIOC = "buy-ioc"
	// OrderPriceTypeSellIOC sell-ioc
	OrderPriceTypeSellIOC = "sell-ioc"
	// OrderPriceTypeBuyLimitMaker buy-limit-maker
	OrderPriceTypeBuyLimitMaker = "buy-limit-maker"
	// OrderPriceTypeSellLimitMaker sell-limit-maker
	OrderPriceTypeSellLimitMaker = "sell-limit-maker"
	// OrderPriceTypeBuyStopLimit buy-stop-limit
	OrderPriceTypeBuyStopLimit = "buy-stop-limit"
	// OrderPriceTypeSellStopLimit sell-stop-limit
	OrderPriceTypeSellStopLimit = "sell-stop-limit"
)

// OrderPlaceReq 下单请求
// 参数名称	数据类型	是否必需	默认值	描述
// account-id	string	true	NA	账户 ID，使用 GET /v1/account/accounts 接口查询。现货交易使用 ‘spot’ 账户的 account-id；杠杆交易，请使用 ‘margin’ 账户的 account-id
// symbol	string	true	NA	交易对, 例如btcusdt, ethbtc
// type	string	true	NA	订单类型，包括buy-market, sell-market, buy-limit, sell-limit, buy-ioc, sell-ioc, buy-limit-maker, sell-limit-maker, buy-stop-limit, sell-stop-limit（说明见下文）
// amount	string	true	NA	订单交易量
// price	string	false	NA	limit order的交易价格
// source	string	false	api	现货交易填写“api”，杠杆交易填写“margin-api”
// client-order-id	string	false	NA	用户自编订单号（须在24小时内保持唯一性）
// stop-price	string	false	NA	止盈止损订单触发价格
// operator	string	false	NA	止盈止损订单触发价运算符 gte – greater than and equal (>=), lte – less than and equal (<=)
type OrderPlaceReq struct {
	AccountID     string `json:"account-id"`
	Symbol        string `json:"symbol"`
	Type          string `json:"type"`
	Amount        string `json:"amount"`
	Price         string `json:"price"`
	Source        string `json:"source"`
	ClientOrderID string `json:"client-order-id"`
	StopPrice     string `json:"stop-price"`
	Operator      string `json:"operator"`
}
