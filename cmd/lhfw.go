package cmd

import (
	"log"
	"strconv"

	js "github.com/bitly/go-simplejson"
	"github.com/rehiy/tencent-cloud-api-cli/api"
)

func GetLighthouseIds() {

	parseFlag()
	checkSecret()

	if region == "" {
		log.Fatal("请设置 -region 参数，-h 查看帮助")
		return
	}

	ready := 1

	for i := 0; ; i++ {

		log.Println("正在获取第", i, "页实例信息")

		payload := `{"Offset": ` + strconv.Itoa(i) + `, "Limit": 100}`

		res, err := api.Request("lighthouse", "2020-03-24", "DescribeInstances", region, payload, secretId, secretKey)

		if err != nil {
			log.Println(err)
			continue
		}

		obj, err := js.NewJson([]byte(res))

		if err != nil {
			log.Println(err)
			continue
		}

		total := obj.GetPath("Response", "TotalCount").MustInt()
		instanceSet := obj.GetPath("Response", "InstanceSet").MustArray()

		if len(instanceSet) == 0 {
			log.Println("未找到实例")
			break
		}

		for _, item := range instanceSet {
			id := item.(map[string]any)["InstanceId"].(string)
			log.Println("正在设置实例 ", id, " 防火墙规则（", ready, "/", total, "）")
			setFirewallRule(id)
			ready++
		}
	}

}

func setFirewallRule(id string) {

	payload := `{
		"InstanceId": "` + id + `",
		"FirewallRules": ` + payload + `
	}`

	res, err := api.Request("lighthouse", "2020-03-24", "CreateFirewallRules", region, payload, secretId, secretKey)

	if err != nil {
		log.Println(err)
	} else {
		log.Println(res)
	}

}
