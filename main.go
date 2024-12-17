package main

import (
	"fmt"
	"log"
	"net/http"
	call "shiny-disco/server"
)

func main() {
	// BuildVehicles()

	originalHandler := http.HandlerFunc(call.HelloHandler)

	loggedHandler := call.NewLoggingMiddleware(originalHandler)

	http.Handle("/", loggedHandler)

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// call constructors from outside package
func BuildVehicles() {
	car := call.NewCar(250, "Petrol", 10.0, "SUV", "Toyota", "Fortuner", 2023)
	plane := call.NewPlane(1500, "Jet Fuel", 500.0, "Cargo", "Jet", 400)
	motorcycle := call.NewMotorcycle(100, "Petrol", 5.0, "Yamaha", "MT-15", 155)
	boat := call.NewBoat(300, "Diesel", 100.0, "Fishing", 25.5, 10)

	vehicles := []call.Vehicle{car, plane, motorcycle, boat}

	for _, v := range vehicles {
		// v.Start()
		v.Details()
		v.Refuel(20)
		// v.Start()
		// v.Stop()
		fmt.Println("-----------")
	}
}