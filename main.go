package main

import (
	"fmt"
)

func main() {
	//Day 1
	//for _, el := range input {
	//	first, last := FindNumbers(el)
	//	sum += first*10 + last
	// }

	//Day 2
	//var sum int64 = 0
	//for _, el := range input {
	//	sum += CubeGame(el)
	//}

	//Day 3
	//matrix := make([][]string, len(input))
	//for ind, el := range input {
	//	matrix[ind] = make([]string, 0)
	//	for _, c := range el {
	//		matrix[ind] = append(matrix[ind], string(c))
	//	}
	//}
	//sum := EngineParts(matrix)
	input := ReadFile("4.txt")
	sum := CalculatePoints(input)
	fmt.Println(sum)
}
