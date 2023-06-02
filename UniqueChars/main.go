package main

import "fmt"

var TestString = "Hello"

func main() {
	fmt.Println(UniqueChars(TestString))
}

func UniqueChars(str string) bool {
	// Create a map to store characters as keys and their existence as values
	chars := make(map[rune]bool)

	// Iterate over each character in the string
	for _, char := range str {
		// Check if the character already exists in the chars map
		if chars[char] {
			// If the character already exists, it's a duplicate, so return false
			return false
		}
		// Mark the character as seen by setting its value to true in the charSet map
		chars[char] = true
	}

	// If we reach this point, it means there are no duplicate characters in the string
	return true
}
