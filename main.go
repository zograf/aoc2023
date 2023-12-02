package main

import "fmt"

func main() {
	input := ReadFile("1.txt")

	var sum int64 = 0
	for _, el := range input {
		first, last := FindNumbers(el)
		sum += first*10 + last
	}

	fmt.Println(sum)
}
