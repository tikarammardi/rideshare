package service

import (
	"fmt"
	. "github.com/tikarammardi/rideshare/utils"
	"math"
)

func (rs *RideShareService) CalculateBill(distance float64, time int) float64 {

	fare := BaseFare + (distance * PerKmCharge) + float64(time)*PerMinuteCharge
	fare += fare * ServiceTax
	return fare
}

func (rs *RideShareService) Bill(rideID string) {
	ride, exists := rs.Rides[rideID]
	if !exists {
		fmt.Println(InvalidRide)
		return
	}
	if !ride.Completed {
		fmt.Println(RideNotCompleted)
		return
	}
	rider := rs.Riders[ride.RiderID]
	distance := rs.EuclideanDistance(rider.X, rider.Y, ride.Destination.X, ride.Destination.Y)

	// round off to 2 decimal places
	roundDistance := math.Round(distance*100) / 100

	amount := rs.CalculateBill(roundDistance, ride.TimeTaken)
	fmt.Printf("BILL %s %s %.2f\n", rideID, ride.DriverID, amount)
}
