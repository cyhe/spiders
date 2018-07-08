package main

import (
	"spiders/distributedCrawfer/distributed/rpcsupport"
	"spiders/distributedCrawfer/distributed/persist"
	"gopkg.in/olivere/elastic.v5"
	"spiders/distributedCrawfer/config"
	"log"
	"fmt"
	"flag"
)



// 命令行参数
var port = flag.Int("port", 0, "the port for me to listen on ")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port ")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d",*port), config.IndexName))
}

func serveRpc(host string, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		client,
		index,
	})
}
