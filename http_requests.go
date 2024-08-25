package main

import (
	"fmt"
	"net/http"
)

// testing the http request for each service for user it will send 50 requests to 1st service and 50 requests to the second service
func main() {
	for i := 0; i < 100; i++ {
		resp, err := http.Get("http://localhost:8080/users/create_user")
		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("Request no: %v was successful", i))
		resp.Body.Close()
	}
	for i := 0; i < 15; i++ {
		resp, err := http.Get("http://localhost:8080/products/create_product")
		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("Request no: %v was successful", i))
		resp.Body.Close()
	}
}
