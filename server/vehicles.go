package server

import (
	"fmt"
	"sync"
)

var (
	carMutex        sync.Mutex
	planeMutex      sync.Mutex
	motorcycleMutex sync.Mutex
	boatMutex       sync.Mutex
)

type Vehicle interface {
	Start()
	Stop()
	Details()
	Refuel(amount float64)
}

type Engine struct {
    HorsePower int
    FuelType   string
	FuelAmount float64
}

func (e *Engine) Start() {
	if e.FuelAmount <= 0 {
		fmt.Println("Cannot start: Fuel tank is empty.")
		return
	}
	fmt.Printf("Engine with %d HP (Fuel: %s) starting... Fuel Level: %.2f\n", e.HorsePower, e.FuelType, e.FuelAmount)
}

func (e *Engine) Stop() {
	fmt.Println("Engine stopping...")
}

func (e *Engine) Refuel(amount float64) {
	if amount <= 0 {
		fmt.Println("Refuel amount must be greater than zero.")
		return
	}
	e.FuelAmount += amount
	fmt.Printf("Refueled %.2f units. Current fuel: %.2f\n", amount, e.FuelAmount)
}

// vehicle structs are unexported to enforce constructor usage.
type plane struct {
    Engine
    Purpose        string
    Classification string
    Capacity       int
}

func (p plane) Details() {
	fmt.Printf("Plane [%s | %s] with capacity %d passengers.\n", p.Purpose, p.Classification, p.Capacity)
}

func NewPlane(hp int, fuelType string, fuelAmount float64, purpose, classification string, capacity int) Vehicle {
	if hp <= 100 || capacity <= 0 {
		panic("Invalid horsepower or capacity")
	}
	return &plane{
		Engine:        Engine{HorsePower: hp, FuelType: fuelType, FuelAmount: fuelAmount},
		Purpose:       purpose,
		Classification: classification,
		Capacity:      capacity,
	}
}

type car struct {
    Engine
    BodyType string
    Make     string
    Model    string
    Year     int
}

func (c car) Details() {
	fmt.Printf("Car [%s %s %d | BodyType: %s | Fuel Level: %.2f]\n", c.Make, c.Model, c.Year, c.BodyType, c.FuelAmount)
}

func NewCar(hp int, fuelType string, fuelAmount float64, bodyType, make, model string, year int) Vehicle {
	if hp <= 0 || year < 1886 {
		panic("Invalid horsepower or year")
	}
	return &car{
		Engine:   Engine{HorsePower: hp, FuelType: fuelType, FuelAmount: fuelAmount},
		BodyType: bodyType,
		Make:     make,
		Model:    model,
		Year:     year,
	}
}

type motorcycle struct {
    Engine
    Make  string
    Model string
    CC    int
}

func (m motorcycle) Details() {
	fmt.Printf("Motorcycle [%s %s | %d CC | Fuel Level: %.2f]\n", m.Make, m.Model, m.CC, m.FuelAmount)
}

func NewMotorcycle(hp int, fuelType string, fuelAmount float64, make, model string, cc int) Vehicle {
	if hp <= 0 || cc <= 0 {
		panic("Invalid horsepower or CC")
	}
	return &motorcycle{
		Engine: Engine{HorsePower: hp, FuelType: fuelType, FuelAmount: fuelAmount},
		Make:   make,
		Model:  model,
		CC:     cc,
	}
}

type boat struct {
    Engine
    Purpose string
    Length  float64
	Capacity int
}

func (b boat) Details() {
	fmt.Printf("Boat [%s | Length: %.2f meters | Fuel Level: %.2f] with capacity %d passengers.\n", b.Purpose, b.Length, b.FuelAmount, b.Capacity)
}

func NewBoat(hp int, fuelType string, fuelAmount float64, purpose string, length float64, capacity int) Vehicle {
	if hp <= 0 || length <= 0 || capacity <= 0 {
		panic("Invalid horsepower or length")
	}
	return &boat{
		Engine: Engine{HorsePower: hp, FuelType: fuelType, FuelAmount: fuelAmount},
		Purpose: purpose,
		Length:  length,
		Capacity:  capacity,
	}
}
