package main

import "gofr.dev/pkg/gofr"

func hello() {
	app := gofr.New() // it initialize the gofr object

	app.GET("/greetmessage", func(ctx *gofr.Context) (interface{}, error) {
		return "Hello Gogfr from Nilabh!", nil
	})

	app.Start()
}
