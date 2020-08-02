package main

import "fmt"

type avengers struct {
	name         string
	power        int
	newStructure *avengers
}

func main() {
	object := &avengers{
		name:  "Iron Man",
		power: 2000,
		//fields of a structures
		newStructure: &avengers{
			name:  "Hulk Buster",
			power: 8000,
		},
	}
	captain := avengers{
		name:  "Captain America",
		power: 5000,
	}
	howkie := avengers{
		name: "Howkie",
	}
	spiderMan := newAvengers("Spider Man", 1000)

	updateIronMan(object)
	fmt.Println("Avengers : ", object.newStructure.name)
	fmt.Println("Avengers : ", object.name)
	fmt.Println("Avengers : ", captain.name)
	fmt.Println("Avengers : ", howkie.name)
	fmt.Println("Howkie Powers : ", howkie.power)
	fmt.Println("New Avengers : ", spiderMan)
}

/* [NEW] keyword
goku := new(Saiyan)
// same as
goku := &Saiyan{}
*/

//Structures doesn't have constructures but we create a function that return structures.
func newAvengers(name string, power int) avengers {
	return avengers{
		name:  name,
		power: power,
	}
}

//Pointers
func updateIronMan(o *avengers) {
	o.power = 3000
}
