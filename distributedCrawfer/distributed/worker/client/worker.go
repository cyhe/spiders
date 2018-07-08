package client

import (
	"spiders/distributedCrawfer/engine"
	"spiders/distributedCrawfer/config"
	"spiders/distributedCrawfer/distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor{
	//client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkPort0))
	//if err != nil {
	//	return nil, err
	//}

	return func(req engine.Request) (engine.ParserResult, error) {

		sReq := worker.SerializeRequest(req)

		var sResult worker.ParseResult
		c := <-clientChan
		err := c.Call(config.CrawlServiceFunName, sReq, &sResult)

		if err != nil {
			return engine.ParserResult{}, err
		}

		return worker.DeserializeResult(sResult)
	}
}
