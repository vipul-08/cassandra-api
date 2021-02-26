package domain

import (
	"github.com/gocql/gocql"
	"github.com/vipul-08/cassandra-api/exception"
)

type Student struct {
	Id     *gocql.UUID  `json:"id" `
	Name   string 		`json:"name"`
	Age    uint8  		`json:"age"`
	Class  string 		`json:"class"`
	Branch string 		`json:"branch"`
}

// swagger:parameters getStudent deleteStudent updateStudent
type studentIdParameter struct {
	// The id of a specific student
	// in: path
	// required: true
	ID string `json:"id"`
}

// A list of students returns in response
// swagger:response studentsResponse
type studentsResponse struct {
	// All students in the DB
	// in: body
	Body []struct{
		Id     string	`json:"id" `
		Name   string 	`json:"name"`
		Age    uint8  	`json:"age"`
		Class  string 	`json:"class"`
		Branch string 	`json:"branch"`
	}
}

// swagger:response noContent
type studentNoContent struct {}


// A student object returns in response
// swagger:response studentResponse
type studentResponse struct {
	// Specific student in the DB
	// in: body
	Body struct{
		Id     string	`json:"id" `
		Name   string 	`json:"name"`
		Age    uint8  	`json:"age"`
		Class  string 	`json:"class"`
		Branch string 	`json:"branch"`
	}
}

// swagger:parameters updateStudent createStudent
type studentParams struct {
	// Student data structure to Update or Create.
	// in: body
	// required: true
	Body struct{
		Name   string 	`json:"name"`
		Age    uint8  	`json:"age"`
		Class  string 	`json:"class"`
		Branch string 	`json:"branch"`
	}
}

//go:generate mockgen -destination=../mocks/domain/mock_StudentRepository.go -package=domain github.com/vipul-08/cassandra-api/domain StudentRepository
type StudentRepository interface {
	FindAll() ([]Student, *exception.AppError)
	FindById(id gocql.UUID) (*Student, *exception.AppError)
	Insert(student *Student) (*Student, *exception.AppError)
	Update(student *Student, id gocql.UUID) (*Student, *exception.AppError)
	Delete(id gocql.UUID) *exception.AppError
}
