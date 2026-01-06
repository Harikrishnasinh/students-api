package student

import (
	"encoding/json"
	"net/http"

	mytypes "github.com/Harikrishnasinh/go-students-api/internal/myTypes"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student mytypes.Student
		json.NewDecoder(r.Body).Decode(&student)

		// For now, just echo back the received student data as JSON
		json.NewEncoder(w).Encode(student)
		w.Header().Set("Content-Type", "application/json")

		// In a real application, you would typically save the student to a database here
		// w.Write([]byte(student.Name + " has been created successfully"))
	}
}
