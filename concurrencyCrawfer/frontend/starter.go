package main

import (
	"net/http"
	"spiders/concurrencyCrawfer/frontend/controller"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("frontend/view")))

	http.Handle("/search",
		controller.CreateSearchresultHandler("frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
