package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var userName string
	var userRating string

	fmt.Println("Enter your name : ")
	reader := bufio.NewReader(os.Stdin)
	userName, _ = reader.ReadString('\n')
	name := strings.TrimSpace(userName)

	fmt.Println("Please enter your rating : ")
	reader = bufio.NewReader(os.Stdin)
	userRating, _ = reader.ReadString('\n')
	ratingConvert, _ := strconv.ParseInt(strings.TrimSpace(userRating), 0, 64)

	if ratingConvert > 4 {
		fmt.Printf("Hey!! %v we are pleasure that you have rated us %d star", name, ratingConvert)
	} else if ratingConvert > 2 || ratingConvert < 4 {
		fmt.Printf("Hey!! %v we always try make our customer happy and satsfy with our services btw thank you for your %d star rating. We try to improve it.", name, ratingConvert)
	}

}
