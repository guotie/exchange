package exchange

const (
	// AccountTypeSpot 现货
	AccountTypeSpot = iota + 1
	// AccountTypeLeverage 期货账户
	AccountTypeLeverage
	// AccountTypeContract 合约账号
	AccountTypeContract
	// AccountTypeSwap 永续合约
	AccountTypeSwap
)

// Account exchange account
type Account struct {
	ID string
	AccessKey   string
	SecretKey   string
	AccountType int
	KycLevel    int
}
