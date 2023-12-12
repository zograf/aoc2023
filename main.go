package main

import "fmt"

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

	//Day 4
	//sum := CalculatePoints(input)
	//fmt.Println(sum)

	// Day 5
	//FindLowestLocation_Pt2(input)
	//FindLowestLocation_Pt2_BruteForce(input)

	// Day 6
	//fmt.Println(BoatRace(input, true))
	//fmt.Println(BoatRace(input, false))

	// Day7
	//fmt.Println(CamelCards(input))

	// Day 8
    //fmt.Println(HauntedWasteland(input))

    // Day 9
    //fmt.Println(Oasis(input))

    // Day 10
    //fmt.Println(PipeMaze_Pt2(input))

    // Day 11
    // fmt.Println(CosmicExpansion(input))

    input := ReadFile("12.txt")
    fmt.Println(HotSprings(input))
}
