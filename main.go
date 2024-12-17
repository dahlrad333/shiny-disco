package main

import (
	"fmt"
	"shiny-disco/helper"
)

func main() {
    helper.Help()
}

// call constructors from outside package
func BuildVehicles() {
	car := helper.NewCar(250, "Petrol", 10.0, "SUV", "Toyota", "Fortuner", 2023)
	plane := helper.NewPlane(1500, "Jet Fuel", 500.0, "Cargo", "Jet", 400)
	motorcycle := helper.NewMotorcycle(100, "Petrol", 5.0, "Yamaha", "MT-15", 155)
	boat := helper.NewBoat(300, "Diesel", 100.0, "Fishing", 25.5, 10)

	vehicles := []helper.Vehicle{car, plane, motorcycle, boat}

	for _, v := range vehicles {
		v.Start()
		v.Details()
		v.Refuel(20)
		v.Start()
		v.Stop()
		fmt.Println("-----------")
	}
}