package api

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

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

	// Email regexp
	m, err := regexp.MatchString("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$", c.Email)
	if !m || err != nil {
		WarningLogger.Printf("checkCustomer(%v) ; regexp.MatchString (email) ; %v\n", c, err)
		return false
	}

	// Name regexp
	m, err = regexp.MatchString("^[a-zA-ZàáâäãåąčćęèéêëėįìíîïłńòóôöõøùúûüųūÿýżźñçčšžÀÁÂÄÃÅĄĆČĖĘÈÉÊËÌÍÎÏĮŁŃÒÓÔÖÕØÙÚÛÜŲŪŸÝŻŹÑßÇŒÆČŠŽ∂ð ,.'-]+$", c.Name)
	if !m || err != nil {
		WarningLogger.Printf("checkCustomer(%v) ; regexp.MatchString (name) ; %v\n", c, err)
		return false
	}

	// Firstname regexp
	m, err = regexp.MatchString("^[a-zA-ZàáâäãåąčćęèéêëėįìíîïłńòóôöõøùúûüųūÿýżźñçčšžÀÁÂÄÃÅĄĆČĖĘÈÉÊËÌÍÎÏĮŁŃÒÓÔÖÕØÙÚÛÜŲŪŸÝŻŹÑßÇŒÆČŠŽ∂ð ,.'-]+$", c.Firstname)
	if !m || err != nil {
		WarningLogger.Printf("checkCustomer(%v) ; regexp.MatchString (firstname) ; %v\n", c, err)
		return false
	}

	// Birthday regexp
	// Compile regexp with group
	regexpBirthday, err := regexp.Compile("^([0-9]{4})-([0-9]{2})-([0-9]{2})T([0-9]{2}):([0-9]{2}):([0-9]{2})Z$")
	if err != nil {
		WarningLogger.Printf("checkCustomer(%v) ; regexp.Compile ; %v\n", c, err)
		return false
	}

	// Simple matching check
	var birthdayString string = c.Birthday.Format(time.RFC3339)
	m = regexpBirthday.MatchString(birthdayString)
	if !m {
		WarningLogger.Printf("checkCustomer(%v) ; regexpBirthday.MatchString(%s)\n", c, birthdayString)
		return false
	}

	// Check all numbers
	listDateNumberStr := regexpBirthday.FindStringSubmatch(birthdayString)
	// Optional with previous matching
	if len(listDateNumberStr) != 7 {
		WarningLogger.Printf("checkCustomer(%v) ; len(listDateNumberStr) (%d) != 7\n", c, len(listDateNumberStr))
		return false
	}

	var listDateNumberInt []int = make([]int, 0, 6)
	for i := 1; i < 7; i++ {
		// By pass err with previous matching
		n, _ := strconv.Atoi(listDateNumberStr[i])
		listDateNumberInt = append(listDateNumberInt, n)
		fmt.Println(n)
	}

	// Simple check year - Date inf. fix ; Date sup. fix
	if !(listDateNumberInt[0] > 1900 && listDateNumberInt[0] < 2100) {
		WarningLogger.Printf("checkCustomer(%v) ; Simple check year (%d)\n", c, listDateNumberInt[0])
		return false
	}

	// Simple check month - Date inf. fix ; Date sup. fix
	if !(listDateNumberInt[1] >= 1 && listDateNumberInt[1] <= 12) {
		WarningLogger.Printf("checkCustomer(%v) ; Simple check month (%d)\n", c, listDateNumberInt[1])
		return false
	}

	// Simple check day - Date inf. fix ; Date sup. fix
	if !(listDateNumberInt[2] >= 1 && listDateNumberInt[2] <= 31) {
		WarningLogger.Printf("checkCustomer(%v) ; Simple check day (%d)\n", c, listDateNumberInt[2])
		return false
	}

	// Get actual date
	dateNow := time.Now()
	dateIn := time.Date(listDateNumberInt[0], time.Month(listDateNumberInt[1]), listDateNumberInt[2], 0, 0, 0, 0, time.UTC)

	diffDurationHours := int(dateNow.Sub(dateIn).Hours())

	// Less than 18 years
	if diffDurationHours < 24*356*18 {
		WarningLogger.Printf("checkCustomer(%v) ; diffDurationHours (%d) < 24*356*18\n", c, diffDurationHours)
		return false
	}

	// DrivingLicenceNumber regexp
	/* m, err = regexp.MatchString("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$", c.DrivingLicenceNumber)
	if !m || err != nil {
		return false
	} */
	if len(c.DrivingLicenceNumber) == 0 {
		WarningLogger.Printf("checkCustomer(%v) ; len(c.DrivingLicenceNumber) == 0\n", c)
		return false
	}

	return true
}

func checkCustomer2Car(c2C customer2car) bool {
	var b bool = false

	// id : primary key ; > 0
	if c2C.CustomerID > 0 && c2C.CarID > 0 {
		b = true
	}

	return b
}
