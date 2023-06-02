package main

import "fmt"

var TestString = "Hello"

func main() {
	fmt.Println(UniqueChars(TestString))
}

func UniqueChars(str string) bool {
	chars := make(map[rune]bool)

	for _, char := range str {
		if chars[char] {
			return false
		}
		chars[char] = true
	}

	return true
}
