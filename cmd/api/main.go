package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/renvins/gocoins_restful/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)        // to get file and line number in logs
	var r *chi.Mux = chi.NewRouter() // pointer to a new router instance
	handlers.Handler(r)              // pass the router to the handler function

	fmt.Println("Starting GoCoins RESTful API server...")
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
