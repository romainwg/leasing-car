package main

import (
	"fmt"
	"log"
	"os"

	"github.com/romainwg/leasing-car/pkg/api"
)

// Initialisation of log file
func init() {
	file, err := os.OpenFile("leasing-car.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	api.InfoLogger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	api.WarningLogger = log.New(file, "WARNING ", log.Ldate|log.Ltime|log.Lshortfile)
	api.ErrorLogger = log.New(file, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
}

// params is a storage containing environment variable in a structure
// to add arguments, change the struct and parseEnvironmentVariables function
type params struct {
	envDBUsername    string
	envDBPassword    string
	envDBHost        string
	envDBPort        string
	envDBName        string
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
			api.ErrorLogger.Fatalf("Environment variable : %s is unavailable", ev)
		}
		envValues = append(envValues, v)
	}

	// Create params structure
	var a params = params{
		envDBUsername:    envValues[0],
		envDBPassword:    envValues[1],
		envDBHost:        envValues[2],
		envDBPort:        envValues[3],
		envDBName:        envValues[4],
		envListeningPort: envValues[5],
	}

	return a
}

func main() {

	var err error = nil

	// Get environment variables
	var as params = parseEnvironmentVariables()

	// Parse databaseURL
	var databaseUrl string = api.GetDataBaseURL(as.envDBUsername, as.envDBPassword, as.envDBHost, as.envDBPort, as.envDBName)

	// Connect to PostgreSQL
	dbPool, err := api.ConnectDB(databaseUrl)
	if err != nil {
		api.ErrorLogger.Fatalln(err)
	}
	defer api.CloseDB(dbPool)
	api.InfoLogger.Println("Connection to database : OK")

	// Create database handler to pass
	h := api.NewBaseHandler(dbPool)

	// Initialize route including environment variable & base handler
	err = h.InitRoute(as.envListeningPort)
	if err != nil {
		api.ErrorLogger.Fatalln(err)
	}

}
