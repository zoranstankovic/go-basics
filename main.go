package main

import "fmt"

func main() {
	// single line comment

	/* This is
	multi line
	comment
	*/

	// zero value
	var zeroValue bool
	fmt.Println("Zero value for boolean is:", zeroValue)

	// how to declare boolean variable using
	// short declaration syntax
	isValid := true
	fmt.Println("Is password correct:", isValid)

	// comparison operators
	// greater than: >
	// less than: <
	// greater than or equal to: >=
	// less than or equal to: <=
	// equal to: ==
	// not equal to: !=
	a := 5
	b := 7
	result := a < b
	fmt.Println(a, "<", b, "is:", result)
	fmt.Println(a, ">", b, "is:", a > b)
	fmt.Println(a, "<=", b, "is:", a <= b)
	fmt.Println(a, ">=", b, "is:", a >= b)
	fmt.Println(a, "==", b, "is:", a == b)
	fmt.Println(a, "!=", b, "is:", a != b)
	// = for assignment, == for comparison

	// Logical Operators
	// NOT - !
	// AND - &&
	// OR - ||
	isUserLoggedIn := false
	role := "editor"
	fmt.Println("User is not logged in:", !isUserLoggedIn)
	canAccess := isUserLoggedIn && role == "admin"
	fmt.Println("User is logged in and can access Admin section:", canAccess)
	canAccess = role == "admin" || role == "editor"
	fmt.Println("User is admin or editor and can access:", canAccess)
	canAccess = isUserLoggedIn && role == "admin" || role == "editor"
	// false && false || true
	// false || true
	// returns true
	fmt.Println("User is logged in and has correct role:", canAccess)

}
