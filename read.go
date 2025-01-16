package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func displayCustomers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var customers []Customer

		if err := db.Find(&customers).Error; err != nil {
			http.Error(w, fmt.Sprintf("Error fetching customers: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
