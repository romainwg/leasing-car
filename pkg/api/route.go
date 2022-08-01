package api

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	log.Println("hello")
    fmt.Fprintf(w, "hello\n")
}
func headers(w http.ResponseWriter, req *http.Request) {
	log.Println("header")
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func InitRoute(lp string) error {

	http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

	err := http.ListenAndServe(":"+lp, nil)

	if err != nil {
		return fmt.Errorf("http.ListenAndServe : %v", err)
	}

	return err
}

