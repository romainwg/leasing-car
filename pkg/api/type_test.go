package api

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"testing"
)

func init() {

	log.Println("Init")

	// file, err := os.OpenFile("leasing-car.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	InfoLogger = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "WARNING ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
}

func TestCheckCustomerOK(t *testing.T) {

	var expected bool = true
	var ans bool = false

	var inputCustomer string = `{
		"email": "test@contact.com",
		"name": "Nametest'",
		"firstname": "Firstname",
		"birthday": "1950-06-20T00:00:00Z",
		"driving_licence_number": "TESTT654051SM9IJ"
	}`
	// Same condition of route with body.Reader
	inputReader := strings.NewReader(inputCustomer)

	var c customer

	// Try to decode incoming json
	err := json.NewDecoder(inputReader).Decode(&c)
	if err != nil {
		t.Errorf("TestCheckCustomerOK() = %v; want %v ; %v", ans, expected, err)
	}

	ans = checkCustomer(c)

	if ans != expected {
		t.Errorf("TestCheckCustomerOK() = %v; want %v", ans, expected)
	}
}

func TestCheckCustomerKOEmail(t *testing.T) {

	var expected bool = false
	var ans bool = false

	var inputCustomer string = `{
		"email": "t@est@contact.com",
		"name": "Nametest'",
		"firstname": "Firstname",
		"birthday": "1950-06-20T00:00:00Z",
		"driving_licence_number": "TESTT654051SM9IJ"
	}`
	// Same condition of route with body.Reader
	inputReader := strings.NewReader(inputCustomer)

	var c customer

	// Try to decode incoming json
	err := json.NewDecoder(inputReader).Decode(&c)
	if err != nil {
		t.Errorf("TestCheckCustomerOK() = %v; want %v ; %v", ans, expected, err)
	}

	ans = checkCustomer(c)

	if ans != expected {
		t.Errorf("TestCheckCustomerOK() = %v; want %v", ans, expected)
	}
}

func TestCheckCustomerKOName(t *testing.T) {

	var expected bool = false
	var ans bool = false

	var inputCustomer string = `{
		"email": "test@contact.com",
		"name": "Nam@etest'",
		"firstname": "Firstname",
		"birthday": "1950-06-20T00:00:00Z",
		"driving_licence_number": "TESTT654051SM9IJ"
	}`
	// Same condition of route with body.Reader
	inputReader := strings.NewReader(inputCustomer)

	var c customer

	// Try to decode incoming json
	err := json.NewDecoder(inputReader).Decode(&c)
	if err != nil {
		t.Errorf("TestCheckCustomerOK() = %v; want %v ; %v", ans, expected, err)
	}

	ans = checkCustomer(c)

	if ans != expected {
		t.Errorf("TestCheckCustomerOK() = %v; want %v", ans, expected)
	}
}

func TestCheckCustomerKOFirstname(t *testing.T) {

	var expected bool = false
	var ans bool = false

	var inputCustomer string = `{
		"email": "test@contact.com",
		"name": "Nametest'",
		"firstname": "Firstnam_e",
		"birthday": "1950-06-20T00:00:00Z",
		"driving_licence_number": "TESTT654051SM9IJ"
	}`
	// Same condition of route with body.Reader
	inputReader := strings.NewReader(inputCustomer)

	var c customer

	// Try to decode incoming json
	err := json.NewDecoder(inputReader).Decode(&c)
	if err != nil {
		t.Errorf("TestCheckCustomerOK() = %v; want %v ; %v", ans, expected, err)
	}

	ans = checkCustomer(c)

	if ans != expected {
		t.Errorf("TestCheckCustomerOK() = %v; want %v", ans, expected)
	}
}

func TestCheckCustomerKOBirthday(t *testing.T) {

	var expected bool = false
	var ans bool = false

	var inputCustomer string = `{
		"email": "test@contact.com",
		"name": "Nametest'",
		"firstname": "Firstname",
		"birthday": "190-06-20T00:00:00Z",
		"driving_licence_number": "TESTT654051SM9IJ"
	}`
	// Same condition of route with body.Reader
	inputReader := strings.NewReader(inputCustomer)

	var c customer

	// Try to decode incoming json
	err := json.NewDecoder(inputReader).Decode(&c)
	if err == nil {
		t.Errorf("TestCheckCustomerOK() = %v; want %v ; %v", ans, expected, err)
	}

	ans = checkCustomer(c)

	if ans != expected {
		t.Errorf("TestCheckCustomerOK() = %v; want %v", ans, expected)
	}
}

func TestCheckCustomerKODrivingLicenceNumber(t *testing.T) {

	var expected bool = false
	var ans bool = false

	var inputCustomer string = `{
		"email": "test@contact.com",
		"name": "Nametest'",
		"firstname": "Firstname",
		"birthday": "1950-06-20T00:00:00Z",
		"driving_licence_number": ""
	}`
	// Same condition of route with body.Reader
	inputReader := strings.NewReader(inputCustomer)

	var c customer

	// Try to decode incoming json
	err := json.NewDecoder(inputReader).Decode(&c)
	if err != nil {
		t.Errorf("TestCheckCustomerOK() = %v; want %v ; %v", ans, expected, err)
	}

	ans = checkCustomer(c)

	if ans != expected {
		t.Errorf("TestCheckCustomerOK() = %v; want %v", ans, expected)
	}
}
