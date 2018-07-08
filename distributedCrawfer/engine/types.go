package engine

type ParserFunc func(contents []byte, url string) ParserResult

type Parser interface {
	Parse(contents []byte, url string) ParserResult
	Serialized() (name string, args interface{})
}

type Request struct {
	Url    string
	Parser Parser
}

type ParserResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

// NilParser 返回一个空的ParserResult
//func NilParser([]byte) ParserResult {
//	return ParserResult{}
//}

type NilParser struct {
}

func (NilParser) Parse(_ []byte, _ string) ParserResult {
	return ParserResult{}

}

func (NilParser) Serialized() (name string, args interface{}) {
	return "NilParser", nil
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParserResult {

	return f.parser(contents, url)
}

func (f *FuncParser) Serialized() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}

}
