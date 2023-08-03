package expediago

import (
	"bytes"
	"compress/gzip"
	"crypto/sha512"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func ApiRequest(endpoint string, method string, data interface{}) ([]byte, int) {

	url := GetApiBaseUrl() + endpoint

	client := &http.Client{}
	var err error
	var req *http.Request
	var reader io.ReadCloser

	if method == "POST" {
		req, err = http.NewRequest(method, url, bytes.NewBuffer([]byte(data.(string))))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		fmt.Println(err.Error())
		return ApiRequest(url, method, data)
	}
	req.Header.Add("Authorization", generateAuthorizationHeader())
	req.Header.Add("Customer-Ip", "127.0.0.1")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Encoding", "gzip")
	//req.Header.Add("Test", "sold_out")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return ApiRequest(url, method, data)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(res.Body)

	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(res.Body)
		defer func(reader io.ReadCloser) {
			err := reader.Close()
			if err != nil {
				fmt.Println(err.Error())
			}
		}(reader)
	default:
		reader = res.Body
	}
	body, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println(err.Error())
		return ApiRequest(url, method, data)
	}

	return body, res.StatusCode
}

func generateAuthorizationHeader() string {
	timestamp := int(time.Now().Unix())
	apiKey := GetApiKey()
	sharedSecret := GetApiSecret()
	sha512v := sha512.New()
	sha512v.Write([]byte(apiKey + sharedSecret + strconv.Itoa(timestamp)))
	hash := fmt.Sprintf("%x", sha512v.Sum(nil))
	auth := "EAN APIKey=" + apiKey + ",Signature=" + hash + ",timestamp=" + strconv.Itoa(timestamp)
	return auth
}
