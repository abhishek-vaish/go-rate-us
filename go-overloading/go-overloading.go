package main

import (
	"fmt"
)

//Person Structure
type Person struct {
	Name string
}

//Introduce from Person.Introduce
func (p *Person) Introduce() {
	fmt.Println("Hi, my name is", p.Name)
}

//Saiyan Structure use Person Structure
type Saiyan struct {
	*Person
	power int
}

//Introduce from Saiyan.Introduce   ----Inheritance
func (s *Saiyan) Introduce() {
	fmt.Println("Name : ", s.Name)
	fmt.Println("Power : ", s.power)
}

func main() {
	object := &Saiyan{
		Person: &Person{"Rohan"},
		power:  1000,
	}
	personObject := Person{
		Name: "Abhishek",
	}

	personObject.Introduce()
	object.Introduce()
}
