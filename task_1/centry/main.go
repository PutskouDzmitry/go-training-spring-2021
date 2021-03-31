package main

import "fmt"

func centuryFromYear(year int) int {
	if year <= 0 {
		return -1
	}
	if year % 100 > 0 {
		return year / 100 + 1
	}
	return year / 100

}


func main() {
	fmt.Println(centuryFromYear(58))
}
