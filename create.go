package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var customers []Customer

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&customers)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		for _, customer := range customers {
			if customer.ID == uuid.Nil {
				customer.ID = uuid.New()
			}

			if err := db.Create(&customer).Error; err != nil {
				http.Error(w, fmt.Sprintf("Error inserting customer: %v", err), http.StatusInternalServerError)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Customers added successfully!"})

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
