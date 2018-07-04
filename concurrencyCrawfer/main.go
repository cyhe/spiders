package main

import (
	"spiders/concurrencyCrawfer/engine"
	"spiders/concurrencyCrawfer/zhenai/parser"
	"spiders/concurrencyCrawfer/scheduler"
	"spiders/concurrencyCrawfer/persist"
	"spiders/concurrencyCrawfer/config"
)

func main() {

	itemChan, err := persist.ItemSaver(config.IndexName)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}
