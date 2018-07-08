package view

import (
	"testing"
	"spiders/distributedCrawfer/frontend/model"
	"os"
	"spiders/distributedCrawfer/config"
	"spiders/distributedCrawfer/engine"
	commonModel "spiders/distributedCrawfer/model"
)

func TestSearchResultView_Render(t *testing.T) {
	//  template 一定要合法

	view := CreateSearchResultView("template.html")


	//tem := template.Must(template.ParseFiles("template.html"))

	out, err := os.Create("template_test.html")
	page := model.SearchResult{}
	page.Hits = 233
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1558719774",
		Type: config.TypeName,
		Id:   "1558719774",
		Payload: commonModel.Profile{
			Age:        35,
			Gender:     "女",
			Height:     168,
			Weight:     60,
			Income:     "20001-50000元",
			Name:       "可惜我爱怀念",
			Marriage:   "离异",
			Education:  "硕士",
			Occupation: "销售经理",
			Hokou:      "江苏连云港",
			Xinzuo:     "处女座",
			House:      "已购房",
			Car:        "已购车",
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out,page)
	if err != nil {
		panic(err)
	}
}
