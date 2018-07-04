package parser

import (
	"testing"
	"io/ioutil"
	"spiders/concurrencyCrawfer/model"
	"fmt"
	"spiders/concurrencyCrawfer/config"
	"spiders/concurrencyCrawfer/engine"
)

func TestParserProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	fmt.Println(contents)

	result := ParseProfile(contents, "可惜我爱怀念","http://album.zhenai.com/u/1558719774")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 "+"element; but was %v", result.Items)
	}

	actual := result.Items[0]
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

	if actual != expected {
		t.Errorf("Items should be %v; but was %v", actual, expected)
	}
}
