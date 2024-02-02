package main

import (
	"bufio"
	"fmt"
	"github.com/tikarammardi/rideshare/service"
	. "github.com/tikarammardi/rideshare/utils"
	"os"
	"strconv"
	"strings"
)

func processLine(service *service.RideShareService, line string) {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return
	}

	switch parts[0] {
	case AddDriver:
		x, _ := strconv.ParseFloat(parts[2], 64)
		y, _ := strconv.ParseFloat(parts[3], 64)
		service.AddDriver(parts[1], x, y)
	case AddRider:
		x, _ := strconv.ParseFloat(parts[2], 64)
		y, _ := strconv.ParseFloat(parts[3], 64)
		service.AddRider(parts[1], x, y)
	case Match:
		service.Match(parts[1])
	case StartRide:
		n, _ := strconv.Atoi(parts[2])
		service.StartRide(parts[1], parts[3], n)
	case StopRide:
		x, _ := strconv.ParseFloat(parts[2], 64)
		y, _ := strconv.ParseFloat(parts[3], 64)
		time, _ := strconv.Atoi(parts[4])
		service.StopRide(parts[1], x, y, time)
	case Bill:
		service.Bill(parts[1])
	default:
		fmt.Println(UnknownCommand, parts[0])
	}
}

func main() {
	rideShareService := service.NewRideShareService()
	if len(os.Args) < 2 {
		fmt.Println("Please provide the file path as an argument.")
		os.Exit(1)
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error reading the file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		processLine(rideShareService, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning the file: %s\n", err)
	}
}
