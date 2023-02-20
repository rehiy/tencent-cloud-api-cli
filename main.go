package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"tcli/helper"
	"time"
)

var (
	service string
	version string
	action  string
	region  string
	payload string
)

func main() {

	usage()

	if service == "" || version == "" || action == "" {
		flag.Usage()
		return
	}

	secretId, ok1 := os.LookupEnv("TENCENTCLOUD_SECRET_ID")
	secretKey, ok2 := os.LookupEnv("TENCENTCLOUD_SECRET_KEY")

	if !ok1 || !ok2 {
		fmt.Println("请设置环境变量 TENCENTCLOUD_SECRET_ID 和 TENCENTCLOUD_SECRET_KEY")
		return
	}

	timestamp := time.Now().Unix()
	host := service + ".tencentcloudapi.com"

	authorization := helper.AuthCode(
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

	res, err := helper.HttpPost("https://"+host, payload, headers)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}

func usage() {

	flag.StringVar(&service, "service", "", "服务名")
	flag.StringVar(&version, "version", "", "服务版本")
	flag.StringVar(&action, "action", "", "执行动作")
	flag.StringVar(&region, "region", "", "地域")
	flag.StringVar(&payload, "payload", "", "结构化数据")

	flag.Usage = func() {

		fmt.Fprintf(os.Stderr, `使用方法:

export TENCENTCLOUD_SECRET_ID=xxxx
export TENCENTCLOUD_SECRET_KEY=yyyy

tcli --service cvm --version 2017-03-12 --action DescribeRegions --region ap-guangzhou --payload "{}"

选项说明:

`)
		flag.PrintDefaults()

	}

	flag.Parse()

}
