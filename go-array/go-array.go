package main

import "fmt"

func main() {
	var number [4]string
	number[0] = "uno"
	number[1] = "dos"
	number[2] = "tres"

	fmt.Println(number)

	var color = [4]string{"rojo", "gris", "azul", "verde"}
	fmt.Println(color)
	fmt.Println(len(color))
}
