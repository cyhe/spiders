package persist

import (
	"testing"
	"spiders/distributedCrawfer/model"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"encoding/json"
	"spiders/distributedCrawfer/config"
	"spiders/distributedCrawfer/engine"
)

func TestItemSaver(t *testing.T) {
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1558719774",
		Type: config.TypeName,
		Id:   "1558719774",
		Payload: model.Profile{
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

	// TODO : Try to start up elastic search.  here using docker go client
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	// Save expected item
	err = Save(index ,client, expected)
	if err != nil {
		panic(err)
	}

	// fetch saved item
	resp, err := client.Get().Index(index).Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)

	var actual engine.Item

	json.Unmarshal(*resp.Source, &actual)

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	// verify saved item
	if actual != expected {
		t.Errorf("got item : %v,expected : %v", actual, expected)
	}
}
