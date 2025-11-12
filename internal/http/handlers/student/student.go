package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/konda-manish/internal/types"
	"github.com/konda-manish/internal/utils/response"
)

func New() http.HandlerFunc {
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

		//handles the successful creation of the student
		response.WriteJson(w, http.StatusCreated, map[string]string{"Success": "OK"})
	} // this is to return a new handler function

}
