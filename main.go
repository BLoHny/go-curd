package main

import (
	"net/http"

	"github.com/blohny/config"
	"github.com/blohny/controller"
	error "github.com/blohny/helper"
	"github.com/blohny/repository"
	"github.com/blohny/router"
	"github.com/blohny/service"
)

func main() {

	db := config.DatabaseConnection()

	bookRepository := repository.NewBookRepository(db)

	bookService := service.NewBookServiceImpl(bookRepository)

	bookController := controller.NewBookController(bookService)

	routes := router.NewRouter(bookController)

	server := http.Server{Addr: "localhost: 8888", Handler: routes}

	err := server.ListenAndServe()
	error.PanicIfError(err)
}
