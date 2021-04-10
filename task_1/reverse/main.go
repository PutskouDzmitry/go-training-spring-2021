package main

func reverse(word string) string {
	if len(word) == 0 {
		return "-1"
	}
	array := []rune(word)
	for i, lastIndex := 0, len(array) - 1; i < lastIndex; i, lastIndex = i + 1, lastIndex - 1{
		array[i], array[lastIndex] = array[lastIndex], array[i]
	}
	return string(array)
}

func main() {
	fmt.Println(reverse("Dima"))
}
