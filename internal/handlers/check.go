package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"bytes"
	"github.com/julienschmidt/httprouter"
	"github.com/Taratukhin/abFactorialApi/model"
)

func Check(next httprouter.Handle) httprouter.Handle { // middleware for checking arguments
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, `{"error":"Incorrect input"}`, http.StatusBadRequest)
			return
		}
		r.Body.Close()
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // restore the body
		var data CheckDataType
		err = json.Unmarshal(bodyBytes, &data)
		if err != nil || data.A == nil || data.B == nil || *data.A < 0 || *data.B < 0 {
			http.Error(w, `{"error":"Incorrect input"}`, http.StatusBadRequest)
			return
		}
		next(w, r, ps)
	}
}