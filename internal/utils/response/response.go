package response

import (
	"encoding/json"
	"net/http"

	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOK    = "ok"
	StatusError = "error"
)

// WriteJson is a helper function to write a json response
func WriteJson(w http.ResponseWriter, status int, data interface{}) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)

}

// GeneralError is a helper function to return a general error response
func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errors []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errors = append(errors, fmt.Sprintf("field %s is required", err.Field()))
		default:
			errors = append(errors, fmt.Sprintf("field %s is not valid", err.Field()))
		}

	}
	return Response{
		Status: StatusError,
		Error:  strings.Join(errors, ","),
	}

}
