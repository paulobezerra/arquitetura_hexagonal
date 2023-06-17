package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bezerrapaulo/go-hexagonal/adapters/web/handler"
	"github.com/bezerrapaulo/go-hexagonal/application"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type Webserver struct {
	Service application.ProductServiceInterface
}

func MakeNewWebserver(service application.ProductServiceInterface) *Webserver {
	return &Webserver{Service: service}
}

func (w Webserver) Start() {

	r := mux.NewRouter()
	n := negroni.New(negroni.NewLogger())

	handler.MakeProductHandlers(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
