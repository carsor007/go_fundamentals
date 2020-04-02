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
// Arrays are zero index, first element is at position 0
func main () {
	var a [5] int
	a[2] = 7
	fmt.Println(a)
} 
