package persist

import (
	"log"
	"context"
	"spiders/concurrencyCrawfer/engine"
	"github.com/pkg/errors"
	"gopkg.in/olivere/elastic.v5"
)

func ItemSaver(index string) (chan engine.Item, error) {

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("ItemSaver Got item #%d: %v", itemCount, item)
			itemCount++

			err := save(index, client, item)
			if err != nil {
				log.Printf("Item Saver: error"+"saving item %v: %v", item, err)
			}

		}
	}()
	return out, nil
}

func save(index string, client *elastic.Client, item engine.Item) (err error) {
	// must turn off sniff in docker
	//client, err := elastic.NewClient(elastic.SetSniff(false))
	//if err != nil {
	//	return err
	//}

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().Index(index).Type(item.Type).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err = indexService.Do(context.Background())

	if err != nil {
		return err
	}
	return nil
}
