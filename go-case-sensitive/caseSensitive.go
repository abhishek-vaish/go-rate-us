package main

import "fmt"

func main() {
	var loggedIn bool = true
	var balance int = 10
	var len, err = fmt.Println("Hello!! My name is Abhishek Vaish")

	if loggedIn {
		fmt.Println("Show web page")
		if balance == 10 || err == nil {
			fmt.Println("You are able to purchase your stuff")
			fmt.Println(len)
		} else {
			fmt.Println("Your balance is below 10")
		}

	} else {
		fmt.Println("Please Logged in!!")
	}
}
