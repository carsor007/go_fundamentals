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
// The only type of loop in go is the for loop
// The for loop also doubles as a while loop if you delete the variable declaration from the beginning and the increment  // at the end
// You can also use a loop to iterate over each element in an array or a slice by using range
// We get the index and value 

func main () {
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)

	}
} 
