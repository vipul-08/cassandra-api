package app

import (
	"bytes"
	"encoding/json"
	"github.com/gocql/gocql"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/vipul-08/cassandra-api/domain"
	"github.com/vipul-08/cassandra-api/exception"
	"github.com/vipul-08/cassandra-api/mocks/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *mux.Router
var mockHandlers StudentHandlers
var mockService *service.MockStudentService

func setup(t *testing.T) func() {
	mockController := gomock.NewController(t)
	mockService = service.NewMockStudentService(mockController)
	mockHandlers = StudentHandlers{mockService}
	router = mux.NewRouter()

	return func() {
		router = nil
		defer mockController.Finish()
	}
}

func Test_get_all_students_test_200(t *testing.T)  {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	id1, id2 := gocql.TimeUUID(), gocql.TimeUUID()
	dummyStudentsList := []domain.Student {
		{
			&id1,
			"Vipul Singh Raghuvanshi",
			23,
			"BE",
			"CE",
		},
		{
			&id2,
			"Vikram Parmar",
			22,
			"BE",
			"EXTC",
		},
	}
	mockService.EXPECT().GetAllStudents().Return(dummyStudentsList, nil)
	router.HandleFunc("/students", mockHandlers.getAllStudents).Methods(http.MethodGet)
	req,_ := http.NewRequest(http.MethodGet, "/students", nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_get_all_students_test_500(t *testing.T)  {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	mockService.EXPECT().GetAllStudents().Return(nil, exception.NewUnexpectedError("Unexpected DB Error"))
	router.HandleFunc("/students", mockHandlers.getAllStudents).Methods(http.MethodGet)
	req,_ := http.NewRequest(http.MethodGet, "/students", nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}

func Test_get_specific_student_404(t *testing.T)  {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	id := gocql.TimeUUID()
	mockService.EXPECT().GetStudentById(id).Return(nil, exception.NewNotFoundError("Student Not Found"))
	router.HandleFunc("/students/{id}", mockHandlers.getStudentById).Methods(http.MethodGet)
	req,_ := http.NewRequest(http.MethodGet, "/students/"+id.String(), nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//Assert
	if recorder.Code != http.StatusNotFound {
		t.Error("Failed while testing the status code")
	}
}

func Test_get_specific_student_200(t *testing.T)  {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	id := gocql.TimeUUID()
	student := domain.Student{
		Id:     &id,
		Name:   "Vipul",
		Age:    23,
		Class:  "BE",
		Branch: "CE",
	}

	mockService.EXPECT().GetStudentById(id).Return(&student, nil)
	router.HandleFunc("/students/{id}", mockHandlers.getStudentById).Methods(http.MethodGet)
	req,_ := http.NewRequest(http.MethodGet, "/students/"+id.String(), nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_post_specific_student_404(t *testing.T)  {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	student := domain.Student{
		Name:   "Vipul",
		Age:    23,
		Class:  "BE",
		Branch: "CE",
	}

	mockService.EXPECT().InsertStudent(&student).Return(nil, exception.NewNotFoundError("Unable to Add"))
	router.HandleFunc("/students", mockHandlers.insertStudent).Methods(http.MethodPost)
	body,_ := json.Marshal(student)
	req,_ := http.NewRequest(http.MethodPost, "/students", bytes.NewReader(body))

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//Assert
	if recorder.Code != http.StatusNotFound {
		t.Error("Failed while testing the status code")
	}
}

func Test_post_specific_student_201(t *testing.T)  {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	id := gocql.TimeUUID()
	student := domain.Student{
		Name:   "Vipul",
		Age:    23,
		Class:  "BE",
		Branch: "CE",
	}

	studentWithId := student
	studentWithId.Id = &id

	mockService.EXPECT().InsertStudent(&student).Return(&studentWithId, nil)
	router.HandleFunc("/students", mockHandlers.insertStudent).Methods(http.MethodPost)
	body,_ := json.Marshal(student)
	req,_ := http.NewRequest(http.MethodPost, "/students", bytes.NewReader(body))

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//Assert
	if recorder.Code != http.StatusCreated {
		t.Error("Failed while testing the status code")
	}
}

func Test_put_specific_student_404(t *testing.T)  {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	id := gocql.TimeUUID()
	student := domain.Student{
		Name:   "Vipul",
		Age:    23,
		Class:  "BE",
		Branch: "CE",
	}

	mockService.EXPECT().UpdateStudent(&student, id).Return(nil, exception.NewNotFoundError("Unable to Update"))
	router.HandleFunc("/students/{id}", mockHandlers.updateStudent).Methods(http.MethodPut)
	body,_ := json.Marshal(student)
	req,_ := http.NewRequest(http.MethodPut, "/students/"+id.String(), bytes.NewReader(body))

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//Assert
	if recorder.Code != http.StatusNotFound {
		t.Error("Failed while testing the status code")
	}
}

func Test_put_specific_student_200(t *testing.T)  {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	id := gocql.TimeUUID()
	student := domain.Student{
		Name:   "Vipul",
		Age:    23,
		Class:  "BE",
		Branch: "CE",
	}

	studentWithId := student
	studentWithId.Id = &id

	mockService.EXPECT().UpdateStudent(&student, id).Return(&studentWithId, nil)
	router.HandleFunc("/students/{id}", mockHandlers.updateStudent).Methods(http.MethodPut)
	body,_ := json.Marshal(student)
	req,_ := http.NewRequest(http.MethodPut, "/students/"+id.String(), bytes.NewReader(body))

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_delete_specific_student_404(t *testing.T)  {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	id := gocql.TimeUUID()
	mockService.EXPECT().DeleteStudent(id).Return(exception.NewNotFoundError("No rows affected"))
	router.HandleFunc("/students/{id}", mockHandlers.deleteStudent).Methods(http.MethodDelete)
	req,_ := http.NewRequest(http.MethodDelete, "/students/"+id.String(), nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//Assert
	if recorder.Code != http.StatusNotFound {
		t.Error("Failed while testing the status code")
	}
}

func Test_delete_specific_student_204(t *testing.T)  {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	id := gocql.TimeUUID()
	mockService.EXPECT().DeleteStudent(id).Return(nil)
	router.HandleFunc("/students/{id}", mockHandlers.deleteStudent).Methods(http.MethodDelete)
	req,_ := http.NewRequest(http.MethodDelete, "/students/"+id.String(), nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//Assert
	if recorder.Code != http.StatusNoContent {
		t.Error("Failed while testing the status code")
	}
}
