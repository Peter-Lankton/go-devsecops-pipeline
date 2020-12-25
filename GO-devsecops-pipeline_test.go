package hi

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	Home(w, r)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("ioutil.ReadAll() err = %s; want nil", err)
	}
	have := string(body)
	want := "<h1>Hello, DevSecOps!</h1>"
	if have != want {
		t.Errorf("GET / = %s; want %s", have, want)
	}
}
