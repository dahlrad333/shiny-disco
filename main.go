package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"shiny-disco/server"
	"time"
)

func main() {
	// >>>>>>>> TOPIC 1.1, 1.2 VEHICLES
	_ = BuildVehicles()

	// >>>>>>>> TOPIC 1.3 HTTP LOGGING WRAPPER
	StartHttpListener()

	// >>>>>>>> TOPIC 2.1 TYPE REFLECTION
	VehicleReflections()

	// >>>>>>>> TOPIC 3.1 CHANNEL & CONTEXT
	RunWorkerpool()
}

// server constructors from outside package
func BuildVehicles() []server.Vehicle {
	car := server.NewCar(250, "Petrol", 10.0, "SUV", "Toyota", "Fortuner", 2023)
	plane := server.NewPlane(1500, "Jet Fuel", 500.0, "Cargo", "Jet", 400)
	motorcycle := server.NewMotorcycle(100, "Petrol", 5.0, "Yamaha", "MT-15", 155)
	boat := server.NewBoat(300, "Diesel", 100.0, "Fishing", 25.5, 10)

	vehicles := []server.Vehicle{car, plane, motorcycle, boat}

	for _, v := range vehicles {
		v.Start()
		v.Details()
		v.Refuel(20)
		// v.Start()
		v.Stop()
		fmt.Println("-----------")
	}

	return vehicles
}

func StartHttpListener() {
	originalHandler := http.HandlerFunc(server.HelloHandler)

	loggedHandler := server.NewLoggingMiddleware(originalHandler)

	http.Handle("/", loggedHandler)

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func VehicleReflections() {
	vehicles := BuildVehicles()

	for _, v := range vehicles {
		server.PrintTypeInfo(v, " ")
	}
}

func RunWorkerpool() {
	inputCh := make(chan server.Job, 10)
	outputCh := make(chan server.Result, 10)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	workerPool := server.NewWorkerPool(ctx, 3, inputCh, outputCh)
	workerPool.Start()

	go func() {
		for i := 1; i <= 20; i++ {
			inputCh <- server.Job{ID: i, Input: i}
		}
		close(inputCh)
	}()

	go func() {
		for result := range outputCh {
			fmt.Printf("Result: Job %d -> Output %d\n", result.JobID, result.Output)
		}
	}()

	go func() {
		for err := range workerPool.Errors() {
			fmt.Println("Error:", err)
		}
	}()

	<-ctx.Done()
	fmt.Println("Context timeout reached")

	workerPool.Shutdown()
}