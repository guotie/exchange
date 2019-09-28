package huobi

import (
	"encoding/json"
	"exchange/utils"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"
)

func convertToValues(reqBody interface{}) (*url.Values, error) {
	return utils.MarshalToValues(reqBody)
}

func convertValuesToJSON(param *url.Values) map[string]string {
	res := map[string]string{}
	for k, v := range *param {
		if len(v) == 1 {
			res[k] = v[0]
		} else if len(v) == 0 {
			log.Printf("url.Values key %s value has zero value", k)
		} else {
			log.Printf("url.Values key %s value has mulit value: %v", k, v)
		}
	}

	return res
}

func (hb *ExHuobi) buildSignParams(method, path string, params *url.Values) {
	params.Set("AccessKeyId", hb.accessKey)
	params.Set("SignatureMethod", "HmacSHA256")
	params.Set("SignatureVersion", "2")
	params.Set("Timestamp", time.Now().UTC().Format("2006-01-02T15:04:05"))
	domain := strings.Replace(hb.baseURL, "https://", "", len(hb.baseURL))
	payload := fmt.Sprintf("%s\n%s\n%s\n%s", method, domain, path, params.Encode())
	sign, _ := utils.HMACSHA256Base64Sign(hb.secretKey, payload)
	params.Set("Signature", sign)
}

func (hb *ExHuobi) buildURL(path string, params *url.Values) string {
	return utils.BuildURI(hb.baseURL, path, params)
}

func (hb *ExHuobi) httpPost(path string, reqBody interface{}, ret interface{}) (err error) {
	var param *url.Values

	if reqBody == nil {
		param = &url.Values{}
	} else {
		param, err = convertToValues(reqBody)
		if err != nil {
			return err
		}
	}
	hb.buildSignParams("POST", path, param)

	resp, err := utils.HTTPPostJSON(hb.client, hb.buildURL(path, nil), param, nil)

	err = json.Unmarshal(resp, ret)

	return err
}

func (hb *ExHuobi) postRequestWithSign() {

}

func (hb *ExHuobi) getRequestWithSign() {

}
