package server

import (
	"net/http"

	"github.com/interns-training/models"
)

func (a *Server) CreateStudent(Student *models.Student) *models.ApiResponse {

	_, err := a.Repo.Student.Add(Student)

	if err != nil {
		return models.NewFRes(http.StatusInternalServerError, "Failed to save Student", "CreateStudent.Student.Save."+err.Error(), models.ERRCODE7)
	}

	return models.NewSRes(http.StatusCreated, "New Student added to the system", Student)
}

func (a *Server) UpdateStudent(Student *models.Student) *models.ApiResponse {

	_, err := a.Repo.Student.Update(Student)

	if err != nil {
		return models.NewFRes(http.StatusInternalServerError, "Failed to update Student", "CreateStudent.Student.Update."+err.Error(), models.ERRCODE7)
	}

	return models.NewSRes(http.StatusCreated, "Student Updated", map[string]string{
		"StudentId": Student.Id,
	})
}

func (a *Server) GetStudents() *models.ApiResponse {

	var err error

	Students, err := a.Repo.Student.List()
	if err != nil {
		return models.NewFRes(http.StatusInternalServerError, "Failed to fetch Students", "GetStudents.Student.GetAll."+err.Error(), models.ERRCODE7)
	}

	return models.NewSRes(http.StatusOK, "All Students", Students)
}

func (a *Server) GetStudentById(StudentId string) *models.ApiResponse {

	var err error

	Student, err := a.Repo.Student.GetStudentById(StudentId)
	if err != nil {
		return models.NewFRes(http.StatusInternalServerError, "Failed to fetch Students", "GetStudentById.Student.GetStudentById."+err.Error(), models.ERRCODE7)
	}

	return models.NewSRes(http.StatusOK, "Student by Id", Student)
}

func (a *Server) DeleteStudent(StudentId string) *models.ApiResponse {

	var err error

	Student, err := a.Repo.Student.Delete(StudentId)
	if err != nil {
		return models.NewFRes(http.StatusInternalServerError, "Failed to fetch Students", "GetStudentById.Student.GetStudentById."+err.Error(), models.ERRCODE7)
	}

	return models.NewSRes(http.StatusOK, "Student by Id", Student)
}
