package api

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// BaseHandler will hold everything that controller needs
type BaseHandler struct {
	db *pgxpool.Pool
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(db *pgxpool.Pool) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}

// GetDataBaseURL returns database url string
func GetDataBaseURL(username string, password string, host string, port string, databaseName string) string {
	return "postgres://" + username + ":" + password + "@" + host + ":" + port + "/" + databaseName
}

// ConnectDB creates connection pool (*pgxpool.Pool) to DB
func ConnectDB(databaseUrl string) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		return dbPool, fmt.Errorf("Unable to connect to database: %v", err)
	}

	return dbPool, err
}

// CloseDB closes database connection (*pgxpool.Pool)
func CloseDB(dbPool *pgxpool.Pool) {
	dbPool.Close()
}

// getCustomer return a customer from DB (select by id)
func getCustomer(dbPool *pgxpool.Pool, customerId int) (customer, error) {

	var c customer = initializeCustomer()
	var err error = nil

	ctx := context.Background()

	var request string = `
	SELECT customers.id, customers.email, customers.name, customers.firstname, customers.birthday, customers.driving_licence_number,
	customer2car.car_id, cars.matriculation_number, cars.brand, cars.model, cars.year
	FROM customers
	LEFT JOIN customer2car
	ON customers.id = customer2car.customer_id
	LEFT JOIN cars
	ON cars.id = customer2car.car_id
	WHERE customers.id = $1
	ORDER BY customers.id, customer2car.car_id;`

	rows, err := dbPool.Query(ctx, request, customerId)
	if err != nil {
		return c, fmt.Errorf("getCustomer - dbPool.Query : %v", err)
	}

	// iterate through the rows
	for rows.Next() {

		var id int32
		var email string
		var name string
		var firstname string
		var birthday time.Time
		var drivingLicenceNumber string

		// Potentially NULL
		var carId *int32
		var matriculationNumber *string
		var brand *string
		var model *string
		var year *int32

		err := rows.Scan(&id, &email, &name, &firstname, &birthday, &drivingLicenceNumber, &carId, &matriculationNumber, &brand, &model, &year)
		if err != nil {
			return c, fmt.Errorf("getCustomer - rows.Scan : %v", err)
		}

		// Write once user information
		if c.ID == 0 {
			// Store DB output
			c.ID = id
			c.Email = email
			c.Name = name
			c.Firstname = firstname
			c.Birthday = birthday
			c.DrivingLicenceNumber = drivingLicenceNumber
		}

		// Append car information
		if carId != nil {
			c.Car = append(c.Car, car{
				ID:                  *carId,
				MatriculationNumber: *matriculationNumber,
				Brand:               *matriculationNumber,
				Model:               *model,
				Year:                *year,
			})
		}
	}

	return c, err
}

// getAllCustomer return all customers from DB
func getAllCustomer(dbPool *pgxpool.Pool) ([]customer, error) {

	var cList []customer = make([]customer, 0)
	var err error = nil

	ctx := context.Background()

	var request string = `
	SELECT customers.id, customers.email, customers.name, customers.firstname, customers.birthday, customers.driving_licence_number,
	customer2car.car_id, cars.matriculation_number, cars.brand, cars.model, cars.year
	FROM customers
	LEFT JOIN customer2car
	ON customers.id = customer2car.customer_id
	LEFT JOIN cars
	ON cars.id = customer2car.car_id
	ORDER BY customers.id, customer2car.car_id;`

	rows, err := dbPool.Query(ctx, request)
	if err != nil {
		return cList, fmt.Errorf("getAllCustomer - dbPool.Query : %v", err)
	}

	// Store customer previous id
	var prevId int32 = 0
	// Store current customer
	var c customer = initializeCustomer()

	// iterate through the rows
	for rows.Next() {

		var id int32
		var email string
		var name string
		var firstname string
		var birthday time.Time
		var drivingLicenceNumber string

		// Potentially NULL
		var carId *int32
		var matriculationNumber *string
		var brand *string
		var model *string
		var year *int32

		err := rows.Scan(&id, &email, &name, &firstname, &birthday, &drivingLicenceNumber, &carId, &matriculationNumber, &brand, &model, &year)
		if err != nil {
			return cList, fmt.Errorf("getAllCustomer - rows.Scan : %v", err)
		}

		// Store each customer
		// New customer detected
		if prevId != id {

			// Store previous customer to list if exist
			if prevId != 0 {
				cList = append(cList, c)
				c = initializeCustomer()
			}

			// Store DB output
			c.ID = id
			c.Email = email
			c.Name = name
			c.Firstname = firstname
			c.Birthday = birthday
			c.DrivingLicenceNumber = drivingLicenceNumber
		}

		// Append car information
		if carId != nil {
			c.Car = append(c.Car, car{
				ID:                  *carId,
				MatriculationNumber: *matriculationNumber,
				Brand:               *matriculationNumber,
				Model:               *model,
				Year:                *year,
			})
		}

		// Update previous customer id
		prevId = id
	}

	return cList, err
}

func TestPostgres(dbPool *pgxpool.Pool) {

	log.Println("Fin connexion ; debut requete")

	rows, err := dbPool.Query(context.Background(), "select * from public.cars")
	if err != nil {
		log.Fatal("error while executing query")
	}

	// iterate through the rows
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}

		// convert DB types to Go types
		id := values[0].(int32)
		matriculation_number := values[1].(string)
		brand := values[2].(string)
		model := values[3].(string)
		year := values[4].(int32)

		log.Println("id", id)
		log.Println("matriculation_number", matriculation_number)
		log.Println("brand", brand)
		log.Println("model", model)
		log.Println("year", year)
	}

}

// updateCustomer return a customer from DB (select by id)
func updateCustomer(dbPool *pgxpool.Pool, customerId int, customerData customer) error {

	var err error = nil

	ctx := context.TODO()

	tx, err := dbPool.Begin(ctx)
	if err != nil {
		return err
	}
	// Rollback is safe to call even if the tx is already closed, so if
	// the tx commits successfully, this is a no-op
	defer tx.Rollback(ctx)

	request := `SELECT COUNT(id) FROM customers
	WHERE id = $1 LIMIT 1;`

	rows, err := tx.Query(ctx, request, customerId)
	if err != nil {
		return fmt.Errorf("updateCustomer - dbPool.Query : %v", err)
	}

	// Count number of customer with customerId
	var count int32 = 0

	// iterate through the rows
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return fmt.Errorf("updateCustomer - rows.Scan : %v", err)
		}
	}

	if count == 0 {
		return fmt.Errorf("updateCustomer - customerId not found : %v", err)
	}

	request = `UPDATE customers
				SET
					email=$1,
					name=$2,
					firstname=$3,
					birthday=$4,
					driving_licence_number=$5
				WHERE id = $6`

	_, err = tx.Exec(ctx, request,
		customerData.Email,
		customerData.Name,
		customerData.Firstname,
		customerData.Birthday,
		customerData.DrivingLicenceNumber,
		customerId)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return err
}

func TestTransaction(dbPool *pgxpool.Pool) error {
	ctx := context.TODO()

	tx, err := dbPool.Begin(ctx)
	if err != nil {
		return err
	}
	// Rollback is safe to call even if the tx is already closed, so if
	// the tx commits successfully, this is a no-op
	defer tx.Rollback(ctx)

	sec := time.Now().Unix()
	secStr := strconv.Itoa(int(sec))
	request := "INSERT INTO \"public\".\"customers\" (\"email\", \"name\", \"firstname\", \"birthday\", \"driving_licence_number\") VALUES ('test" + secStr + "@a-pro.fr', 'Duchene', 'Olivier', '1985-06-20', 'TEST" + secStr + "IJ');"

	log.Println(request)

	_, err = tx.Exec(ctx, request)
	if err != nil {
		return err
	}

	request = "INSERT INTO \"public\".\"cars\" (\"matriculation_number\", \"brand\", \"model\", \"year\") VALUES ('LM" + secStr + "OP', 'Fiat', '500', '2008');"

	log.Println(request)

	_, err = tx.Exec(ctx, request)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return err
}
