package main

type Car struct {
	plateNumber string
}

func NewCar(plateNumber string) *Car {
	return &Car{plateNumber}
}
