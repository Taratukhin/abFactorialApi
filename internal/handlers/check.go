package handlers

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"bytes"
	"github.com/julienschmidt/httprouter"
	. "github.com/Taratukhin/abFactorialApi/internal/model"
)

func Check(next httprouter.Handle) httprouter.Handle { // middleware for checking arguments
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-control", "no-cashe")
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"Incorrect input"}`))
			return
		}
		r.Body.Close()
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // restore the body
		var data CheckDataType
		err = json.Unmarshal(bodyBytes, &data)
		if err != nil || data.A == nil || data.B == nil || *data.A < 0 || *data.B < 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"Incorrect input"}`))
			return
		}
		next(w, r, ps)
	}
}
