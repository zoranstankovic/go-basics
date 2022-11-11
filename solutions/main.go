package main

import "fmt"

func main() {
	// Assignment 1
	// Declare a variable c to be of type byte
	// then assign a value of 99 to it
	// print decimal value, character representation of it
	// and type using the fmt.Printf function.

	// Output should look like this:
	// The value of c is 99.
	// The character representation is c.
	// The type is uint8.
	var c byte
	c = 99
	// also you can assign in one line
	// var c byte = 99
	fmt.Printf("The value of c is %v.\n", c)
	fmt.Printf("The character representation is %c.\n", c)
	fmt.Printf("The type is %T.\n", c)

	// Assignment 2
	// Visit this website: https://unicode-table.com/en/
	// Choose any letter, symbol or emoji that you like
	// and use it to create new rune variable.
	// Then print the variable character representation and
	// also its Unicode value.
	r := 'か'
	fmt.Printf("The character representation is %c.\n", r)
	fmt.Printf("The Unicode value is %U.\n", r)
}
