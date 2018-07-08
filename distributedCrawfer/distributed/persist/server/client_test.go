package main

import (
	"testing"
	"spiders/distributedCrawfer/distributed/rpcsupport"
	"spiders/distributedCrawfer/engine"
	"spiders/distributedCrawfer/config"
	"spiders/distributedCrawfer/model"
	"time"
)

func TestIteamSaver(t *testing.T) {
	const host = ":1234"

	go serveRpc(host, "test1")

	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1558719774",
		Type: config.TypeName,
		Id:   "1558719774",
		Payload: model.Profile{
			Age:        35,
			Gender:     "女",
			Height:     168,
			Weight:     60,
			Income:     "20001-50000元",
			Name:       "可惜我爱怀念",
			Marriage:   "离异",
			Education:  "硕士",
			Occupation: "销售经理",
			Hokou:      "江苏连云港",
			Xinzuo:     "处女座",
			House:      "已购房",
			Car:        "已购车",
		},
	}
	result := ""
	err = client.Call("ItemSaverService.Save", item, &result)
	if err != nil || result != "OK" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
