package controller

import (
	"gopkg.in/olivere/elastic.v5"
	"spiders/distributedCrawfer/frontend/view"
	"net/http"
	"strings"
	"strconv"
	"spiders/distributedCrawfer/frontend/model"
	"spiders/distributedCrawfer/config"
	"context"
	"reflect"
	"spiders/distributedCrawfer/engine"
	"regexp"
)

// TODO
// fill in query string
// support search button
// support paging
// add start page
type SearchresultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchresultHandler(template string) SearchresultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchresultHandler{
		view.CreateSearchResultView(template),
		client,
	}

}

// localhost:8888/search?q=男 已购房&from=20
func (h SearchresultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 取出参数,并去除空格
	q := strings.TrimSpace(req.FormValue("q"))
	//q = rewriteQueryString(q)
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}

	//var page model.SearchResult

	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}

func (h SearchresultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	result.Query = q
	resp, err := h.client.Search(config.IndexName).Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).From(from).Do(context.Background())
	if err != nil {
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))

	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}

func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")

}
