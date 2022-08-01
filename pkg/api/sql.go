package api

import "fmt"

func GetDataBaseURL(username string, password string, host string, port string, databaseName string) string {
	var url string

	url = "postgres://"+username+":"+password+"@"+host+":"+port+"/"+databaseName

	return url
}

func TestAPI() {
	fmt.Println("TestAPI")
}