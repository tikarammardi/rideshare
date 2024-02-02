package service

import (
	"fmt"
	"math"
)

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
