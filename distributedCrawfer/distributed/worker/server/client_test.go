package main

import (
	"testing"
	"spiders/distributedCrawfer/distributed/rpcsupport"
	"spiders/distributedCrawfer/distributed/worker"
	"time"
	"spiders/distributedCrawfer/config"
	"fmt"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/1558719774",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "可惜我爱怀念",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceFunName, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
