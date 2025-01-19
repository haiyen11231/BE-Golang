package main

import (
	"fmt"
	"log"
)

func main() {
	employees, err := GetEmployees()
	if err != nil {
		log.Fatalf("Error fetching employees: %v", err)
	}

	for _, employee := range employees {
        fmt.Printf("ID: %d, Name: %s, Salary: %d, Age: %d, Image: %s\n",
			employee.ID, employee.Name, employee.Salary, employee.Age, employee.Image)
    }
}