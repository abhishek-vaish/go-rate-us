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

func power(a, b int) (int, bool) {
	return a * b, true
}

func main() {
	power, boolean := power(4, 5)
	add := add(56, 78)
	log("abhishek")
	fmt.Println(add)
	fmt.Println(power, boolean)
}
