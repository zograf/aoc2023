package main

import (
	"slices"
	"strconv"
)

type EngineNumbers struct {
	value int
	i     int
	j     int
	size  int
}

func EngineParts(matrix [][]string) int {
	var sum int = 0
	nums := FindEngineNumbers(matrix)
	var seen []EngineNumbers = make([]EngineNumbers, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if matrix[i][j] == "*" {
				sum += SumNeighbors(matrix, i, j, nums, seen)
			}
		}
	}
	return sum
}

func FindEngineNumbers(matrix [][]string) []EngineNumbers {
	var ret []EngineNumbers = make([]EngineNumbers, 0)
	for i := 0; i < len(matrix); i++ {
		num := 0
		k := 1
		start_j := len(matrix) - 1
		for j := len(matrix) - 1; j >= 0; j-- {
			if matrix[i][j] >= "0" && matrix[i][j] <= "9" {
				val, _ := strconv.Atoi(matrix[i][j])
				num += val * k
				k *= 10
			} else {
				if num != 0 {
					ret = append(ret, EngineNumbers{num, i, j + 1, start_j - j - 1})
				}
				start_j = j
				num = 0
				k = 1
			}
		}
		if num != 0 {
			ret = append(ret, EngineNumbers{num, i, 0, start_j})
		}
	}
	return ret
}

func SumNeighbors(matrix [][]string, i int, j int, nums []EngineNumbers, seen []EngineNumbers) int {
	var ret int = 1
	var count int = 0
	directions := [][]int{{0, -1}, {-1, -1}, {0, 1}, {-1, 0}, {1, 1}, {1, 0}, {1, -1}, {-1, 1}}
	for _, pair := range directions {
		new_i := pair[0] + i
		new_j := pair[1] + j
		if new_i >= len(matrix) || new_i < 0 || new_j >= len(matrix) || new_j < 0 {
			continue
		}
		for _, el := range nums {
			if slices.Contains(seen, el) {
				continue
			}
			if el.i == new_i && new_j >= el.j && new_j <= el.j+el.size-1 {
				count++
				ret *= el.value
				seen = append(seen, el)
				break
			}
		}
	}
	if count == 2 {
		return ret
	}
	return 0
}
