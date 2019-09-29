package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/wy3148/autoptesting/routers"

	"github.com/astaxie/beego"
)

func init() {
	path := os.Getenv("GOPATH") + "/src/github.com/wy3148/autoptesting/"
	beego.TestBeegoInit(path)
}

// TestGet is a sample to run an endpoint test
func TestGet(t *testing.T) {
	r, _ := http.NewRequest("GET", "/contact/person_447AA9AE-B311-45E7-9FF2-E846FA5D939E", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	if w.Code != 200 {
		t.Errorf("failed to get contact, error http code %d", w.Code)
	}
	//more testing code can be added here
}
