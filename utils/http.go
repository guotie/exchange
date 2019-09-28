package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// BuildURI build uri
func BuildURI(host, path string, param *url.Values) string {
	var res = host

	if strings.HasSuffix(res, "/") {
		if strings.HasPrefix(path, "/") {
			res += path[1:]
		} else {
			res += path
		}
	} else {
		if strings.HasPrefix(path, "/") {
			res += path
		} else {
			res += "/" + path
		}
	}
	if param != nil {
		if strings.Contains(path, "?") {
			res += "&"
		} else {
			res += "?"
		}
		res += param.Encode()
	}
	return res
}

// NewHTTPRequest create http request
func NewHTTPRequest(client *http.Client, reqType string, reqURL string, postData string, requstHeaders map[string]string) ([]byte, error) {
	req, _ := http.NewRequest(reqType, reqURL, strings.NewReader(postData))
	if req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36")
	}
	if requstHeaders != nil {
		for k, v := range requstHeaders {
			req.Header.Add(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HttpStatusCode: %d ,Desc: %s", resp.StatusCode, string(bodyData))
	}

	return bodyData, nil
}

// HTTPGet get method
func HTTPGet(client *http.Client, reqURL string) (map[string]interface{}, error) {
	respData, err := NewHTTPRequest(client, "GET", reqURL, "", nil)
	if err != nil {
		return nil, err
	}

	var bodyDataMap map[string]interface{}
	err = json.Unmarshal(respData, &bodyDataMap)
	if err != nil {
		log.Println(string(respData))
		return nil, err
	}
	return bodyDataMap, nil
}

// HTTPGetWithValues get method
func HTTPGetWithValues(client *http.Client, reqURL string) (map[string]interface{}, error) {
	respData, err := NewHTTPRequest(client, "GET", reqURL, "", nil)
	if err != nil {
		return nil, err
	}

	var bodyDataMap map[string]interface{}
	err = json.Unmarshal(respData, &bodyDataMap)
	if err != nil {
		log.Println(string(respData))
		return nil, err
	}
	return bodyDataMap, nil
}

// HTTPPost post with application/json
func HTTPPost(client *http.Client, reqURL string, postData string, headers map[string]string) ([]byte, error) {
	return NewHTTPRequest(client, "POST", reqURL, postData, headers)
}

// HTTPPostJSON http post with json
func HTTPPostJSON(client *http.Client, reqURL string, postData interface{}, headers map[string]string) ([]byte, error) {
	if headers == nil {
		headers = map[string]string{}
	}

	headers["Content-Type"] = "application/json"
	data, _ := json.Marshal(postData)
	return NewHTTPRequest(client, "POST", reqURL, string(data), headers)
}
