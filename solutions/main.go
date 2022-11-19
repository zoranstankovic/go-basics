package main

import "fmt"

func main() {
	// Assignment 1
	// Declare two variables, firstName and lastName
	// and initialize them to your first and last names
	// Then define a third variable, fullName, that will be the concatenated string
	// of firstName and lastName. Don't forget to use space between them.
	// And in the end, print a message
	// "I am [your full name], and my favorite emoji is [emoji here]."
	// using the fmt.Printf function.
	// visit this site to find an emoji https://unicode-table.com/en/

	firstName, lastName := "Zoran", "Stankovic"

	fullName := firstName + " " + lastName

	fmt.Printf("I am %s, and my favorite emoji is %c.\n", fullName, 'ᚱ')

	// Assignment 2
	// Using the Printf function and double quotes (string literal)
	// print this message to the console.
	// Use Windows as an argument for Printf function.
	// Output:
	// On "Windows" GOPATH is in "C:\Users\[profile_name]\go".

	fmt.Printf("On %q GOPATH is in \"C:\\Users\\smith\\go\".\n", "Windows")
	// or if you used %s format verb
	fmt.Printf("On \"%s\" GOPATH is in \"C:\\Users\\smith\\go\".\n", "Windows")
}
