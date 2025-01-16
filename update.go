package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		var customer Customer

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&customer)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		if customer.ID == uuid.Nil {
			http.Error(w, "Customer ID is required", http.StatusBadRequest)
			return
		}

		var existingCustomer Customer
		result := db.First(&existingCustomer, "id = ?", customer.ID)
		if result.Error != nil {
			http.Error(w, "Customer not found", http.StatusNotFound)
			return
		}

		existingCustomer.Name = customer.Name
		existingCustomer.City = customer.City
		existingCustomer.Email = customer.Email

		if err := db.Save(&existingCustomer).Error; err != nil {
			http.Error(w, fmt.Sprintf("Error updating customer: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Customer updated successfully!"})

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
