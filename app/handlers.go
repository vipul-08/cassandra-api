// Package Classification Student API
//
// Documentation of Student API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package app

import (
	"encoding/json"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"github.com/vipul-08/cassandra-api/domain"
	"github.com/vipul-08/cassandra-api/exception"
	"github.com/vipul-08/cassandra-api/service"
	"io/ioutil"
	"net/http"
)

type StudentHandlers struct {
	service service.StudentService
}

// swagger:route GET /students students listStudents
// Returns list of all students
// responses:
//	200: studentsResponse
//

func (handler *StudentHandlers) getAllStudents(w http.ResponseWriter, r *http.Request) {
	students, err := handler.service.GetAllStudents()

	if err != nil {
		writeResponse(w,err.Code, err)
	} else {
		w.Header().Add("Content-Type", "application/json")
		writeResponse(w,http.StatusOK, students)
	}
}

// swagger:route GET /students/{id} students getStudent
// Returns a student object
// responses:
//	200: studentResponse
//

func (handler *StudentHandlers) getStudentById(w http.ResponseWriter, r *http.Request) {
	id, _ := gocql.ParseUUID(mux.Vars(r)["id"])
	student, err := handler.service.GetStudentById(id)

	if err != nil {
		writeResponse(w,err.Code, err)
	} else {
		writeResponse(w,http.StatusOK, student)
	}
}

// swagger:route POST /students students createStudent
// Create a new student
//
// responses:
//	200: studentResponse

func (handler *StudentHandlers) insertStudent(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e := exception.NewUnprocessableEntityError(err.Error())
		writeResponse(w, e.Code, e)
	}
	var student domain.Student
	err = json.Unmarshal(body, &student)

	if err != nil {
		e := exception.NewUnprocessableEntityError(err.Error())
		writeResponse(w, e.Code, e)
	}

	if student.Id != nil {
		e := exception.NewUnprocessableEntityError("Field id should not be present")
		writeResponse(w, e.Code, e)
	}

	s,e := handler.service.InsertStudent(&student)
	if e != nil {
		writeResponse(w, e.Code, e)
	}
	writeResponse(w, http.StatusCreated, s)
}

// swagger:route PUT /students/{id} students updateStudent
// Update a student details
//
// responses:
//	200: studentResponse

func (handler *StudentHandlers) updateStudent(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e := exception.NewUnprocessableEntityError(err.Error())
		writeResponse(w, e.Code, e)
	}
	var student domain.Student
	err = json.Unmarshal(body, &student)

	if err != nil {
		e := exception.NewUnprocessableEntityError(err.Error())
		writeResponse(w, e.Code, e)
	}

	if student.Id != nil {
		e := exception.NewUnprocessableEntityError("Field id should not be present")
		writeResponse(w, e.Code, e)
	}

	id, _ := gocql.ParseUUID(mux.Vars(r)["id"])

	s,e := handler.service.UpdateStudent(&student, id)
	if e != nil {
		writeResponse(w, e.Code, e)
	}
	writeResponse(w, http.StatusOK, s)
}

// swagger:route DELETE /students/{id} students deleteStudent
// Returns a student object
// responses:
//	204: noContent
//

func (handler *StudentHandlers) deleteStudent(w http.ResponseWriter, r *http.Request) {
	id, _ := gocql.ParseUUID(mux.Vars(r)["id"])
	err := handler.service.DeleteStudent(id)
	if err != nil {
		writeResponse(w,err.Code, err)
	} else {
		w.Header().Set("Entity", id.String())
		w.WriteHeader(204)
	}
}

func writeResponse(w http.ResponseWriter, code int, body interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		panic(err)
	}
}
