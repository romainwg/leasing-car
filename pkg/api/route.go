package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (h *BaseHandler) hello(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	log.Println("hello")
	fmt.Fprintf(w, "hello\n")
}

func (h *BaseHandler) headers(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	log.Println("header")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func (h *BaseHandler) helloCustomer(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Println("hello")
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("customerId"))
	fmt.Fprintf(w, "hello\n")

	// Test DB
	TestPostgres(h.db)
}

func (h *BaseHandler) InitRoute(lp string) error {

	// Create router and routess
	router := httprouter.New()

	router.GET("/customer/get/:customerId", h.getCustomer)
	router.GET("/customer/getall", h.getAllCustomers)
	router.POST("/customer/create", h.addCustomer)
	router.PUT("/customer/update/:customerId", h.updateCustomer)
	router.POST("/customer-car/associate", h.associateCustomer2Car)
	router.POST("/customer-car/disassociate", h.disassociateCustomer2Car)

	log.Println("Create routes : OK\nTrying to launch server...")

	// Start server
	err := http.ListenAndServe(":"+lp, router)
	if err != nil {
		return fmt.Errorf("http.ListenAndServe : %v", err)
	}

	return err
}

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
		WarningLogger.Printf("PUT customer/update/%d - 405: successful ; json.NewDecoder(req.Body).Decode(%v) : %v\n", customerId, c, err)
		return
	}

	if !checkCustomer(c) {
		// 405
		w.WriteHeader(http.StatusMethodNotAllowed)
		WarningLogger.Printf("PUT customer/update/%d - 405: successful ; checkCustomer(%v) : %v\n", customerId, c, err)
		return
	}

	err = updateCustomer(h.db, customerId, c)
	if err != nil {
		// 500
		w.WriteHeader(http.StatusInternalServerError)
		WarningLogger.Printf("PUT customer/update/%d  - 500: successful ; updateCustomer(h.db, %d, %v) : %v\n", customerId, customerId, c, err)
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
	InfoLogger.Printf("PUT customer/update/%d - 200: successful", customerId)
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
	log.Println("addCustomer")

	fmt.Fprintf(w, "addCustomer!\n")

	// Test DB
	// TestPostgres(h.db)

	err := TestTransaction(h.db)
	log.Println("TestTransaction", err)
}

func (h *BaseHandler) associateCustomer2Car(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	log.Println("associateCustomer2Car")

	fmt.Fprintf(w, "associateCustomer2Car!\n")

	// Test DB
	// TestPostgres(h.db)
}

func (h *BaseHandler) disassociateCustomer2Car(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	log.Println("disassociateCustomer2Car")

	fmt.Fprintf(w, "disassociateCustomer2Car!\n")

	// Test DB
	// TestPostgres(h.db)
}
