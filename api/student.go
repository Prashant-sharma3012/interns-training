package api

import (
	"net/http"

	"github.com/interns-training/models"
)

func (api *API) InitStudent() {
	api.BaseRoutes.Student.Handle("", http.HandlerFunc(api.createStudent)).Methods("POST")
	api.BaseRoutes.Student.Handle("/list", http.HandlerFunc(api.getStudents)).Methods("GET")
	api.BaseRoutes.Student.Handle("/{Id}", http.HandlerFunc(api.getStudentById)).Methods("GET")
	api.BaseRoutes.Student.Handle("", http.HandlerFunc(api.updateStudent)).Methods("PUT")
	api.BaseRoutes.Student.Handle("/{Id}", http.HandlerFunc(api.deleteStudent)).Methods("DELETE")
}

func (api *API) createStudent(w http.ResponseWriter, r *http.Request) {
	Student := models.FromJson(r.Body)
	sendResponse(w, api.Srv.CreateStudent(Student))
	return
}

func (api *API) updateStudent(w http.ResponseWriter, r *http.Request) {
	Student := models.FromJson(r.Body)

	sendResponse(w, api.Srv.UpdateStudent(Student))
	return
}

func (api *API) getStudents(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, api.Srv.GetStudents())
	return
}

func (api *API) getStudentById(w http.ResponseWriter, r *http.Request) {
	params := models.GetParams(r)
	sendResponse(w, api.Srv.GetStudentById(params.Id))
	return
}

func (api *API) deleteStudent(w http.ResponseWriter, r *http.Request) {
	params := models.GetParams(r)
	sendResponse(w, api.Srv.DeleteStudent(params.Id))
	return
}
