package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		var input struct {
			ID uuid.UUID `json:"id"`
		}

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&input)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		result := db.Delete(&Customer{}, "id = ?", input.ID)
		if result.Error != nil {
			http.Error(w, fmt.Sprintf("Error deleting customer: %v", result.Error), http.StatusInternalServerError)
			return
		}

		if result.RowsAffected == 0 {
			http.Error(w, "Customer not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Customer deleted successfully!"})

	} else {
		// Handle invalid request method
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
