// 2. Write a function with a string as input, returning a map where:
// 	The key is a character from the string.
// 	The value is the number of times that character appears in the string.

package main

import "fmt"

func getCharOccurrences (s string) map[string]int {
	m := make(map[string]int)
	for idx := range s {
		m[string(s[idx])]++
	}
	return m
}

func Exercise_2() {
	var s string
	fmt.Println("Enter a string: ")
	fmt.Scan(&s)
	fmt.Println("Character occurences:", getCharOccurrences(s))
}