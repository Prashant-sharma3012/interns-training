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
	srv := &Server{
		Repo:    &repository.Repository{},
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

	logger.Info("Created Server")
	srv.InitRoutes()

	return srv
}

func (s *Server) InitRoutes() {
	s.Handler.HandleFunc("/", s.RootHandler)
}
