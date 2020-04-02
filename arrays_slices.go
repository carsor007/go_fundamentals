//program always start with package main

package main

//can import multiple packages

import (
	"fmt"
)

//always starts with function main which does not take or return any arguments
//some example data types; int, byte, string, byte, rune, float32, complex64
//If no initial value is provided for a variable, it will default to the zero value for the variable type
// when you give variables values, you can use a short term syntax instead of var x varType = 1
// An arrays holds multiple elements of the same type
// The number of elements an arrays holds is fixed has a specific length
// All elements are initialized with their type 0 values if not values initialized.
// To overcome the fixed size in arrays, you can use slices instead
// They are an abstraction on  top of arrays to make them easier to work with
// They dont have a fixed number of elements -- like lists in python
// Now you can use the built in append function to add something new to the end of the slice.
func main () {
	a:= []int {5, 4, 3, 2, 1, 0, 1, 2, 3, 4}
	a = append(a, 13, 25)

	fmt.Println(a)
} 
