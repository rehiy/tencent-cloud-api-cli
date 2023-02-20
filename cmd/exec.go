package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/rehiy/tencent-cloud-api-cli/api"
)

var (
	service string
	version string
	action  string
	region  string
	payload string
)

func Exec() {

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

	res, err := api.Request(service, version, action, region, payload, secretId, secretKey)

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

tcapi --service cvm --version 2017-03-12 --action DescribeRegions --region ap-guangzhou --payload "{}"

选项说明:

`)
		flag.PrintDefaults()

	}

	flag.Parse()

}
