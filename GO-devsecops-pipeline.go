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
		a.mux.HandleFunc("/login", a.login)
		a.mux.HandleFunc("/admin", cookieAuthMw(a.admin))
		a.mux.HandleFunc("/header-admin", headerAuthMw(a.admin))
	})
	a.mux.ServeHTTP(w, r)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, DevSecOps!</h1>")
}
