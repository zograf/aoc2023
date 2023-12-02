package main

import "fmt"

func main() {
	//Day 1
	//for _, el := range input {
	//	first, last := FindNumbers(el)
	//	sum += first*10 + last
	// }

	input := ReadFile("2.txt")
	var sum int64 = 0

	for _, el := range input {
		sum += CubeGame(el)
	}

	fmt.Println(sum)
}
