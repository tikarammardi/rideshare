package utils

const (
	AddDriver   = "ADD_DRIVER"
	AddRider    = "ADD_RIDER"
	Match       = "MATCH"
	StartRide   = "START_RIDE"
	StopRide    = "STOP_RIDE"
	RideStarted = "RIDE_STARTED"
	RideStopped = "RIDE_STOPPED"
	Bill        = "BILL"

	InvalidRide        = "INVALID_RIDE"
	RideNotCompleted   = "RIDE_NOT_COMPLETED"
	NoDriversAvailable = "NO_DRIVERS_AVAILABLE"
	DriversMatched     = "DRIVERS_MATCHED"
	UnknownCommand     = "Unknown command:"

	BaseFare        = 50.0
	PerKmCharge     = 6.5
	PerMinuteCharge = 2.0
	ServiceTax      = 0.20
)
