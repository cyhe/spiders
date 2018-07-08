package worker

import (
	"spiders/distributedCrawfer/engine"
	"spiders/distributedCrawfer/config"
	"spiders/distributedCrawfer/zhenai/parser"
	"errors"
	"fmt"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

// {"ParseCityList",nil},{"ProfileParser", username}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items   []engine.Item
	Request []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialized()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParserResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		result.Request = append(result.Request, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult(r ParseResult) (engine.ParserResult, error) {
	result := engine.ParserResult{
		Items: r.Items,
	}
	for _, req := range r.Request {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing"+"request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result, nil
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		} else {
			return nil, fmt.Errorf("invalid"+"args: %v", p.Args)
		}
	default:
		return nil, errors.New("unDefine Method")
	}
}

//type CrawService struct{}

//engine.Request 无法再网络上直接传输, 需要包装一层可供传输
//func (CrawService) Process(req engine.Request, result engine.ParserResult) {
//
//}

//func (CrawService) Process(req engine.Request, result engine.ParserResult) {
//
//}
