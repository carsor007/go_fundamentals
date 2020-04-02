//program always start with package main

package main

//can import multiple packages

import (
	"fmt"
	
)

// A struct is a collection of fields 
// So you can group things together to create a more logical type
// You create a struct outside of a function, use type followed by the name
// Then the word struct
// You then list the fields you want the struct to have
// The fields need a name and a type, e.g string, int or another struct



type person struct {
	name string
	age int

}


// You can now use the name of the struct 
// To create a struct of that type
// Use curly brackets to set the fields

func main () {
	p := person {name: "John", age: 26}
	fmt.Println(p)
        fmt.Println(p.age)

}

