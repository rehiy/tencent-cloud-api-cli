package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func sha256hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}

func hmacsha256(s, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	return string(hashed.Sum(nil))
}

var (
	h       bool
	service string
	version string
	action  string
	region  string
	payload string
)

func init() {

	flag.BoolVar(&h, "h", false, "查看帮助")

	flag.StringVar(&service, "service", "", "服务名")
	flag.StringVar(&version, "version", "", "服务版本")
	flag.StringVar(&action, "action", "", "服务接口名")
	flag.StringVar(&region, "region", "", "地域")
	flag.StringVar(&payload, "payload", "", "接口入参")

	// 改变默认的 Usage
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `注意：使用之前需检查系统环境变量TENCENTCLOUD_SECRET_ID 和 TENCENTCLOUD_SECRET_KEY是否设置
若果未设置，请执行命令设置：
export TENCENTCLOUD_SECRET_KEY=xxxx
export TENCENTCLOUD_SECRET_ID=yyyy

使用方法: tencentcloudapi_curl -service cvm -version 2017-03-12 -action "DescribeRegions" -region "ap-guangzhou" -payload "{}"

Options:
`)
	flag.PrintDefaults()
}

func HttpPost(url, msg string, headers map[string]string) (string, error) {

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

func main() {

	flag.Parse()

	if h {
		flag.Usage()
		return
	}

	if service == "" || version == "" || action == "" {
		flag.Usage()
		return
	}

	secretId, ok1 := os.LookupEnv("TENCENTCLOUD_SECRET_ID")
	secretKey, ok2 := os.LookupEnv("TENCENTCLOUD_SECRET_KEY")

	if !ok1 || !ok2 {
		fmt.Println("请执行export命令检查是否设置系统环境变量TENCENTCLOUD_SECRET_ID 和 TENCENTCLOUD_SECRET_KEY")
		return
	}

	host := service + ".tencentcloudapi.com"
	algorithm := "TC3-HMAC-SHA256"

	var timestamp int64 = time.Now().Unix()

	// step 1: build canonical request string

	httpRequestMethod := "POST"
	canonicalURI := "/"
	canonicalQueryString := ""
	canonicalHeaders := "content-type:application/json; charset=utf-8\n" + "host:" + host + "\n"
	signedHeaders := "content-type;host"
	hashedRequestPayload := sha256hex(payload)
	canonicalRequest := fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s\n%s",
		httpRequestMethod,
		canonicalURI,
		canonicalQueryString,
		canonicalHeaders,
		signedHeaders,
		hashedRequestPayload,
	)
	//fmt.Println(canonicalRequest)

	// step 2: build string to sign

	date := time.Unix(timestamp, 0).UTC().Format("2006-01-02")
	credentialScope := fmt.Sprintf("%s/%s/tc3_request", date, service)
	hashedCanonicalRequest := sha256hex(canonicalRequest)
	string2sign := fmt.Sprintf(
		"%s\n%d\n%s\n%s",
		algorithm,
		timestamp,
		credentialScope,
		hashedCanonicalRequest,
	)
	//fmt.Println(string2sign)

	// step 3: sign string

	secretDate := hmacsha256(date, "TC3"+secretKey)
	secretService := hmacsha256(service, secretDate)
	secretSigning := hmacsha256("tc3_request", secretService)
	signature := hex.EncodeToString([]byte(hmacsha256(string2sign, secretSigning)))
	//fmt.Println(signature)

	// step 4: build authorization

	authorization := fmt.Sprintf(
		"%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		algorithm,
		secretId,
		credentialScope,
		signedHeaders,
		signature,
	)
	//fmt.Println(authorization)

	// step 5: send https request

	url := "https://" + host
	headers := make(map[string]string)
	headers["Authorization"] = authorization
	headers["Content-Type"] = "application/json; charset=utf-8"
	headers["Host"] = host
	headers["X-TC-Action"] = action
	headers["X-TC-Timestamp"] = strconv.FormatInt(timestamp, 10)
	headers["X-TC-Version"] = version
	if region != "" {
		headers["X-TC-Region"] = region
	}

	res, err := HttpPost(url, payload, headers)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
