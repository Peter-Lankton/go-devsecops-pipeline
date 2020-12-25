package main

import (
	"log"
	"net/http"

	hi "github.com/llcranmer/GO-devsecops-pipeline"
)

func main() {
	err := http.ListenAndServe(":3000", &hi.Server{})
	if err != nil {
		log.Fatal("Server is not working.")
	}
}
