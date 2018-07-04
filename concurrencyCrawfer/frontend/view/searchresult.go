package view

import (
	"html/template"
	"io"
	"spiders/concurrencyCrawfer/frontend/model"
)

type SearchResultView struct {
	template *template.Template
}

// 创建一个结果页面
func CreateSearchResultView(filename string) SearchResultView {
	return SearchResultView{
		template: template.Must(template.ParseFiles(filename)),
	}
}

// 渲染
func (s SearchResultView)Render(w io.Writer, data model.SearchResult) error  {
	return  s.template.Execute(w,data)
	
}
