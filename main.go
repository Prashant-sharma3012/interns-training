package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/interns-training/helpers"
	"github.com/interns-training/server"
	log "github.com/sirupsen/logrus"
)

func main() {

	// setup logger
	logger := log.New()

	logger.Info("Setting things up")
	logger.Info("Load Config")

	//load config
	config := helpers.LoadConfig()

	srv := &server.Server{}

	logger.Info("Creating end points")
	// setup routes
	r := mux.NewRouter()
	r.HandleFunc("/", srv.RootHandler)

	logger.Info("creating server")
	// listen and serve
	s := &http.Server{
		Addr:           config.Port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())

	logger.Info("Listening on port 3000")
}
