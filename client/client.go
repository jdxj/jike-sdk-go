package client

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func MockRegister() {
	URL := "https://api.ruguoapp.com/1.0/users/register"
	body := strings.NewReader(`{"username":"71929ad7-186d-4baa-864a-c2067656278c","password":"6ESo0rbzGGcghcOx"}`)
	req, err := http.NewRequest(http.MethodPost, URL, body)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Host", "api.ruguoapp.com")
	req.Header.Set("x-jike-app-id", "XeITUMa6kGKF")
	req.Header.Set("os-version", "29")
	req.Header.Set("os", "Android")
	req.Header.Set("app-buildno", "1071")
	req.Header.Set("manufacturer", "Xiaomi")
	req.Header.Set("model", "Redmi K20 Pro")
	req.Header.Set("app-version", "7.2.1")
	req.Header.Set("x-jike-device-id", "426883b4-b8e0-49bb-b105-d79b298e8c7f")
	req.Header.Set("applicationid", "com.ruguoapp.jike")
	req.Header.Set("market", "yingyongbao")
	req.Header.Set("resolution", "1080x2208")
	req.Header.Set("source", "")
	req.Header.Set("app-permissions", "4")
	req.Header.Set("x-jike-device-properties", "{\"uuid\":\"426883b4-b8e0-49bb-b105-d79b298e8c7f\",\"android_id\":\"42b3efa6d8113f7d\",\"oaid\":\"\",\"vaid\":\"\",\"aaid\":\"\"}")
	req.Header.Set("content-type", "application/json;charset=utf-8")
	req.Header.Set("user-agent", "okhttp/4.7.2")
	req.Header.Set("accept-encoding", "gzip")

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	r, err := gzip.NewReader(resp.Body)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)

	accessToken := resp.Header.Get("x-jike-access-token")
	refreshToken := resp.Header.Get("x-jike-refresh-token")
	fmt.Printf("at: %s\n", accessToken)
	fmt.Printf("rt: %s\n", refreshToken)

	URL = "https://api.ruguoapp.com/1.0/users/loginWithPhoneAndPassword"
	body = strings.NewReader(`{"password":"","mobilePhoneNumber":"","areaCode":"+86"}`)
	req, err = http.NewRequest(http.MethodPost, URL, body)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Host", "api.ruguoapp.com")
	req.Header.Set("x-jike-app-id", "XeITUMa6kGKF")
	req.Header.Set("os-version", "29")
	req.Header.Set("os", "Android")
	req.Header.Set("app-buildno", "1071")
	req.Header.Set("manufacturer", "Xiaomi")
	req.Header.Set("model", "Redmi K20 Pro")
	req.Header.Set("app-version", "7.2.1")
	req.Header.Set("x-jike-device-id", "426883b4-b8e0-49bb-b105-d79b298e8c7f")
	req.Header.Set("applicationid", "com.ruguoapp.jike")
	req.Header.Set("market", "yingyongbao")
	req.Header.Set("resolution", "1080x2208")
	req.Header.Set("source", "")
	req.Header.Set("app-permissions", "4")
	req.Header.Set("x-jike-device-id", accessToken)
	req.Header.Set("x-jike-refresh-token", refreshToken)
	req.Header.Set("x-jike-device-properties", "{\"uuid\":\"426883b4-b8e0-49bb-b105-d79b298e8c7f\",\"android_id\":\"42b3efa6d8113f7d\",\"oaid\":\"\",\"vaid\":\"\",\"aaid\":\"\"}")
	req.Header.Set("content-type", "application/json;charset=utf-8")
	req.Header.Set("accept-encoding", "gzip")
	req.Header.Set("user-agent", "okhttp/4.7.2")

	resp2, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp2.Body.Close()

	r2, err := gzip.NewReader(resp2.Body)
	if err != nil {
		panic(err)
	}

	data, err = ioutil.ReadAll(r2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)
}
