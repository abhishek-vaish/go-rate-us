package main

import (
	"fmt"
)

func log(message string) {
	fmt.Println("this is message : ", message)
}

func add(a, b int) int {
	return a + b
}

// Go Function returns multiple values
func power(a, b int) (int, bool) {
	results := a * b
	if a == 5 {
		return results, true
	}
	return results, false
}

func main() {
	powerVar, boolean := power(4, 5)
	//we declare _ if we doesnot care about the second return value from the function
	powers, booleanVar := power(5, 8)
	add := add(56, 78)
	log("abhishek")
	fmt.Println(add)
	fmt.Println(powerVar, boolean)
	fmt.Println(powers, booleanVar)
}
