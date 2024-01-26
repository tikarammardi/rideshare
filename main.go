package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Driver struct {
	ID        string
	X         float64
	Y         float64
	Available bool
}

type Rider struct {
	ID string
	X  float64
	Y  float64
}

type Ride struct {
	ID          string
	RiderID     string
	DriverID    string
	Started     bool
	Completed   bool
	Destination struct {
		X float64
		Y float64
	}
	TimeTaken int
}

type RideShareService struct {
	Drivers        map[string]Driver
	Riders         map[string]Rider
	Rides          map[string]Ride
	MatchedDrivers map[string][]string
}

func NewRideShareService() *RideShareService {
	return &RideShareService{
		Drivers:        make(map[string]Driver),
		Riders:         make(map[string]Rider),
		Rides:          make(map[string]Ride),
		MatchedDrivers: make(map[string][]string),
	}
}

func (rs *RideShareService) AddDriver(id string, x, y float64) {
	rs.Drivers[id] = Driver{ID: id, X: x, Y: y, Available: true}
}

func (rs *RideShareService) AddRider(id string, x, y float64) {
	rs.Riders[id] = Rider{ID: id, X: x, Y: y}
}

func (rs *RideShareService) Match(riderID string) {
	rider, exists := rs.Riders[riderID]
	if !exists {
		fmt.Printf("Rider with ID does not exist: %s\n", riderID)
		return
	}

	var availableDrivers []Driver
	for _, driver := range rs.Drivers {
		if driver.Available && rs.EuclideanDistance(rider.X, rider.Y, driver.X, driver.Y) <= 5 {
			availableDrivers = append(availableDrivers, driver)
		}
	}

	sort.Slice(availableDrivers, func(i, j int) bool {
		d1, d2 := availableDrivers[i], availableDrivers[j]
		dist1, dist2 := rs.EuclideanDistance(rider.X, rider.Y, d1.X, d1.Y), rs.EuclideanDistance(rider.X, rider.Y, d2.X, d2.Y)
		if dist1 == dist2 {
			return d1.ID < d2.ID
		}
		return dist1 < dist2
	})

	var matchedDriverIDs []string
	for _, driver := range availableDrivers {
		matchedDriverIDs = append(matchedDriverIDs, driver.ID)
	}

	rs.MatchedDrivers[riderID] = matchedDriverIDs

	if len(availableDrivers) == 0 {
		fmt.Println("NO_DRIVERS_AVAILABLE")
	} else {
		fmt.Println("DRIVERS_MATCHED", strings.Join(matchedDriverIDs, " "))
	}
}

func (rs *RideShareService) StartRide(rideID, riderID string, n int) {
	if _, exists := rs.Rides[rideID]; exists {
		fmt.Println("INVALID_RIDE")
		return
	}

	if n-1 >= len(rs.MatchedDrivers[riderID]) {
		fmt.Println("INVALID_RIDE")
		return
	}

	driverID := rs.MatchedDrivers[riderID][n-1]
	driver, _ := rs.Drivers[driverID]
	if !driver.Available {
		fmt.Println("INVALID_RIDE")
		return
	}

	driver.Available = false
	rs.Drivers[driverID] = driver

	rs.Rides[rideID] = Ride{
		ID:       rideID,
		RiderID:  riderID,
		DriverID: driverID,
		Started:  true,
	}
	fmt.Println("RIDE_STARTED", rideID)
}

func (rs *RideShareService) StopRide(rideID string, x, y float64, time int) {
	ride, exists := rs.Rides[rideID]
	if !exists || !ride.Started || ride.Completed {
		fmt.Println("INVALID_RIDE")
		return
	}

	ride.Completed = true
	ride.Destination.X = x
	ride.Destination.Y = y
	ride.TimeTaken = time
	rs.Rides[rideID] = ride

	fmt.Println("RIDE_STOPPED", rideID)
}

func (rs *RideShareService) CalculateBill(distance float64, time int) float64 {
	baseFare := 50.0
	perKmCharge := 6.5
	perMinuteCharge := 2.0
	serviceTax := 0.20
	fare := baseFare + (distance * perKmCharge) + float64(time)*perMinuteCharge
	fare += fare * serviceTax
	return fare
}

func (rs *RideShareService) Bill(rideID string) {
	ride, exists := rs.Rides[rideID]
	if !exists {
		fmt.Println("INVALID_RIDE")
		return
	}
	if !ride.Completed {
		fmt.Println("RIDE_NOT_COMPLETED")
		return
	}
	rider := rs.Riders[ride.RiderID]
	distance := rs.EuclideanDistance(rider.X, rider.Y, ride.Destination.X, ride.Destination.Y)

	// round off to 2 decimal places
	roundDistance := math.Round(distance*100) / 100

	amount := rs.CalculateBill(roundDistance, ride.TimeTaken)
	fmt.Printf("BILL %s %s %.2f\n", rideID, ride.DriverID, amount)
}

func (rs *RideShareService) EuclideanDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func processLine(service *RideShareService, line string) {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return
	}

	switch parts[0] {
	case "ADD_DRIVER":
		x, _ := strconv.ParseFloat(parts[2], 64)
		y, _ := strconv.ParseFloat(parts[3], 64)
		service.AddDriver(parts[1], x, y)
	case "ADD_RIDER":
		x, _ := strconv.ParseFloat(parts[2], 64)
		y, _ := strconv.ParseFloat(parts[3], 64)
		service.AddRider(parts[1], x, y)
	case "MATCH":
		service.Match(parts[1])
	case "START_RIDE":
		n, _ := strconv.Atoi(parts[2])
		service.StartRide(parts[1], parts[3], n)
	case "STOP_RIDE":
		x, _ := strconv.ParseFloat(parts[2], 64)
		y, _ := strconv.ParseFloat(parts[3], 64)
		time, _ := strconv.Atoi(parts[4])
		service.StopRide(parts[1], x, y, time)
	case "BILL":
		service.Bill(parts[1])
	default:
		fmt.Println("Unknown command:", parts[0])
	}
}

func main() {
	service := NewRideShareService()
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
		processLine(service, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning the file: %s\n", err)
	}
}
