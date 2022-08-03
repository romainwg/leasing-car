package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// InitRoute creates routes and start HTTP server
func (h *BaseHandler) InitRoute(lp string) error {

	// Create router and routess
	router := httprouter.New()

	router.GET("/customer/get/:customerId", h.getCustomer)
	router.GET("/customer/getall", h.getAllCustomers)
	router.POST("/customer/create", h.addCustomer)
	router.PUT("/customer/update/:customerId", h.updateCustomer)
	router.POST("/customer-car/associate", h.associateCustomer2Car)
	router.POST("/customer-car/disassociate", h.disassociateCustomer2Car)

	router.GET("/", h.home)

	InfoLogger.Println("Create routes : OK\nTrying to launch server on port " + lp + "...")

	// Start server
	err := http.ListenAndServe(":"+lp, router)
	if err != nil {
		return fmt.Errorf("http.ListenAndServe : %v", err)
	}

	return err
}

// parseId tries to convert string id to integer
func parseId(i string) (int, error) {
	o, err := strconv.Atoi(i)
	return o, err
}

func (h *BaseHandler) getCustomer(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	// Add header : Content-Type: application/json
	w.Header().Add("Content-Type", "application/json")

	// Parse customer id
	customerId, err := parseId(ps.ByName("customerId"))
	if err != nil {
		WarningLogger.Printf("parseId(ps.ByName(\"customerId\")) : %v\n", err)
	}

	// Get customer from BDD
	c, err := getCustomer(h.db, customerId)
	if err != nil {
		// 500
		w.WriteHeader(http.StatusInternalServerError)
		WarningLogger.Printf("getCustomer(h.db, %d) : %v\n", customerId, err)
		return
	}

	// Empty customer
	if c.ID == 0 {
		// Output
		w.WriteHeader(http.StatusBadRequest)
		// Log
		InfoLogger.Printf("GET /customer/get/%d - 400: invalid ID supplied", customerId)
		return
	}

	// Parse to JSON
	b, err := json.Marshal(c)
	if err != nil {
		// 500
		w.WriteHeader(http.StatusInternalServerError)
		WarningLogger.Printf("json.Marshal(c) : %v\n", err)
		return
	}

	// Output
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", string(b))

	// Log
	InfoLogger.Printf("GET /customer/get/%d - 200: successful", customerId)
}

func (h *BaseHandler) updateCustomer(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	// Add header : Content-Type: application/json
	w.Header().Add("Content-Type", "application/json")

	// Parse customer id
	customerId, err := parseId(ps.ByName("customerId"))
	if err != nil {
		WarningLogger.Printf("parseId(ps.ByName(\"customerId\")) : %v\n", err)
	}

	var c customer

	// Try to decode incoming json
	err = json.NewDecoder(req.Body).Decode(&c)
	if err != nil {
		// 405
		w.WriteHeader(http.StatusMethodNotAllowed)
		WarningLogger.Printf("PUT /customer/update/%d - 405: error ; json.NewDecoder(req.Body).Decode(%v) : %v\n", customerId, c, err)
		return
	}

	if !checkCustomer(c) {
		// 405
		w.WriteHeader(http.StatusMethodNotAllowed)
		WarningLogger.Printf("PUT /customer/update/%d - 405: error ; checkCustomer(%v) : %v\n", customerId, c, err)
		return
	}

	err = updateCustomer(h.db, customerId, c)
	if err != nil {
		// 500
		w.WriteHeader(http.StatusInternalServerError)
		WarningLogger.Printf("PUT /customer/update/%d  - 500: error ; updateCustomer(h.db, %d, %v) : %v\n", customerId, customerId, c, err)
		return
	}

	// Output
	w.WriteHeader(http.StatusOK)

	// Log
	InfoLogger.Printf("PUT /customer/update/%d - 200: successful", customerId)
}

func (h *BaseHandler) getAllCustomers(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// Add header : Content-Type: application/json
	w.Header().Add("Content-Type", "application/json")

	// Get customer from BDD
	c, err := getAllCustomer(h.db)
	if err != nil {
		// 500
		w.WriteHeader(http.StatusInternalServerError)
		WarningLogger.Printf("getAllCustomer(h.db) : %v\n", err)
		return
	}

	// Parse to JSON
	b, err := json.Marshal(c)
	if err != nil {
		// 500
		w.WriteHeader(http.StatusInternalServerError)
		WarningLogger.Printf("json.Marshal(c) : %v\n", err)
		return
	}

	// Output
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", string(b))

	// Log
	InfoLogger.Printf("GET /customer/getall - 200: successful")
}

func (h *BaseHandler) addCustomer(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// Add header : Content-Type: application/json
	w.Header().Add("Content-Type", "application/json")

	var c customer

	// Try to decode incoming json
	err := json.NewDecoder(req.Body).Decode(&c)
	if err != nil {
		// 405
		w.WriteHeader(http.StatusMethodNotAllowed)
		WarningLogger.Printf("POST /customer/create - 405: error ; json.NewDecoder(req.Body).Decode(%v) : %v\n", c, err)
		return
	}

	if !checkCustomer(c) {
		// 405
		w.WriteHeader(http.StatusMethodNotAllowed)
		WarningLogger.Printf("POST /customer/create - 405: error ; checkCustomer(%v) : %v\n", c, err)
		return
	}

	err = addCustomer(h.db, c)
	if err != nil {
		// 500
		w.WriteHeader(http.StatusInternalServerError)
		WarningLogger.Printf("POST /customer/create  - 500: error ; addCustomer(h.db, %v) : %v\n", c, err)
		return
	}

	// Output
	w.WriteHeader(http.StatusOK)

	// Log
	InfoLogger.Printf("POST /customer/create - 200: successful")
}

func (h *BaseHandler) associateCustomer2Car(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// Add header : Content-Type: application/json
	w.Header().Add("Content-Type", "application/json")

	var c2C customer2car

	// Try to decode incoming json
	err := json.NewDecoder(req.Body).Decode(&c2C)
	if err != nil {
		// 405
		w.WriteHeader(http.StatusMethodNotAllowed)
		WarningLogger.Printf("POST /customer-car/associate - 405 : error ; json.NewDecoder(req.Body).Decode(%v) : %v\n", c2C, err)
		return
	}

	if !checkCustomer2Car(c2C) {
		// 405
		w.WriteHeader(http.StatusMethodNotAllowed)
		WarningLogger.Printf("POST /customer-car/associate - 405 : error ; checkCustomer2Car(%v) : %v\n", c2C, err)
		return
	}

	err = addCustomer2Car(h.db, c2C)
	if err != nil {
		// 500
		w.WriteHeader(http.StatusInternalServerError)
		WarningLogger.Printf("POST /customer-car/associate - 500 : error ; addCustomer2Car(h.db, %v) : %v\n", c2C, err)
		return
	}

	// Output
	w.WriteHeader(http.StatusOK)

	// Log
	InfoLogger.Printf("POST /customer-car/associate - 200 : successful")
}

func (h *BaseHandler) disassociateCustomer2Car(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// Add header : Content-Type: application/json
	w.Header().Add("Content-Type", "application/json")

	var c2C customer2car

	// Try to decode incoming json
	err := json.NewDecoder(req.Body).Decode(&c2C)
	if err != nil {
		// 405
		w.WriteHeader(http.StatusMethodNotAllowed)
		WarningLogger.Printf("POST /customer-car/disassociate - 405 : error ; json.NewDecoder(req.Body).Decode(%v) : %v\n", c2C, err)
		return
	}

	if !checkCustomer2Car(c2C) {
		// 405
		w.WriteHeader(http.StatusMethodNotAllowed)
		WarningLogger.Printf("POST /customer-car/disassociate - 405 : error ; checkCustomer2Car(%v) : %v\n", c2C, err)
		return
	}

	err = removeCustomer2Car(h.db, c2C)
	if err != nil {
		// 500
		w.WriteHeader(http.StatusInternalServerError)
		WarningLogger.Printf("POST /customer-car/disassociate - 500 : error ; removeCustomer2Car(h.db, %v) : %v\n", c2C, err)
		return
	}

	// Output
	w.WriteHeader(http.StatusOK)

	// Log
	InfoLogger.Printf("POST /customer-car/disassociate - 200 : successful")
}

func (h *BaseHandler) home(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// Output
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "https://github.com/romainwg/leasing-car")

	// Log
	InfoLogger.Printf("GET /home - 200: successful")
}
