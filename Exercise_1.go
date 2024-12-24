// 1. Write a function to create a struct representing a person (including name, profession, and year of birth). The struct should have:
// 	A method to calculate the age of the person.
// 	A method to check if the person is suitable for their profession (if the year of birth is divisible by the number of characters in their name).

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Person struct {
	name, profession string
	yearOfBirth int
}

func (p *Person) CalculateAge() int {
	return time.Now().Year() - p.yearOfBirth
}

func (p *Person) IsSuitableProfession() bool {
	fmt.Println("Year of Birth: ", p.yearOfBirth)
	fmt.Println("Length of Name: ", len(p.name))
	return p.yearOfBirth % len(p.name) == 0
}

func Exercise_1() {
	reader := bufio.NewReader(os.Stdin)

	var p Person
	fmt.Println("Enter your name: ")
	// fmt.Scan(&p.name)
	p.name, _ = reader.ReadString('\n')
	p.name = p.name[:len(p.name)-1]

	fmt.Println("Enter your profession: ")
	// fmt.Scan(&p.profession)
	p.profession, _ = reader.ReadString('\n')
	p.profession = p.profession[:len(p.profession)-1]

	fmt.Println("Enter your year of birth: ")
	fmt.Scan(&p.yearOfBirth)

	fmt.Println("Your age: ", p.CalculateAge())
	if p.IsSuitableProfession() {
		fmt.Println("Your profession is fit!!!")
	} else {
		fmt.Println("Your profession is not fit!!!")
	}
}

// Linux/Unix/macOS: A newline is typically represented as \n.
// Windows: A newline is often represented as \r\n (carriage return \r followed by a newline \n).
//  => if run in Windows, then the number of characters in 'name' and 'profession' are 1 more than actual number b/c of '\r'