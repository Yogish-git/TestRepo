package main

import "fmt"

func main() {
	fmt.Println("hello Yogish!")
	var a int = 7 //declared a int variable

	i := 113 //shortcut for declaration
	fmt.Println(a, i)

	//arrays in go
	array := [4]string{"apple", "orange", "grapes", "pear"}
	fmt.Println(array)

	//structs
	//declaring struct
	type Employee struct {
		firstName string
		lastName  string
		age       int
	}

	//initialising struct instance
	emp1 := Employee{
		firstName: "Ram",
		lastName:  "Kumar",
		age:       20,
	}

	emp2 := Employee{"Yogish", "M", 22}

	fmt.Println(emp1)
	fmt.Println(emp2)

	//added test comment

}
