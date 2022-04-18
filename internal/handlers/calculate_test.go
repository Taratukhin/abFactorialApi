package handlers

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"strings"
)

func TestCalculate(t *testing.T) {

	data := []struct {
		arg, res string
		code     int
	}{
		{`{"a":3,"b":10}`, `{"a":6,"b":3628800}`, 200},
		{`{"a":3,"b":-10}`, `{"error":"Incorrect input"}`, 400}, // -10 is not uint64
		{`{"a":3}`, `{"a":6,"b":1}`, 200}, // the argument b is 0 by default
	}
	for _, d := range data {
		req := httptest.NewRequest("GET", "http://localhost/calculate", strings.NewReader(d.arg))
		w := httptest.NewRecorder()
		Calculate(w, req, nil)
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		if strings.TrimSpace(string(body)) != d.res || resp.StatusCode!= d.code{
			t.Errorf("Calculate(\"%s\") return \"%s\" code=%d, want \"%s\" code=%d", d.arg, strings.TrimSpace(string(body)),resp.StatusCode, d.res, d.code)
		}
	}

}
