//program always start with package main

package main

//can import multiple packages

import (
	"fmt"
	"errors"
	"math"
)

//always starts with function main which does not take or return any arguments
//some example data types; int, byte, string, byte, rune, float32, complex64
//If no initial value is provided for a variable, it will default to the zero value for the variable type
// when you give variables values, you can use a short term syntax instead of var x varType = 1
// function with multiple return values 

func main () {
   result, err := sqrt(64)

   if err != nil {
      fmt.Println(err)
   } else {
     fmt.Println(result)
   }
} 


// takes a float64
// multiple return values float64 and error
// go doesnt have exceptions if you're used to working with those

func sqrt (x float64)(float64, error) {
      if x < 0 {
          return 0, errors.New("Undefined for negative numbers") 
      }

      return math.Sqrt(x), nil
}


