package main

import (
	"spiders/singleTaskCrawfer/engine"
	"spiders/singleTaskCrawfer/zhenai/parser"
)

func main() {

	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
//

//
//// printCityList 城市列表解析器
//func printCityList(constents []byte) {
//
//}
