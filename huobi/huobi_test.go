package huobi

import (
	"net"
	"net/http"
	"net/url"
	"testing"
	"time"
)

var (
	accessKey = "edrfhh5h53-1d7638ce-da3e4fec-0f1f0"
	secretKey = "5a65ad00-d029223e-de893d0a-e4681"
	accountID = "101754951"
)

var httpProxyClient = &http.Client{
	Transport: &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return &url.URL{
				Scheme: "http",
				Host:   "127.0.0.1:25378"}, nil
		},
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
	},
	Timeout: 10 * time.Second,
}

func TestAccountInfo(t *testing.T) {
	hb := NewHuoBiPro(httpProxyClient, accessKey, secretKey, accountID)

	info, err := hb.GetAccountInfo(AccountTypeSpot)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(info)
}

func TestBalance(t *testing.T) {

}

func TestUSDT(t *testing.T) {

}
