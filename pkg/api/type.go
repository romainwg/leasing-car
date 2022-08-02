package api

import "time"

type customer struct {
	ID                   int32     `json:"id"`
	Email                string    `json:"email"`
	Name                 string    `json:"name"`
	Firstname            string    `json:"firstname"`
	Birthday             time.Time `json:"birthday"`
	DrivingLicenceNumber string    `json:"driving_licence_number"`
	Car                  []car     `json:"car"`
}

type car struct {
	ID                  int32  `json:"id"`
	MatriculationNumber string `json:"matriculation_number"`
	Brand               string `json:"brand"`
	Model               string `json:"model"`
	Year                int32  `json:"year"`
}

type customer2car struct {
	CustomerID int32 `json:"customer_id"`
	CarID      int32 `json:"car_id"`
}

// initializeCustomer return a customer with car array initialize
func initializeCustomer() customer {
	var c customer
	c.Car = make([]car, 0)
	return c
}

func checkCustomer(c customer) bool {
	var b bool = false

	// Email regexp
	// Name regexp
	// Firstname regexp
	// Birthday regexp
	// DrivingLicenceNumber regexp

	b = true
	return b
}
