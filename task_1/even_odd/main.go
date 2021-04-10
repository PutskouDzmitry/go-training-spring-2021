package main

import "fmt"

func evenOrOdd(number int) string {
	if number < 0 {
		return "-1"
	}
	if number % 2 == 0 {
		return "Even"
	} else{
		return "Odd"
	}
}

func main() {
	fmt.Println(evenOrOdd(1))
}
