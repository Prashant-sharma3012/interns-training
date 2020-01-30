package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/interns-training/models"
	"github.com/interns-training/server"
)

type Routes struct {
	Root    *mux.Router // ''
	ApiRoot *mux.Router // 'api/v1'
	Student *mux.Router // 'api/v1/product'
}
type API struct {
	Srv        *server.Server
	BaseRoutes *Routes
}

func Init(a *server.Server) {
	api := &API{
		Srv:        a,
		BaseRoutes: &Routes{},
	}

	api.BaseRoutes.Root = a.Handler
	api.BaseRoutes.ApiRoot = a.Handler.PathPrefix("/api/v1").Subrouter()
	api.BaseRoutes.Student = api.BaseRoutes.ApiRoot.PathPrefix("/student").Subrouter()

	a.Handler.Handle("/api/v1/{anything:.*}", http.HandlerFunc(PreFlightHandle))

	api.InitStudent()
}

func PreFlightHandle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, x-access-token")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("There doesn't appear to be an api call for the url='" + r.URL.Path + "'."))
	}

}

func sendResponse(w http.ResponseWriter, r *models.ApiResponse) {
	w.WriteHeader(r.StatusCode)
	w.Write([]byte(r.ToJson()))
	return
}
