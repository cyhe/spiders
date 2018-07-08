package main

import (
	"spiders/distributedCrawfer/engine"
	"spiders/distributedCrawfer/scheduler"
	"spiders/distributedCrawfer/persist"
	"spiders/distributedCrawfer/config"
	"spiders/distributedCrawfer/zhenai/parser"
)

func main() {

	itemChan, err := persist.ItemSaver(config.IndexName)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		//ParserFunc: parser.ParseCityList,
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})

	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}
