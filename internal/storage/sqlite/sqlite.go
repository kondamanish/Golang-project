package sqlite

import (
	"database/sql"

	"fmt"

	"github.com/konda-manish/internal/config"
	"github.com/konda-manish/internal/types"
	_ "github.com/mattn/go-sqlite3" //sqlite3 driver
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students 
	(id INTEGER PRIMARY KEY AUTOINCREMENT, 
	name TEXT NOT NULL, 
	email TEXT NOT NULL, 
	age INTEGER NOT NULL)`) //this is to create the students table if it doesn't exist

	if err != nil {
		return nil, err
	}

	return &Sqlite{Db: db}, nil

}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	stmt, err := s.Db.Prepare(`INSERT INTO students (name, email, age) values (?, ?, ?)`) //this is to prepare the statement for the execution
	if err != nil {
		return 0, err
	}
	defer stmt.Close() //this is to close the statement after the execution

	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId() //tis is part of the sqlite driver to get the last inserted id

	if err != nil {
		return 0, err
	}

	return lastId, nil //this is to return the last inserted id
}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare(`SELECT id, name, email, age FROM students WHERE id = ? LIMIT 1`)
	if err != nil {
		return types.Student{}, err
	}
	defer stmt.Close()

	var student types.Student
	err = stmt.QueryRow(id).Scan(&student.ID, &student.Name, &student.Email, &student.Age) //this is to scan the result into the student struct
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("student not found: %w", err)
		}
		return types.Student{}, fmt.Errorf(" query error: %w", err) //this is to return the query error
	}

	return student, nil
}

func (s *Sqlite) GetStudents() ([]types.Student, error) {
	rows, err := s.Db.Query(`SELECT id, name, email, age FROM students`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []types.Student

	for rows.Next() {
		var student types.Student
		err = rows.Scan(&student.ID, &student.Name, &student.Email, &student.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func (s *Sqlite) DeleteStudent(id int64) error {
	stmt, err := s.Db.Prepare(`DELETE FROM students WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return err
	}

	return nil
}
