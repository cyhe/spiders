package engine

import (
	"log"
	"spiders/distributedCrawfer/fetcher"
)

func Worker(r Request) (ParserResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)

	if err != nil {
		log.Printf("Fetcher: error"+"fetching url %s: %v", r.Url, err)
		return ParserResult{}, err
	}
	//return r.ParserFunc(body, r.Url), nil
	return r.Parser.Parse(body, r.Url), nil
}
