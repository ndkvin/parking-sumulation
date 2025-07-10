package main

import "fmt"

type parkingLot struct {
	slots []*Car
}

func NewParkingLot(size int) *parkingLot {
	return &parkingLot{slots: make([]*Car, size)}
}

func (p *parkingLot) Park(car *Car) {
	for i := range p.slots {
		if p.slots[i] == nil {
			p.slots[i] = car
			fmt.Printf("Allocated slot number: %d\n", i+1)
			return
		}
	}

	fmt.Printf("Sorry, parking lot is full\n")
}

func (p *parkingLot) Leave(plateNumber string, hours int) {
	bill := 10 + max(hours-2, 0) * 10
	for i := range p.slots {
		if p.slots[i] != nil && p.slots[i].plateNumber == plateNumber {
			p.slots[i] = nil
			fmt.Printf("Registration number %s with Slot Number %d is free with Charge $%d\n", plateNumber, i+1, bill)
			return
		}
	}

	fmt.Printf("Registration number %s not found\n", plateNumber)
}

func (p *parkingLot) Status() {
	fmt.Printf("\nSlot No. Registration No.\n")
	for i := range p.slots {
		if p.slots[i] != nil {
			fmt.Printf("%d %s\n", i+1, p.slots[i].plateNumber)
		}
	}
}
