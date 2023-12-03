package main

import "fmt"

func main() {
	//Day 1
	//for _, el := range input {
	//	first, last := FindNumbers(el)
	//	sum += first*10 + last
	// }

	input := ReadFile("3.txt")

	//Day 2
	//var sum int64 = 0
	//for _, el := range input {
	//	sum += CubeGame(el)
	//}
	matrix := make([][]string, len(input))
	for ind, el := range input {
		matrix[ind] = make([]string, 0)
		for _, c := range el {
			matrix[ind] = append(matrix[ind], string(c))
		}
	}
	sum := EngineParts(matrix)
	fmt.Println(sum)
}
