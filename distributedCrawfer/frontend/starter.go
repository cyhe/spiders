package main

import (
	"net/http"
	"spiders/distributedCrawfer/frontend/controller"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("frontend/resources/views")))

	http.Handle("/search",
		controller.CreateSearchresultHandler("frontend/resources/views/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
