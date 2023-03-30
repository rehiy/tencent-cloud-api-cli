package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/rehiy/tencent-cloud-api-cli/api"
)

func Exec() {

	usage()

	if service == "" {
		log.Fatal("请设置 -service 参数，-h 查看帮助")
		return
	}

	if version == "" {
		log.Fatal("请设置 -version 参数，-h 查看帮助")
		return
	}

	if action == "" {
		log.Fatal("请设置 -action 参数，-h 查看帮助")
		return
	}

	secretId, ok1 := os.LookupEnv("TENCENTCLOUD_SECRET_ID")
	secretKey, ok2 := os.LookupEnv("TENCENTCLOUD_SECRET_KEY")

	if !ok1 || !ok2 {
		log.Fatal("请设置环境变量 TENCENTCLOUD_SECRET_ID 和 TENCENTCLOUD_SECRET_KEY")
	}

	res, err := api.Request(service, version, action, region, payload, secretId, secretKey)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

}
