package main

import (
	"fmt"
)

func main() {
	var variable float64 = 99.2
	variableRef := &variable

	variable = variable * 2.2

	fmt.Println(variable)
	fmt.Println(*variableRef)
}
