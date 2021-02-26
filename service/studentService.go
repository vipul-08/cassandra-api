package service

import (
	"github.com/gocql/gocql"
	"github.com/vipul-08/cassandra-api/domain"
	"github.com/vipul-08/cassandra-api/exception"
)

//go:generate mockgen -destination=../mocks/service/mock_StudentService.go -package=service github.com/vipul-08/cassandra-api/service StudentService
type StudentService interface {
	GetAllStudents() ([]domain.Student, *exception.AppError)
	GetStudentById(id gocql.UUID) (*domain.Student, *exception.AppError)
	InsertStudent(s *domain.Student) (*domain.Student, *exception.AppError)
	UpdateStudent(s *domain.Student, id gocql.UUID) (*domain.Student, *exception.AppError)
	DeleteStudent(id gocql.UUID) *exception.AppError
}

type DefaultStudentService struct {
	repo domain.StudentRepository
}

func (service DefaultStudentService) GetAllStudents() ([]domain.Student, *exception.AppError) {
	return service.repo.FindAll()
}

func (service DefaultStudentService) GetStudentById(id gocql.UUID) (*domain.Student, *exception.AppError) {
	return service.repo.FindById(id)
}

func (service DefaultStudentService) InsertStudent(s *domain.Student) (*domain.Student, *exception.AppError) {
	return service.repo.Insert(s)
}

func (service DefaultStudentService) UpdateStudent(s *domain.Student, id gocql.UUID) (*domain.Student, *exception.AppError) {
	return service.repo.Update(s, id)
}

func (service DefaultStudentService) DeleteStudent(id gocql.UUID) *exception.AppError {
	return service.repo.Delete(id)
}

func NewStudentService(repository domain.StudentRepository) DefaultStudentService {
	return DefaultStudentService{repository}
}
