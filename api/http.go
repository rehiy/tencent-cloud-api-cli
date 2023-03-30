package api

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Request(service, version, action, region, payload, secretId, secretKey string) (string, error) {

	timestamp := time.Now().Unix()
	host := service + ".tencentcloudapi.com"

	authorization := AuthCode(
		service,
		payload,
		secretId,
		secretKey,
		timestamp,
	)

	// send https request

	headers := map[string]string{
		"Authorization":  authorization,
		"Content-Type":   "application/json; charset=utf-8",
		"Host":           host,
		"X-TC-Action":    action,
		"X-TC-Timestamp": strconv.FormatInt(timestamp, 10),
		"X-TC-Version":   version,
	}

	if region != "" {
		headers["X-TC-Region"] = region
	}

	return httpPost("https://"+host, payload, headers)

}

func httpPost(url, msg string, headers map[string]string) (string, error) {

	var err error
	var req *http.Request
	var resp *http.Response
	var body []byte

	rd := strings.NewReader(msg)
	if req, err = http.NewRequest("POST", url, rd); err != nil {
		return "", err
	}

	for key, header := range headers {
		req.Header.Set(key, header)
	}

	client := &http.Client{}
	if resp, err = client.Do(req); err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return "", err
	}

	return string(body), err

}
