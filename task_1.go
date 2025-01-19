package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var dummyResponse = Response{
	Status: "success",
	Data: []Employee{
		{
			ID:             1,
			Name:   "John Snow",
			Salary: 1000,
			Age:    30,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=John+Snow",
		},
		{
			ID:             2,
			Name:   "Maria Onizuka",
			Salary: 2000,
			Age:    25,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=Maria+Onizuka",
		},
		{
			ID:             3,
			Name:   "Emily",
			Salary: 3000,
			Age:    28,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=Emily",
		},
		{
			ID:             4,
			Name:   "Tom",
			Salary: 4000,
			Age:    35,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=Tom",
		},
		{
			ID:             5,
			Name:   "John Snow",
			Salary: 5000,
			Age:    40,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=John+Snow",
		},
		{
			ID:             6,
			Name:   "Sarah Connor",
			Salary: 6000,
			Age:    32,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=Sarah+Connor",
		},
		{
			ID:             7,
			Name:   "Michael Smith",
			Salary: 7000,
			Age:    45,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=Michael+Smith",
		},
		{
			ID:             8,
			Name:   "Jessica Jones",
			Salary: 8000,
			Age:    29,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=Jessica+Jones",
		},
		{
			ID:             9,
			Name:   "Clark Kent",
			Salary: 9000,
			Age:    33,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=Clark+Kent",
		},
		{
			ID:             10,
			Name:   "Bruce Wayne",
			Salary: 10000,
			Age:    41,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=Bruce+Wayne",
		},
		{
			ID:             11,
			Name:   "Diana Prince",
			Salary: 11000,
			Age:    28,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=Diana+Prince",
		},
		{
			ID:             12,
			Name:   "Peter Parker",
			Salary: 12000,
			Age:    22,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=Peter+Parker",
		},
		{
			ID:             13,
			Name:   "Tony Stark",
			Salary: 13000,
			Age:    45,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=Tony+Stark",
		},
		{
			ID:             14,
			Name:   "Natasha Romanoff",
			Salary: 14000,
			Age:    35,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=Natasha+Romanoff",
		},
		{
			ID:             15,
			Name:   "Steve Rogers",
			Salary: 15000,
			Age:    38,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=Steve+Rogers",
		},
		{
			ID:             16,
			Name:   "Wade Wilson",
			Salary: 16000,
			Age:    34,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=Wade+Wilson",
		},
		{
			ID:             17,
			Name:   "Barry Allen",
			Salary: 17000,
			Age:    27,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=Barry+Allen",
		},
		{
			ID:             18,
			Name:   "Hal Jordan",
			Salary: 18000,
			Age:    39,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=Hal+Jordan",
		},
		{
			ID:             19,
			Name:   "Arthur Curry",
			Salary: 19000,
			Age:    33,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=Arthur+Curry",
		},
		{
			ID:             20,
			Name:   "T'Challa",
			Salary: 20000,
			Age:    37,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=T'Challa",
		},
		{
			ID:             21,
			Name:   "Bucky Barnes",
			Salary: 21000,
			Age:    32,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=Bucky+Barnes",
		},
		{
			ID:             22,
			Name:   "Sam Wilson",
			Salary: 22000,
			Age:    35,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=Sam+Wilson",
		},
		{
			ID:             23,
			Name:   "Wanda Maximoff",
			Salary: 23000,
			Age:    29,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=Wanda+Maximoff",
		},
		{
			ID:             24,
			Name:   "Vision",
			Salary: 24000,
			Age:    36,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=Vision",
		},
		{
			ID:             25,
			Name:   "Scott Lang",
			Salary: 25000,
			Age:    40,
			Image:   "https://dummyimage.com/600x400/000/fff.png&text=Scott+Lang",
		},
	},
}

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
	// Message string     `json:"message"`
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

	// Check status response 
	if response.StatusCode != http.StatusOK {
		fmt.Println("API response is not OK, using dummy data!!!")
		return dummyResponse.Data, nil
	}

	// Parse JSON
	var responseObject Response
	if err := json.Unmarshal(responseData, &responseObject); err != nil {
		return nil, fmt.Errorf("Failed to parse response JSON: %w", err)
	}

	// Log response for debugging
	log.Printf("API Response: %+v\n", responseObject)
	
	return responseObject.Data, nil
}
