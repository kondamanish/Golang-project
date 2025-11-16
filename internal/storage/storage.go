package storage

import "github.com/konda-manish/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error) //this is to create a new student interface
	GetStudentById(id int64) (types.Student, error)                  //this is to get a student by id interface
	GetStudents() ([]types.Student, error)                           //this is to get a list of students interface
	DeleteStudent(id int64) error
}
