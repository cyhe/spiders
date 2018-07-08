package client

import (
	"log"
	"spiders/distributedCrawfer/engine"
	"spiders/distributedCrawfer/distributed/rpcsupport"
	"fmt"
	"spiders/distributedCrawfer/config"
)

func ItemSaver(host string) (chan engine.Item, error) {

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Client ItemSaver Got item #%d: %v", itemCount, item)

			// call rpc to save item
			result := ""
			fmt.Println("++++++++++")
			err = client.Call(config.ItemSaverRpcFunName, item, &result)
			if err != nil {
				log.Printf("Item Saver: error"+"saving item %v: %v", item, err)
			}

			fmt.Println("----------")

			itemCount++

		}
	}()
	return out, nil
}
