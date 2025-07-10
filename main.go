package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var parkingLot parkingLot

	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		part := strings.Fields(line)

		switch strings.ToLower(part[0]) {
			case "create_parking_lot":
				size, _ := strconv.Atoi(part[1])
				parkingLot = *NewParkingLot(size)
			case "park":
				parkingLot.Park(NewCar(part[1]))
			case "leave":
				hours, _ := strconv.Atoi(part[2])
				parkingLot.Leave(part[1], hours)
			case "status":
				parkingLot.Status()
		}		
	}
}
