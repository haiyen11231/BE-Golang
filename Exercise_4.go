// 4. Given a file a.txt, write a function that returns a slice of structs representing people, with information extracted from the file. The requirements are:
// The name must be converted to all uppercase.
// The profession must be converted to all lowercase.
// (Refer to file reading instructions at: https://zetcode.com/golang/readfile/).
// 	Tom|Software engineer|1995
// 	John Snow|Teacher|1997
// 	Maria Onitsuka|Actor|1993
// 	Emil|Football player|1987

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func convertToStruct() []Person {
	var persons []Person
	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var person Person
		info := scanner.Text()
		fmt.Println(info)

		infoSlice := strings.Split(info, "|")

		person.name = strings.ToUpper(infoSlice[0]) 
		person.profession = strings.ToLower(infoSlice[1])
		person.yearOfBirth,_ = strconv.Atoi(infoSlice[2])

		persons = append(persons, person)
	}
	return persons
}

func Exercise_4() {
	fmt.Println("Slice of persons: ", convertToStruct())
}