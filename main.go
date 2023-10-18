package main

import (
	"fmt"
	"net/http"

	error "github.com/blohny/helper"
	"github.com/julienschmidt/httprouter"
)

func main() {

	routes := httprouter.New()

	routes.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "GOLANG")
	})

	server := http.Server{Addr: "localhost: 8888", Handler: routes}

	err := server.ListenAndServe()
	error.PanicIfError(err)
}
