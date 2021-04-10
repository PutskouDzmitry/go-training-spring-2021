package main

import "fmt"

func removeChar(word string) string{
	return word[1:len(word) - 1]
}

func main() {
	fmt.Println(removeChar("Dima"))
}
