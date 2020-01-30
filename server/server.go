package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/interns-training/models"
	"github.com/interns-training/repository"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Repo    *repository.Repository
	Logger  *logrus.Logger
	Config  *models.Config
	Handler *mux.Router
	Server  *http.Server
}

func InitServer(cnf *models.Config, logger *logrus.Logger, handler *mux.Router) *Server {

	logger.Info("Connecting to DB")
	repo, err := repository.Connect(cnf.DB)

	if err != nil {
		logger.Error(err)
	}

	logger.Info("Creating Server")

	srv := &Server{
		Repo:    repo,
		Logger:  logger,
		Config:  cnf,
		Handler: handler,
		Server: &http.Server{
			Addr:           cnf.Port,
			Handler:        handler,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}

	logger.Info("Server Created")

	return srv
}
