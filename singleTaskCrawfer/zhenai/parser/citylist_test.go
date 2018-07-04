package parser

import (
	"testing"
	"io/ioutil"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)
	const resultSize = 470
	expectedUls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCitys := []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d"+"requests; but had %d", resultSize, len(result.Requests))
	}

	for i, url := range expectedUls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but "+"was %s", i, url, result.Requests[i].Url)
		}
	}

	for i, city := range expectedCitys {
		if result.Items[i].(string) != city {
			t.Errorf("expected city #%d: %s; but "+"was %s", i, city, result.Items[i].(string))
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d"+"requests; but had %d", resultSize, len(result.Items))
	}
}