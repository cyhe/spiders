package parser

import (
	"regexp"
	"spiders/distributedCrawfer/engine"
	"spiders/distributedCrawfer/config"
)

// const cityRe = `<a href="http://album.zhenai.com/u/1646394243" target="_blank">不点</a>`
const profileRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
const cityUrlRe = `href="(http://www.zhenai.com/zhenghun/[^"]+)"`

var (
	profilePattern = regexp.MustCompile(profileRe)
	cityUrlPattern = regexp.MustCompile(cityUrlRe)
)

func ParseCity(contents []byte, _ string) engine.ParserResult {
	matches := profilePattern.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			//ParserFunc: ProfileParser(string(m[2])),
			Parser:NewProfileParser(string(m[2])),
		})
	}

	matches = cityUrlPattern.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			Parser: engine.NewFuncParser(ParseCity,config.ParseCity),
		})
	}

	return result
}
