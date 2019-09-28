package huobi

import (
	"errors"
	"exchange"
	"exchange/utils"
	"fmt"
	"net/http"
	"net/url"
)

const (
	api   = ""
	wsAPI = ""
)

const (
	pathAccountInfo = "/v1/account/accounts"
	pathCancelOrder = "/v1/order/orders/%s/submitcancel"
	pathStablePrice = "/v1/stable-coin/quote"
)

// ExHuobi Exchange huobi
type ExHuobi struct {
	client    *http.Client
	baseURL   string
	accountID string
	accessKey string
	secretKey string
}

// NewDefaultHuoBiPro create client
func NewDefaultHuoBiPro(apikey, secretkey, accountID string) *ExHuobi {
	client := &http.Client{}

	return NewHuoBiPro(client, apikey, secretkey, accountID)
}

// NewHuoBiPro new huobi pro client
func NewHuoBiPro(client *http.Client, apikey, secretkey, accountID string) *ExHuobi {
	hbpro := &ExHuobi{
		client:    client,
		baseURL:   "https://api.huobi.br.com",
		accessKey: apikey,
		secretKey: secretkey,
		accountID: accountID,
	}

	return hbpro
}

// LimitBuy limit buy
func (hb *ExHuobi) LimitBuy(amount, price string, symbol string) (*exchange.Order, error) {
	return nil, nil
}

// CancelOrder cancel order
func (hb *ExHuobi) CancelOrder(orderID string, symbol string) (bool, error) {
	uri := fmt.Sprintf(hb.buildURL(pathCancelOrder, nil), orderID)
	if _, err := utils.HTTPPost(hb.client, uri, "{}", nil); err != nil {
		return false, err
	}
	return true, nil
}

// GetAccountInfo get account info
func (hb *ExHuobi) GetAccountInfo(acc string) (AccountInfo, error) {
	params := &url.Values{}
	hb.buildSignParams("GET", pathAccountInfo, params)

	//log.Println(hbpro.baseUrl + path + "?" + params.Encode())

	respmap, err := utils.HTTPGet(hb.client, hb.buildURL(pathAccountInfo, params))
	if err != nil {
		return AccountInfo{}, err
	}
	//log.Println(respmap)
	if respmap["status"].(string) != "ok" {
		return AccountInfo{}, errors.New(respmap["err-code"].(string))
	}

	var info AccountInfo

	data := respmap["data"].([]interface{})
	for _, v := range data {
		iddata := v.(map[string]interface{})
		if iddata["type"].(string) == acc {
			info.ID = utils.ConvertToUint64Must(iddata["id"])
			info.Type = acc
			info.State = iddata["state"].(string)
			break
		}
	}
	//log.Println(respmap)
	return info, nil
}

// GetStableCoinPrice 获取稳定币的价格
// func (hb *ExHuobi) GetStableCoinPrice(coin, side string, amount string) {
// 	params := &url.Values{}
// 	params.Add("currency", coin)
// 	params.Add("type", side)
// 	params.Add("amount", amount)
// 	hb.buildSignParams("GET", pathStablePrice, params)

// 	//log.Println(hbpro.baseUrl + path + "?" + params.Encode())

// 	respmap, err := utils.HTTPGet(hb.client, hb.buildURL(pathStablePrice, params))
// 	if err != nil {
// 		return AccountInfo{}, err
// 	}
// }
