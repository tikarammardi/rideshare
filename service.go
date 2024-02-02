package main

import (
	"fmt"
	"github.com/tikarammardi/rideshare/models"
	"math"
	"sort"
	"strings"
)

type RideShareService struct {
	Drivers        map[string]models.Driver
	Riders         map[string]models.Rider
	Rides          map[string]models.Ride
	MatchedDrivers map[string][]string
}

func NewRideShareService() *RideShareService {
	return &RideShareService{
		Drivers:        make(map[string]models.Driver),
		Riders:         make(map[string]models.Rider),
		Rides:          make(map[string]models.Ride),
		MatchedDrivers: make(map[string][]string),
	}
}

func (rs *RideShareService) AddDriver(id string, x, y float64) {
	rs.Drivers[id] = models.Driver{ID: id, X: x, Y: y, Available: true}
}

func (rs *RideShareService) AddRider(id string, x, y float64) {
	rs.Riders[id] = models.Rider{ID: id, X: x, Y: y}
}

func (rs *RideShareService) Match(riderID string) {
	rider, exists := rs.Riders[riderID]
	if !exists {
		fmt.Printf("Rider with ID does not exist: %s\n", riderID)
		return
	}

	var availableDrivers []models.Driver
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

	rs.Rides[rideID] = models.Ride{
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

func (rs *RideShareService) EuclideanDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
