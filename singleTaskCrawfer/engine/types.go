package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request
	Items    []interface{}
}

// NilParser 返回一个空的ParserResult
func NilParser([]byte) ParserResult {
	return ParserResult{}
}
