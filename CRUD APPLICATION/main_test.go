package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

func setupApp() *gofr.App {
	app := gofr.New()

	// Routes
	app.GET("/cars", ListCars)
	app.POST("/cars", AddCar)
	app.PUT("/cars/:id/repair", UpdateCarRepair)
	app.DELETE("/cars/:id", RemoveCar)

	return app
}

func TestCarCRUD(t *testing.T) {
	// Setup
	app := setupApp()

	// Test AddCar
	newCar := Car{Model: "Convertible", Status: "In Garage"}
	newCarJSON, _ := json.Marshal(newCar)
	reqAddCar := httptest.NewRequest("POST", "/cars", bytes.NewBuffer(newCarJSON))
	reqAddCar.Header.Set("Content-Type", "application/json")
	respAddCar := httptest.NewRecorder()
	app.ServeHTTP(respAddCar, reqAddCar)

	assert.Equal(t, http.StatusOK, respAddCar.Code)

	// Test ListCars
	reqListCars := httptest.NewRequest("GET", "/cars", nil)
	respListCars := httptest.NewRecorder()
	app.ServeHTTP(respListCars, reqListCars)

	assert.Equal(t, http.StatusOK, respListCars.Code)

	var retrievedCars []Car
	err := json.Unmarshal(respListCars.Body.Bytes(), &retrievedCars)
	assert.NoError(t, err)
	assert.Len(t, retrievedCars, 3) // Initial 2 + Newly Added 1

	// Test UpdateCarRepair
	idToUpdate := retrievedCars[2].ID
	reqUpdateCar := httptest.NewRequest("PUT", fmt.Sprintf("/cars/%d/repair", idToUpdate), nil)
	respUpdateCar := httptest.NewRecorder()
	app.ServeHTTP(respUpdateCar, reqUpdateCar)

	assert.Equal(t, http.StatusOK, respUpdateCar.Code)

	// Test RemoveCar
	idToRemove := retrievedCars[2].ID
	reqRemoveCar := httptest.NewRequest("DELETE", fmt.Sprintf("/cars/%d", idToRemove), nil)
	respRemoveCar := httptest.NewRecorder()
	app.ServeHTTP(respRemoveCar, reqRemoveCar)

	assert.Equal(t, http.StatusOK, respRemoveCar.Code)

	// Verify ListCars after removal
	reqListCarsAfterRemove := httptest.NewRequest("GET", "/cars", nil)
	respListCarsAfterRemove := httptest.NewRecorder()
	app.ServeHTTP(respListCarsAfterRemove, reqListCarsAfterRemove)

	assert.Equal(t, http.StatusOK, respListCarsAfterRemove.Code)

	Car := carsAfterRemove []Car
	err = json.Unmarshal(respListCarsAfterRemove.Body.Bytes(), &carsAfterRemove)
	assert.NoError(t, err)
	assert.Len(t, carsAfterRemove, 2) // Initial 2
}
