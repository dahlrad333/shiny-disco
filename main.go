package main

import (
	"fmt"
	"shiny-disco/server"
)

func main() {
	// >>>>>>>> TOPIC 1.1, 1.2 VEHICLES
	// BuildVehicles()

	// >>>>>>>> TOPIC 1.3 HTTP LOGGING WRAPPER
	// originalHandler := http.HandlerFunc(server.HelloHandler)

	// loggedHandler := server.NewLoggingMiddleware(originalHandler)

	// http.Handle("/", loggedHandler)

	// log.Println("Starting server on :8080...")
	// log.Fatal(http.ListenAndServe(":8080", nil))

	// >>>>>>>> TOPIC 2 TYPE REFLECTION
	VehicleReflections()

}

// server constructors from outside package
func BuildVehicles() {
	car := server.NewCar(250, "Petrol", 10.0, "SUV", "Toyota", "Fortuner", 2023)
	plane := server.NewPlane(1500, "Jet Fuel", 500.0, "Cargo", "Jet", 400)
	motorcycle := server.NewMotorcycle(100, "Petrol", 5.0, "Yamaha", "MT-15", 155)
	boat := server.NewBoat(300, "Diesel", 100.0, "Fishing", 25.5, 10)

	vehicles := []server.Vehicle{car, plane, motorcycle, boat}

	for _, v := range vehicles {
		// v.Start()
		v.Details()
		v.Refuel(20)
		// v.Start()
		// v.Stop()
		fmt.Println("-----------")
	}
}

func VehicleReflections() {
	car := server.NewCar(250, "Petrol", 10.0, "SUV", "Toyota", "Fortuner", 2023)
	plane := server.NewPlane(1500, "Jet Fuel", 500.0, "Cargo", "Jet", 400)
	motorcycle := server.NewMotorcycle(100, "Petrol", 5.0, "Yamaha", "MT-15", 155)
	boat := server.NewBoat(300, "Diesel", 100.0, "Fishing", 25.5, 10)

	vehicles := []server.Vehicle{car, plane, motorcycle, boat}

	for _, v := range vehicles {
		server.PrintTypeInfo(v, " ")
	}
}