package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Employee struct {
	ID     int    `json:"id"`
	Name   string `json:"employee_name"`
	Salary int    `json:"employee_salary"`
	Age    int    `json:"employee_age"`
	Image  string `json:"profile_image"`
}

// type Response struct {
// 	Data []Employee `json:"data"`
// }
type Response struct {
	Status  string     `json:"status"`
	Data    []Employee `json:"data"`
	Message string     `json:"message"`
}

func GetEmployees() ([]Employee, error) {
	// // Create an HTTP client with insecure TLS settings
	// transport := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	// client := &http.Client{Transport: transport}
	
	// // Send GET request
	// response, err := client.Get("https://dummy.restapiexample.com/api/v1/employees")
	response, err := http.Get("http://dummy.restapiexample.com/api/v1/employees")
	if err != nil {
		// log.Fatal(err)
		return nil, fmt.Errorf("Failed to make request: %w", err)
	}
	defer response.Body.Close()

	fmt.Println("Response: ", response)

	// Read response body
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		// log.Fatal(err)
		return nil, fmt.Errorf("Failed to read response: %w", err)
	}

	fmt.Printf("Raw Response Body:\n%s\n", string(responseData))
	
	fmt.Println("ResponseBody: ", response.Body)
	fmt.Println("ResponseData: ", responseData)

	// Parse JSON
	var responseObject Response
	if err := json.Unmarshal(responseData, &responseObject); err != nil {
		return nil, fmt.Errorf("Failed to parse response JSON: %w", err)
	}

	// Log response for debugging
	log.Printf("API Response: %+v\n", responseObject)
	
	return responseObject.Data, nil
}
