package main

import (
	"net/http"

	hi "github.com/llcranmer/GO-devsecops-pipeline"
)

func main() {
	http.ListenAndServe(":3000", &hi.Server{})
}
