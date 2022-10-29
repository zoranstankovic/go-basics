package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x = 5
	fmt.Println("Value of x is:", x)
	fmt.Println("Default type of x is:", reflect.TypeOf(x))
	fmt.Println("Zero value for", reflect.TypeOf(x), "is:", reflect.Zero(reflect.TypeOf(x)))

	// var y = 7.85
	// fmt.Println("Value of y is", y)
	// fmt.Println("Default type of y is:", reflect.TypeOf(y))
	// fmt.Println("Zero value for", reflect.TypeOf(y), "is:", reflect.Zero(reflect.TypeOf(y)))

	// var z1 = complex(2, 5)
	// var z2 = 3 + 2i
	// fmt.Println("Value of z1 is:", z1)
	// fmt.Println("Value of z2 is:", z2)
	// fmt.Println("Default type of z1 is:", reflect.TypeOf(z1))
	// fmt.Println("Zero value for", reflect.TypeOf(z1), "is:", reflect.Zero(reflect.TypeOf(z1)))
	// var realPart = real(z1)
	// var imaginary = imag(z1)
	// fmt.Println("Real part is:", realPart)
	// fmt.Println("Imaginary part is", imaginary)

	// var a int8 = 200 // max value for int8 is 127
	// fmt.Println("The value of a is:", a)

	// var a int8 = 127
	// fmt.Println("The value of a is:", a)
	// a = a + 1
	// fmt.Println("New value of a is:", a)

	// Arithmetic operations
	// fmt.Println("2+3i + 3+2i =", (2+3i)+(3+2i)) // addition using complex numbers
	// fmt.Println("7.5 - 2.3 =", 7.5-2.3)
	// fmt.Println("2 * 5 =", 2*5)
	// fmt.Println("5 / 2 =", 5/2)
	// fmt.Println("12 % 5 =", 12%5) // only for integers

	// Comparison operators
	// var i1 = 20
	// var i2 = 25
	// fmt.Println("Is i1 greater than i2:", i1 > i2)
	// fmt.Println("Is i1 less than i2:", i1 < i2)
	// fmt.Println("Is i1 greater than or equal to i2:", i1 >= i2)
	// fmt.Println("Is i1 less than or equal to i2:", i1 <= i2)
	// var c1 = 2 + 3i
	// var c2 = 2 + 5i
	// fmt.Println("Is c1 equal to c2:", c1 == c2)
	// fmt.Println("Is c1 not equal to c2:", c1 != c2)

	// var a = 25   // int
	// var b = 4.23 // float64
	// var c = a + b

	// var a = 25 // int
	// var b int64 = 7
	// var c = a + b

	// var a = 25
	// var b = 4.23
	// var c = a + int(b)
	// fmt.Println("Value of c is:", c)

}
