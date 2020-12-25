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

func (a *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.once.Do(func() {
		a.mux = http.NewServeMux()
		a.mux.HandleFunc("/", a.home)
	})
	a.mux.ServeHTTP(w, r)
}

func (a *Server) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome!</h1>")
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, DevSecOps!</h1>")
}
