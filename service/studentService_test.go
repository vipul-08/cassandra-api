package service

import (
	"github.com/gocql/gocql"
	"github.com/golang/mock/gomock"
	mainDomain "github.com/vipul-08/cassandra-api/domain"
	"github.com/vipul-08/cassandra-api/exception"
	"github.com/vipul-08/cassandra-api/mocks/domain"
	"testing"
)

var mockRepository *domain.MockStudentRepository
var service StudentService

func setup(t *testing.T) func()  {
	mockController := gomock.NewController(t)
	mockRepository = domain.NewMockStudentRepository(mockController)
	service = NewStudentService(mockRepository)
	return func() {
		service = nil
		defer mockController.Finish()
	}
}

func Test_create_student_successfully(t *testing.T) {
	// Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	//Act
	id := gocql.TimeUUID()
	payloadStudent := mainDomain.Student{
		Id:		&id,
		Name:   "Vipul Singh Raghuvanshi",
		Age:    24,
		Class:  "BE",
		Branch: "CE",
	}

	mockRepository.EXPECT().Insert(&payloadStudent).Return(&payloadStudent, nil)
	newStudent,appError := service.InsertStudent(&payloadStudent)

	//Assert
	if appError != nil {
		t.Error("Test failed while creating account")
	}
	if newStudent.Id != payloadStudent.Id {
		t.Error("Failed while matching new student id")
	}
}

func Test_create_student_failed(t *testing.T) {
	// Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	//Act
	id := gocql.TimeUUID()
	payloadStudent := mainDomain.Student{
		Id:		&id,
		Name:   "Vipul Singh Raghuvanshi",
		Age:    24,
		Class:  "BE",
		Branch: "CE",
	}

	mockRepository.EXPECT().Insert(&payloadStudent).Return(nil, exception.NewNotFoundError("Unable to Add"))
	_,appError := service.InsertStudent(&payloadStudent)

	//Assert
	if appError == nil {
		t.Error("Test failed while creating student")
	}
}

func Test_get_student_failed(t *testing.T) {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	//Act
	id := gocql.TimeUUID()
	mockRepository.EXPECT().FindById(id).Return(nil, exception.NewNotFoundError("Student Not Found"))
	_,appError := service.GetStudentById(id)

	//Assert
	if appError == nil {
		t.Error("Test failed while getting specific student")
	}
}

func Test_get_student_success(t *testing.T) {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	//Act
	id := gocql.TimeUUID()
	student := mainDomain.Student{
		Id: 	&id,
		Name:   "Vipul Singh Raghuvanshi",
		Age:    24,
		Class:  "BE",
		Branch: "CE",
	}
	mockRepository.EXPECT().FindById(id).Return(&student, nil)
	_,appError := service.GetStudentById(id)

	//Assert
	if appError != nil {
		t.Error("Test failed while getting specific student")
	}
}

func Test_update_student_success(t *testing.T) {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	//Act
	id := gocql.TimeUUID()
	student := mainDomain.Student{
		Name:   "Vipul Singh Raghuvanshi",
		Age:    24,
		Class:  "BE",
		Branch: "CE",
	}
	mockRepository.EXPECT().Update(&student, id).Return(&student, nil)
	_,appError := service.UpdateStudent(&student, id)

	//Assert
	if appError != nil {
		t.Error("Test failed while updating specific student")
	}
}

func Test_update_student_failure(t *testing.T) {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	//Act
	id := gocql.TimeUUID()
	student := mainDomain.Student{
		Name:   "Vipul Singh Raghuvanshi",
		Age:    24,
		Class:  "BE",
		Branch: "CE",
	}
	mockRepository.EXPECT().Update(&student, id).Return(nil, exception.NewNotFoundError("Unable to Update"))
	_,appError := service.UpdateStudent(&student, id)

	//Assert
	if appError == nil {
		t.Error("Test failed while updating specific student")
	}
}

func Test_delete_student_failure(t *testing.T) {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	//Act
	id := gocql.TimeUUID()
	mockRepository.EXPECT().Delete(id).Return(exception.NewNotFoundError("No rows affected"))
	appError := service.DeleteStudent(id)

	//Assert
	if appError == nil {
		t.Error("Test failed while deleting specific student")
	}
}

func Test_delete_student_success(t *testing.T) {
	//Arrange
	callOnFinish := setup(t)
	defer callOnFinish()

	//Act
	id := gocql.TimeUUID()
	mockRepository.EXPECT().Delete(id).Return(nil)
	appError := service.DeleteStudent(id)

	//Assert
	if appError != nil {
		t.Error("Test failed while deleting specific student")
	}
}


