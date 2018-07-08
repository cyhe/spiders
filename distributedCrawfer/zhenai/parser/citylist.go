/*
	citylist的解析器
*/
package parser

import (
	"spiders/distributedCrawfer/engine"
	"regexp"
	"spiders/distributedCrawfer/config"
)

//// pingxiang2 [0-9a-z]+
//// class=""   [^>]*  匹配到多个非>  遇到>暂停
//// 萍乡 [^<]+
//// <a href="http://www.zhenai.com/zhenghun/pingxiang2"class="">萍乡</a>
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	//limit := 10
	for _, m := range matches {
		//result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			Parser: engine.NewFuncParser(ParseCity,config.ParseCity),
		})
		//limit--
		//if limit == 0 {
		//	break;
		//}
	}
	return result
}
