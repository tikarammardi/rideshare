package models

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
