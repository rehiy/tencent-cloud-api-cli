package cmd

import (
	"fmt"
	"log"

	"github.com/rehiy/tencent-cloud-api-cli/api"
)

func Exec() {

	parseFlag()
	checkSecret()

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

	res, err := api.Request(service, version, action, region, payload, secretId, secretKey)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

}
