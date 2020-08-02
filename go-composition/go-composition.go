package main

import (
	"fmt"
)

//Composition are like Inheritance by which we can reuse our code in different structures

//Person Structure
type Person struct {
	Name string
}

//Introduce from Person.Introduce
func (p *Person) Introduce() {
	fmt.Println("Hi, my name is ", p.Name)
}

//Saiyan Structure use Person Structure
type Saiyan struct {
	*Person
	power int
}

func main() {
	object := &Saiyan{
		Person: &Person{"Abhishek"},
		power:  1000,
	}
	fmt.Println(object.Person.Name)
	fmt.Println(object.Person)
}
