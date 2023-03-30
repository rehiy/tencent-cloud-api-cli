package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	service   string
	version   string
	action    string
	region    string
	payload   string
	secretId  string
	secretKey string
)

func parseFlag() {

	flag.StringVar(&service, "service", "", "服务名")
	flag.StringVar(&version, "version", "", "服务版本")
	flag.StringVar(&action, "action", "", "执行动作")
	flag.StringVar(&region, "region", "", "地域")
	flag.StringVar(&payload, "payload", "", "结构化数据")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, readme)
		flag.PrintDefaults()
	}

	flag.Parse()

}

func checkSecret() {

	secretId, _ = os.LookupEnv("TENCENTCLOUD_SECRET_ID")
	secretKey, _ = os.LookupEnv("TENCENTCLOUD_SECRET_KEY")

	if secretId == "" || secretKey == "" {
		log.Fatal("请设置环境变量 TENCENTCLOUD_SECRET_ID 和 TENCENTCLOUD_SECRET_KEY")
	}

}

const readme = `
使用方法:

tcapi --service cvm --version 2017-03-12 --action DescribeRegions --region ap-guangzhou --payload "{}"

选项说明:

`
