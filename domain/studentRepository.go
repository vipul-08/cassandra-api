package domain

import (
	"github.com/gocql/gocql"
	"github.com/mitchellh/mapstructure"
	"github.com/vipul-08/cassandra-api/database"
	"github.com/vipul-08/cassandra-api/exception"
)

type StudentRepositoryDb struct {
	client database.DbConnection
}

func NewStudentRepository() StudentRepositoryDb {
	client := database.GetDbConnection()
	return StudentRepositoryDb{client}
}

func (d StudentRepositoryDb) FindAll() ([]Student, *exception.AppError) {
	students := make([]Student, 0)
	rows, err := d.client.Session.Query(`SELECT * FROM students`).Iter().SliceMap()
	if err != nil {
		return students, exception.NewUnexpectedError("Unable to fetch students")
	}
	for _,row := range rows {
		var student Student
		e := mapstructure.Decode(row, &student)
		if e != nil {
			return students, exception.NewUnexpectedError("Unable to parse Student")
		}
		students = append(students, student)
	}
	return students, nil
}

func (d StudentRepositoryDb) FindById(id gocql.UUID) (*Student, *exception.AppError) {
	rows, err := d.client.Session.Query(`SELECT * FROM students WHERE id = ? LIMIT 1`, id).Consistency(gocql.One).Iter().SliceMap()
	if err != nil {
		return nil, exception.NewUnexpectedError("Unable to fetch student")
	}
	var student Student
	found := false
	for _,row := range rows {
		found = true
		e := mapstructure.Decode(row, &student)
		if e != nil {
			return nil, exception.NewUnexpectedError("Unable to parse Student")
		}
	}
	if !found {
		return nil, exception.NewNotFoundError("Student not found")
	}
	return &student, nil
}

func (d StudentRepositoryDb) Insert(s *Student) (*Student, *exception.AppError) {
	uuid := gocql.TimeUUID()
	err := d.client.Session.Query(
		`INSERT INTO students (id, name, age, class, branch) VALUES (?, ?, ?, ?, ?)`,
		uuid,
		s.Name,
		s.Age,
		s.Class,
		s.Branch,
	).Exec()

	if err != nil {
		return nil, exception.NewUnexpectedError("Unable to insert Row")
	}
	s.Id = &uuid

	return s,nil
}

func (d StudentRepositoryDb) Update(s *Student, id gocql.UUID) (*Student, *exception.AppError) {
	err := d.client.Session.Query(
		`UPDATE students SET name = ?, age = ?, class = ?, branch = ? WHERE id = ?`,
		s.Name,
		s.Age,
		s.Class,
		s.Branch,
		id,
	).Exec()
	if err != nil {
		return nil, exception.NewUnexpectedError("Unable to update student with this id")
	}
	s.Id = &id
	return s,nil
}

func (d StudentRepositoryDb) Delete(id gocql.UUID) *exception.AppError {
	rows, err := d.client.Session.Query(`SELECT * FROM students WHERE id = ? LIMIT 1`, id).Consistency(gocql.One).Iter().SliceMap()
	if err != nil {
		return exception.NewUnexpectedError("Unable to fetch student")
	}
	found := false
	for range rows {
		found = true
	}
	if !found {
		return exception.NewNotFoundError("Unable to find student with this id")
	}
	err = d.client.Session.Query(`DELETE FROM students WHERE id = ?`, id).Exec()
	if err != nil {
		return exception.NewNotFoundError("Unable to delete student with this id")
	}
	return nil
}
