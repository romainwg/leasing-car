package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/romainwg/leasing-car/pkg/api"
)

// Global variable for logs
var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

// Initialisation of log file
func init() {
	file, err := os.OpenFile("leasing-car.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
}

// params is a storage containing environment variable in a structure
// to add arguments, change the struct and parseEnvironmentVariables function
type params struct{
	envDBUsername string
	envDBPassword string
	envDBHost string
	envDBPort string
	envDBName string
	envListeningPort string
}

// parseEnvironmentVariables
// parse and store environment variables to "params" structure
func parseEnvironmentVariables() params {
	
	/* 
	ENV ENV_LC_DB_USERNAME="postgres"
	ENV ENV_LC_DB_PASSWORD=""
	ENV ENV_LC_DB_HOST="postgres"
	ENV ENV_LC_DB_PORT="5432"
	ENV ENV_LC_DB_NAME="postgres"
	ENV ENV_LC_LISTENING_PORT="6432"
	*/

	// Environment list
	ENV_LIST := [...]string{
		"ENV_LC_DB_USERNAME",
		"ENV_LC_DB_PASSWORD",
		"ENV_LC_DB_HOST",
		"ENV_LC_DB_PORT",
		"ENV_LC_DB_NAME",
		"ENV_LC_LISTENING_PORT",
	}

	// Stockage of environment variable values
	var envValues []string = make([]string, 0)

	// If one env. var. doesn't exist -> exit(1)
	for _, ev := range ENV_LIST {
		v, b := os.LookupEnv(ev)
		if !b {
			fmt.Printf("Environment variable : %s is unavailable", ev)
			ErrorLogger.Fatalf("Environment variable : %s is unavailable", ev)
		}
		envValues = append(envValues, v)
	}

	// Create params structure
	var a params = params{
		envDBUsername: envValues[0],
		envDBPassword: envValues[1],
		envDBHost: envValues[2],
		envDBPort: envValues[3],
		envDBName: envValues[4],
		envListeningPort: envValues[5],
	}

	return a
}

func main() {

	var as params = parseEnvironmentVariables()

	var databaseUrl string = api.GetDataBaseURL(as.envDBUsername, as.envDBPassword, as.envDBHost, as.envDBPort, as.envDBName)

	fmt.Println(databaseUrl)
	fmt.Printf("as : %v\n", as)

	api.TestAPI()

	testPostgres(databaseUrl)

	err := api.InitRoute(as.envListeningPort)
	if err != nil {
		ErrorLogger.Fatalln(err)
	}

}

func testPostgres(databaseUrl string) {
	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// to close DB pool
	defer dbPool.Close()

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
		
		log.Println("id",id)
		log.Println("matriculation_number",matriculation_number)
		log.Println("brand",brand)
		log.Println("model",model)
		log.Println("year",year)
	}
}