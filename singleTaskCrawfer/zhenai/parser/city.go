package parser

import (
	"regexp"
	"spiders/singleTaskCrawfer/engine"
)

// const cityRe = `<a href="http://album.zhenai.com/u/1646394243" target="_blank">不点</a>`
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityRe)
	//matches := re.FindAll(constents, -1)
	matches := re.FindAllSubmatch(contents, -1)
	// [][][]byte
	result := engine.ParserResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParserResult {
				return ParserProfile(c, name)
			},
		})
	}
	return result
}
