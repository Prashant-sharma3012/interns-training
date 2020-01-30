package models

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Param struct {
	Id string
}

func GetParams(r *http.Request) *Param {
	params := &Param{}
	vars := mux.Vars(r)

	if val, ok := vars["Id"]; ok {
		params.Id = val
	}

	return params
}
