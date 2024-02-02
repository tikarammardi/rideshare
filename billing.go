package main

func (rs *RideShareService) CalculateBill(distance float64, time int) float64 {
	baseFare := 50.0
	perKmCharge := 6.5
	perMinuteCharge := 2.0
	serviceTax := 0.20
	fare := baseFare + (distance * perKmCharge) + float64(time)*perMinuteCharge
	fare += fare * serviceTax
	return fare
}
