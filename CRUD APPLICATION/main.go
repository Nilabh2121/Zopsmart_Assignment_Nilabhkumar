package main

import (
	"fmt"

	"gofr.dev/pkg/gofr"
)

type Car struct {
	ID     int    `json:"id"`
	Model  string `json:"model"`
	Status string `json:"status"`
}

var cars = []Car{
	{ID: 1, Model: "Sedan", Status: "In Garage"},
	{ID: 2, Model: "SUV", Status: "In Garage"},
	{ID: 3, Model: "Nano", Status: "In Garage"},
	{ID: 4, Model: "Duster", Status: "In Garage"},
	{ID: 5, Model: "Mercedes", Status: "In Garage"},
	{ID: 6, Model: "Lambhorghini", Status: "In Garage"},
}

func main() {
	app := gofr.New()

	// Routes
	app.GET("/cars", ListCars)
	app.POST("/cars", AddCar)
	app.PUT("/cars/:id/repair", UpdateCarRepair)
	app.DELETE("/cars/:id", RemoveCar)

	// Start the server
	app.Start()
}

func ListCars(ctx *gofr.Context) (interface{}, error) {
	return cars, nil
}

func AddCar(ctx *gofr.Context) (interface{}, error) {
	var newCar Car
	if err := ctx.Bind(&newCar); err != nil {
		return nil, fmt.Errorf("Invalid JSON format: %v", err)
	}

	// Assign a unique ID to the new car
	newCar.ID = len(cars) + 1

	// Add the new car to the list
	cars = append(cars, newCar)
	return newCar, nil
}

func UpdateCarRepair(ctx *gofr.Context) (interface{}, error) {
	id := ctx.Param("id")
	for i := range cars {
		if carID := ctx.Param("id"); carID == id {
			// Update the status to "Under Repair"
			cars[i].Status = "Under Repair"
			return cars[i], nil
		}
	}

	// If the loop completes and no car is found, return a 404
	return nil, fmt.Errorf("Car not found: %s", id)
}

func RemoveCar(ctx *gofr.Context) (interface{}, error) {
	id := ctx.Param("id")
	for i, car := range cars {
		if carID := ctx.Param("id"); carID == id {
			// Remove the car from the list
			cars = append(cars[:i], cars[i+1:]...)
			return car, nil
		}
	}

	// If the loop completes and no car is found, return a 404
	return nil, fmt.Errorf("Car not found: %s", id)
}
