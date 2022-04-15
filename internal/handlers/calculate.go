package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	. "github.com/Taratukhin/abFactorialApi/internal/functions"
	. "github.com/Taratukhin/abFactorialApi/internal/model"
)

func Calculate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var data DataType
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil { // this check is superfluous, but what happens if we remove the middleware?
		http.Error(w, `{"error":"Incorrect input"}`, http.StatusBadRequest)
		return
	}
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	go func () {
		data.A = Factorial(data.A)
		c1 <- struct{}{}
	}()
	go func () {
		data.B = Factorial(data.B)
		c2 <- struct{}{}
	}()
	for i:=0;i<2;i++ {
		select {
			case <-c1 :
			case <-c2 :
		}
	}
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, `{"error":"Incorrect output"}`, http.StatusBadRequest)
		return
	}
}
