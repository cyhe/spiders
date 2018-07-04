package parser

import (
	"testing"
	"io/ioutil"
	"spiders/singleTaskCrawfer/model"
	"fmt"
)

func TestParserProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	fmt.Println(contents)

	result := ParserProfile(contents, "可惜我爱怀念")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 "+"element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
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
	}

	if profile != expected {
		t.Errorf("Items should be %v; but was %v", profile, expected)
	}
}
