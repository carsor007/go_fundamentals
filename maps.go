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
// Maps hold key value pairs - like dictionaries in Python
// The type definition is the word map, in square brackets the type of keys followed by the type of values in the map
// map[string] int
// To create a map, you use the built in make function and give it the vertices type
// We can set key value pairs using square brackets similar to indexing arrays 

func main () {
	vertices :=  make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

        delete(vertices, "square")

	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])
} 
