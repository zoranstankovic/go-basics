package main

import (
	"fmt"
)

func main() {
	var tvShow string

	fmt.Printf("The zero value is %q, and type is %T\n", tvShow, tvShow)

	// shoppingList := "\tBread\n\tApples\n\tEggs"
	// shoppingList2 := `	Bread
	// Apples
	// Eggs
	// `

	// fmt.Println("Shopping List 1:")
	// fmt.Println(shoppingList)
	// fmt.Println("Shopping List 2:")
	// fmt.Println(shoppingList2)

	// fmt.Println("apple < bear =", "apple" < "bear")
	// fmt.Println("appLe < apple =", "appLe" < "apple")
	// fmt.Println("app1e < appLe =", "app1e" < "appLe")
	// fmt.Println("apple < apples =", "apple" < "apples")

	// tvShow = "Good Doctor"
	// myFavoriteTvShow := "Good Doctor"

	// fmt.Println("Is this my favorite tv show:", tvShow == myFavoriteTvShow)

	// fmt.Printf("My favorite tv show is %q.\n", tvShow)
	// fmt.Printf("My favorite tv show is %s.\n", tvShow)

	// fmt.Printf("The value of tvShow is %q, and its type is %T\n", tvShow, tvShow)
	// fmt.Printf("The value of tvShow is %[1]q, and its type is %[1]T\n", tvShow)
	// fmt.Printf("My favorite color is %[1]s, my favorite food is %[2]s, and the sky is %[1]s.\n", "blue", "pizza")

	// fmt.Println("My favorite tv show is " + tvShow)

	// fmt.Println("The first letter of tvShow is:", tvShow[0])
	// fmt.Printf("Type of first letter is %T\n", tvShow[0])

	// tvShow[0] = 'B' // we cannot do this.

	// tvShow = "The Lord of the Rings"

	// fmt.Println("My favorite tv show is", tvShow)

	// song := "What does the fox say?"
	// fmt.Println("the length of the song string is:", len(song))
	// song2 := "What does the 🦊 say?"
	// fmt.Println("the length of the song2 string is:", len(song2))

	// fmt.Println(len("fox"), len("🦊"))

	// fmt.Println("the length of the song string is:", utf8.RuneCountInString(song2))
	// fmt.Println(utf8.RuneCountInString("fox"), utf8.RuneCountInString("🦊"))

	// Assignment 1
	// Declare two variables, firstName and lastName
	// and initialize them to your first and last names
	// Then define a third variable, fullName, that will be the concatenated string
	// of firstName and lastName. Don't forget to use space between them.
	// And in the end, print a message
	// "I am [your full name], and my favorite emoji is [emoji here]."
	// using the fmt.Printf function.
	// visit this site to find an emoji https://unicode-table.com/en/

	// Assignment 2
	// Using the Printf function and double quotes (string literal)
	// print this message to the console.
	// Use Windows as an argument for Printf function.
	// Output:
	// On "Windows" GOPATH is in "C:\Users\[profile_name]\go".

}
