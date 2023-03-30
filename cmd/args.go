package cmd

import (
	"flag"
	"fmt"
	"os"
)

var (
	service string
	version string
	action  string
	region  string
	payload string
)

func usage() {

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

const readme = `
使用方法:

tcapi --service cvm --version 2017-03-12 --action DescribeRegions --region ap-guangzhou --payload "{}"

选项说明:

`
