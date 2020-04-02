//program always start with package main

package main

//can import multiple packages

import (
	"fmt"
	
)

//ampersand represents the address of
//pointer is represented using an asterisk
//It's also used to dereference a pointer 
//This gives us access to the value the pointer points to

func main () {
    i := 7
    inc(&i)
    fmt.Println(i)
}

func inc (x *int) {
   *x++

}

