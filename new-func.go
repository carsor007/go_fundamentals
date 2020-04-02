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
// We can create a new function in addition to main 

func main () {
   result := sum(2, 3)
   fmt.Println(result)
} 

// new function that accepts arguments
// the return value type is also declared after the brackets

func sum(x int, y int) int  {
    return x + y
}

