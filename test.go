package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateCustomer(t *testing.T) {
	// Create a new HTTP request
	customer := Customer{
		Name:  "John Doe",
		City:  "New York",
		Email: "john.doe@example.com",
	}
	body, _ := json.Marshal(customer)
	req, err := http.NewRequest("POST", "/createCustomers", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Create)

	// Call the handler with the request and recorder
	handler.ServeHTTP(rr, req)

	// Assert the status code and response body
	assert.Equal(t, http.StatusOK, rr.Code)
	var createdCustomer Customer
	err = json.NewDecoder(rr.Body).Decode(&createdCustomer)
	assert.Nil(t, err)
	assert.Equal(t, customer.Name, createdCustomer.Name)
	assert.Equal(t, customer.City, createdCustomer.City)
	assert.Equal(t, customer.Email, createdCustomer.Email)
	assert.NotEqual(t, uuid.Nil, createdCustomer.ID) // Check that the ID is generated
}

func TestGetCustomerByID(t *testing.T) {
	// First, create a customer to retrieve
	customer := Customer{
		Name:  "Jane Doe",
		City:  "Los Angeles",
		Email: "jane.doe@example.com",
	}
	db.Create(&customer)

	// Make the GET request to fetch the customer by ID
	req, err := http.NewRequest("GET", "/getCustomerByID?id="+customer.ID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCustomerByID)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, rr.Code)
	var retrievedCustomer Customer
	err = json.NewDecoder(rr.Body).Decode(&retrievedCustomer)
	assert.Nil(t, err)
	assert.Equal(t, customer.ID, retrievedCustomer.ID)
}

func TestDisplayCustomers(t *testing.T) {
	// Create a customer to test
	customer := Customer{
		Name:  "Alice",
		City:  "San Francisco",
		Email: "alice@example.com",
	}
	db.Create(&customer)

	// Make a GET request to display all customers
	req, err := http.NewRequest("GET", "/displayCustomers", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(displayCustomers)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, rr.Code)
	var customers []Customer
	err = json.NewDecoder(rr.Body).Decode(&customers)
	assert.Nil(t, err)
	assert.True(t, len(customers) > 0)
}

func TestUpdateCustomer(t *testing.T) {
	// First, create a customer to update
	customer := Customer{
		Name:  "Bob",
		City:  "Chicago",
		Email: "bob@example.com",
	}
	db.Create(&customer)

	// Update the customer's name and city
	customer.Name = "Robert"
	customer.City = "Houston"
	body, _ := json.Marshal(customer)
	req, err := http.NewRequest("POST", "/updateCustomer", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateCustomer)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, rr.Code)
	var updatedCustomer Customer
	err = json.NewDecoder(rr.Body).Decode(&updatedCustomer)
	assert.Nil(t, err)
	assert.Equal(t, customer.Name, updatedCustomer.Name)
	assert.Equal(t, customer.City, updatedCustomer.City)
}

func TestDeleteCustomer(t *testing.T) {
	// First, create a customer to delete
	customer := Customer{
		Name:  "Eve",
		City:  "Miami",
		Email: "eve@example.com",
	}
	db.Create(&customer)

	// Make a DELETE request to remove the customer
	req, err := http.NewRequest("GET", "/deleteCustomer?id="+customer.ID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteCustomer)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Assert the response
	assert.Equal(t, http.StatusNoContent, rr.Code)
}
