package hi

import (
	"fmt"
	"net/http"
	"sync"
)

type Server struct {
	mux  *http.ServeMux
	once sync.Once
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, DevSecOps!</h1>")
}
