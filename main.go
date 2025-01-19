package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Task_1.1")
	employees, err := GetEmployees()
	if err != nil {
		log.Fatalf("Error fetching employees: %v", err)
	}

	for _, employee := range employees {
        fmt.Printf("ID: %d, Name: %s, Salary: %d, Age: %d, Image: %s\n",
			employee.ID, employee.Name, employee.Salary, employee.Age, employee.Image)
    }

	fmt.Println("Task_1.2")
	numJobs := len(employees)
	jobs := make(chan Employee, numJobs)
	results := make(chan string, numJobs)

	for w := 1; w <= 5; w++ {
		go Worker(w, jobs, results)
	}

	for _, employee := range employees {
		jobs <- employee
	}
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		fmt.Println(<- results)
	}
}