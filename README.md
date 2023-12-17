# Car Management API

A simple HTTP (REST) API for managing a car collection using the GoFr framework.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Unit Tests](#unit-tests)
- [Bonus Features](#bonus-features)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This project implements a CRUD API for managing a collection of cars. It is built using the GoFr framework, providing a clean and structured approach to develop HTTP APIs.
First Folder i.e GOLOANG_MINI_PROJECT Includes the self learning of go language and how to run gofr framework with the help of documentation after that with help of gofr i made 
first https REST API by taking help from gofr.dev then i started learning how to connect databse using mysql and docker image firstly i downlaod the docker engine and run the 
instances after that in my project directory i run the command based on the code in connecting database i.e 

[docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3306:3306 -d mysql:8.0.30]

Access test_db database and create table customer with columns id and name

[docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE customers (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL);"]

after this in my docker instances there is database named test_db after that it connected sucessfully with my local database. 

After that i made a new project based on the assignmnet given by Zopsmart i.e CRUD application of CAR Management system using Gofr for which i write the go lang code
with the help pf gofr framework and sucessfully executed the following features..

## Features

- **List Cars:** Retrieve the list of cars.
- **Add Car:** Add a new car to the collection.
- **Update Car Repair Status:** Set a car's status to "Under Repair."
- **Remove Car:** Remove a car from the collection.

## Getting Started

Follow these steps to set up and run the project locally:

1. Clone the repository:

   ```bash
   git clone https://github.com/Nilabh2121/Zopsmart_Assignment_Nilabhkumar.git
   
2. Navigate to the project directory:

    cd Zopsmart_Assignment_Nilabhkumar

3. Run the application:

    go run main.go

   
4. Test the API using your preferred method (e.g., Postman, curl).

API Endpoints



#List Cars 
   Endpoint: GET /cars
   Description: Retrieve the list of cars.
   Response:[{"ID": 1, "Model": "Sedan", "Status": "In Garage"},"ID": 2, "Model": "SUV", "Status": "In Garage"},...]

   
#Add Car 
   Endpoint: POST /cars 
   Description: Add a new car to the collection. 
   Request Body: json 
   {"Model": "Convertible", "Status": "In Garage"} 
   Response:json  
   {"ID": 3, "Model": "Convertible", "Status": "In Garage"} 

   
#Update Car Repair Status
   Endpoint: PUT /cars/:id/repair
   Description: Set a car's status to "Under Repair."
   Response:json
   {"ID": 1, "Model": "Sedan", "Status": "Under Repair"}


   
#Remove Car
   Endpoint: DELETE /cars/:id
   Description: Remove a car from the collection.
   Response:json
   {"ID": 2, "Model": "SUV", "Status": "In Garage"}





   


**********// CODE EXPLAINATION //*************

package main

import (
	"fmt"

	"gofr.dev/pkg/gofr"
)
*The main package is the entry point for the executable.
*It imports the necessary packages, including the fmt package for formatting and printing, and the gofr.dev/pkg/gofr package for the GoFr framework.

type Car struct {
	ID     int    `json:"id"`
	Model  string `json:"model"`
	Status string `json:"status"`
}
*Defines a Car struct representing a car with ID, Model, and Status fields. The json tags specify the field names when serializing to JSON.

var cars = []Car{
	{ID: 1, Model: "Sedan", Status: "In Garage"},
	{ID: 2, Model: "SUV", Status: "In Garage"},
}
*Initializes a slice of Car structs representing the initial list of cars in the garage.
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
*The main function sets up the GoFr application, defines routes, and starts the server.
*Creates a new GoFr application instance using gofr.New().

func ListCars(ctx *gofr.Context) (interface{}, error) {
	return cars, nil
}

*Defines a handler function ListCars for the GET /cars route.
*Returns the list of cars as a response.

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
*Defines a handler function AddCar for the POST /cars route.
*Binds the request body to a new Car struct and adds it to the list of cars.
*Returns the newly added car as a response.

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

*Defines a handler function UpdateCarRepair for the PUT /cars/:id/repair route.
*Updates the repair status of the car with the specified ID.
*Returns the updated car as a response.

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
*Defines a handler function RemoveCar for the DELETE /cars/:id route.
*Removes the car with the specified ID from the list.
*Returns the removed car as a response.








***********// SCREENSHOTS OF RESPONSE AND POSTMAN //**************

This shows the car entry in localhost server on port 8000


![Screenshot 2023-12-17 at 3 49 07 PM](https://github.com/Nilabh2121/Zopsmart_Assignment_Nilabhkumar/assets/74805255/57b62e0a-90ef-44f6-b21d-9fd3395c66bf)


This shows the postman Create API sucessfully running by getting the cars models and id 


![Screenshot 2023-12-17 at 3 57 49 PM](https://github.com/Nilabh2121/Zopsmart_Assignment_Nilabhkumar/assets/74805255/cef03cde-3ac1-401d-b44d-bf437584710a)


This shows that Add cars Api sucessfully running and able to update the car according to the need


![Screenshot 2023-12-17 at 4 01 17 PM](https://github.com/Nilabh2121/Zopsmart_Assignment_Nilabhkumar/assets/74805255/109d1d62-f737-4e1b-b68f-07c3a7fd8092)


This shows that Delete operation in the database


![Screenshot 2023-12-17 at 9 17 52 PM](https://github.com/Nilabh2121/Zopsmart_Assignment_Nilabhkumar/assets/74805255/d6e85699-5c3a-4743-a2d8-aa58e4c582e6)


***********// SEQUENCE DIAGRAM  //**************

  +--------+     +--------------+     +-------------+     +-----------+
  |        |     |              |     |             |     |           |
  |        | --> |   User       | --> | ListCars    | --> | List of   |
  |        |     |              |     |   Handler   |     | Cars      |
  |        |     |              |     |             |     |           |
  +--------+     +--------------+     +-------------+     +-----------+

  +--------+     +--------------+     +-------------+     +-----------+
  |        |     |              |     |             |     |           |
  |        | --> |   User       | --> |  AddCar     | --> | New Car   |
  |        |     |              |     |   Handler   |     |           |
  |        |     |              |     |             |     |           |
  +--------+     +--------------+     +-------------+     +-----------+

  +--------+     +--------------+     +------------------+     +-----------+
  |        |     |              |     |                  |     |           |
  |        | --> |   User       | --> | UpdateCarRepair | --> | Updated   |
  |        |     |              |     |    Handler       |     | Car       |
  |        |     |              |     |                  |     |           |
  +--------+     +--------------+     +------------------+     +-----------+

  +--------+     +--------------+     +------------+     +-----------+
  |        |     |              |     |            |     |           |
  |        | --> |   User       | --> | RemoveCar  | --> | Removed   |
  |        |     |              |     |  Handler   |     | Car       |
  |        |     |              |     |            |     |           |
  +--------+     +--------------+     +------------+     +-----------+


+---------+            +---------------+           +---------------+            +---------+
|   User  |            |               |           |               |            |  Cars   |
+---------+            |   Main Code   |           |   ListCars    |            |   DB    |
    |                 |               |           |   Handler     |            |         |
    | POST /cars       +---------------+           +---------------+            |         |
    | -------------------------->|                    |                        |         |
    |                 |  Car Object (JSON)             |                        |         |
    |                 |<------------------------|     |                        |         |
    |                 |                               |                        |         |
    |                 |                               |                        |         |
    |                 |                               |                        |         |
    |                 |                               |                        |         |
    |                 |                               |                        |         |
    |                 |                               |                        |         |
    |                 |                               |                        |         |
    |                 |                               |                        |         |
    |                 |                               |                        |         |






![Sequence_Diagram](https://github.com/Nilabh2121/Zopsmart_Assignment_Nilabhkumar/assets/74805255/6d9fa288-3a61-427a-b518-90f76a523e0c)




***********// CLASS DIAGRAM  //**************

  +-----------------+     +------------------+
  |  Car            |     |   gofr.Context   |
  +-----------------+     +------------------+
  | ID  : string    |     | ...              |
  | Model : string  |     | ...              |
  | Status : string |     | ...              |
  +-----------------+     +------------------+

   
  
![Class_Diagram](https://github.com/Nilabh2121/Zopsmart_Assignment_Nilabhkumar/assets/74805255/9dcc7c30-4556-44cc-b54d-88801476bac8)


  

