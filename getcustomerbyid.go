package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the ID from the request body
	var input struct {
		ID uuid.UUID `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Validate the ID
	if input.ID == uuid.Nil {
		http.Error(w, "Customer ID is required", http.StatusBadRequest)
		return
	}

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Channels to handle results or errors
	resultChan := make(chan *Customer, 1)
	errorChan := make(chan error, 1)

	// Simulate a long-running operation in a goroutine
	go func() {
		// Simulated delay (2 minutes)
		time.Sleep(10 * time.Minute)

		var customer Customer
		result := db.First(&customer, "id = ?", input.ID)
		if result.Error != nil {
			errorChan <- result.Error
		} else {
			resultChan <- &customer
		}
	}()

	// Handle the response or timeout
	select {
	case <-ctx.Done():
		// Context timeout occurred
		http.Error(w, "Request timed out", http.StatusRequestTimeout)
	case err := <-errorChan:
		// Database error occurred
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Customer not found", http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Error fetching customer: %v", err), http.StatusInternalServerError)
		}
	case customer := <-resultChan:
		// Successfully retrieved the customer
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
