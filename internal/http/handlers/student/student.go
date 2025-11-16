package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/konda-manish/internal/storage"
	"github.com/konda-manish/internal/types"
	"github.com/konda-manish/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student

		//log the request
		slog.Info("Creating a new student")

		//handles the empty request body errors
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty request body")))
			return
		}

		//handles the other json decode errors
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		//request validation
		if err := validator.New().Struct(student); err != nil {
			validationErrors := err.(validator.ValidationErrors) //typeassertion the error to validator.ValidationErrors
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validationErrors))
			return
		}

		lastId, err := storage.CreateStudent(student.Name, student.Email, student.Age) //this is to create a new student

		slog.Info("user created successfully", slog.Int64("id", lastId))
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}

		//handles the successful creation of the student
		response.WriteJson(w, http.StatusCreated, map[string]int64{"Id": lastId})
	} // this is to return a new handler function

}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("Getting a student by id", slog.String("id", id))

		idInt, err := strconv.ParseInt(id, 10, 64) //this is to parse the id to int64
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		student, err := storage.GetStudentById(idInt) //this is to get the student by id
		if err != nil {
			slog.Error("failed to get student by id", slog.String("error", err.Error()))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		response.WriteJson(w, http.StatusOK, student)

	}
}

func GetList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Getting a list of students")
		students, err := storage.GetStudents() //this is to get the list of students
		if err != nil {
			slog.Error("failed to get list of students", slog.String("error", err.Error()))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		response.WriteJson(w, http.StatusOK, students)

	}
}

func DeleteById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("deleting the student by id", slog.String("id", id))
		idInt, err := strconv.ParseInt(id, 10, 64) //this is to parse the id to int64
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		err = storage.DeleteStudent(idInt) //this is to get the student by id
		if err != nil {
			slog.Error("failed to delete by id", slog.String("error", err.Error()))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		response.WriteJson(w, http.StatusOK, map[string]interface{}{"message": "student deleted successfully", "id": idInt})
	}
}
