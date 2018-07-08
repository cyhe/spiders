package main

import (
	"spiders/distributedCrawfer/engine"
	"spiders/distributedCrawfer/scheduler"
	itemSaver "spiders/distributedCrawfer/distributed/persist/client"
	"spiders/distributedCrawfer/config"
	"spiders/distributedCrawfer/zhenai/parser"
	worker "spiders/distributedCrawfer/distributed/worker/client"
	"net/rpc"
	"spiders/distributedCrawfer/distributed/rpcsupport"
	"log"
	"flag"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host","","itemsaver host")
	wokerhosts = flag.String("worker_hosts","","worker hosts (comma seprated)")
)


func main() {

	flag.Parse()

	itemChan, err := itemSaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*wokerhosts,","))

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("success connecting to %s", h)
		} else {
			log.Printf("error connecting to %s : %v", h, err)
		}
	}
	out := make(chan *rpc.Client)
	// 轮流分发
	go func() {
		// 始终轮流分发,多套一层for
		for {
			for _, client := range clients {
				out <- client
			}
		}

	}()
	return out
}
