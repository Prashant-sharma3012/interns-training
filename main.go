package main

import (
	"github.com/gorilla/mux"
	"github.com/interns-training/helpers"
	"github.com/interns-training/server"
	"github.com/sirupsen/logrus"
)

func main() {

	// setup logger
	logger := logrus.New()

	logger.Info("Setting things up")
	logger.Info("Load Config")

	//load config
	config := helpers.LoadConfig()

	logger.Info("Creating end points")
	// setup routes
	r := mux.NewRouter()
	srv := server.InitServer(config, logger, r)

	logger.Info("Listening on port " + srv.Config.Port)

	logger.Fatal(srv.Server.ListenAndServe())
}
