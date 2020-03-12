package main

import "fmt"

var x int = 11

func main() {

	if x <= 1 {

		fmt.Println("Oh...dear")

	} else if x == 2 || x == 3 || x == 5 || x == 7 {

		fmt.Println("Ueaaasss", x, "is a PRIME number")

	} else if x%2 != 0 && x%3 != 0 && x%5 != 0 && x%7 != 0 {

		fmt.Println("YES babe !! Whoohooo", x, "is a PRIME number")

	} else {

		fmt.Println("NO babe", x, "is not a prime number...")
	}
}
