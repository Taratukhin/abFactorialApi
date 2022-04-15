package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	. "github.com/Taratukhin/abFactorialApi/internal/functions"
	. "github.com/Taratukhin/abFactorialApi/internal/handlers"
	. "github.com/Taratukhin/abFactorialApi/internal/model"
)

func main() {
	router := httprouter.New()
	router.GET("/calculate", Check(Calculate))
	http.ListenAndServe(":8989", router)
}
