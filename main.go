package main

import (
	"log"
	"net/http"

	"github.com/inari111/layered-architecture-example-2020/di"
)

func main() {
	http.Handle("/", di.InitializeAPIHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
