package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/rehiy/tencent-cloud-api-cli/api"
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
