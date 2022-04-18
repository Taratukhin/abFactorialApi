package main

import (
	"testing"
	"io/ioutil"
	"strings"
	"net/http"
	"net/http/httptest"
	"github.com/julienschmidt/httprouter"
	. "github.com/Taratukhin/abFactorialApi/internal/handlers"

)

func TestServer(t *testing.T) {

	data := []struct {
		arg, res string
		code     int
	}{
		{`{"a":2,"b":10}`, `{"a":2,"b":3628800}`, 200},
		{`{"a":3,"b":-10}`, `{"error":"Incorrect input"}`, 400},
		{`{"a":2}`, `{"error":"Incorrect input"}`, 400}, 
	}
	router := httprouter.New()
	router.GET("/calculate", Check(Calculate))
	ts := httptest.NewServer(router)
	defer ts.Close()

	for _, d := range data {
		req, err := http.NewRequest("GET", ts.URL+"/calculate", strings.NewReader(d.arg))
    		if err != nil {
        		t.Error(err)
			return
    		}
		client := &http.Client{}
    		resp, err := client.Do(req)
    		if err != nil {
        		t.Error(err)
			return
    		}
		body, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if strings.TrimSpace(string(body)) != d.res || resp.StatusCode!= d.code{
			t.Errorf("GET \"%s\" return \"%s\" code=%d, want \"%s\" code=%d", d.arg, strings.TrimSpace(string(body)),resp.StatusCode, d.res, d.code)
		}
	}

}
