package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	mytypes "github.com/Harikrishnasinh/go-students-api/internal/myTypes"
	"github.com/Harikrishnasinh/go-students-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student mytypes.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			fmt.Println("it is here,", err.Error())
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("Empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// Request validation
		if err := validator.New().Struct(student); err != nil {
			// This is type casting from err to validator.ValidationErrors
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		response.WriteJson(w, http.StatusOK, student)
	}
}
