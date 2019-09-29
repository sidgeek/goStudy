package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHanndler() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)

	return router
}

func main() {
	r := RegisterHanndler()
	http.ListenAndServe(":8000", r)
}
