package main

import (
	"strconv"
	"strings"
)

func CubeGame(input string) int64 {
	var max map[string]int64 = map[string]int64{"red": 0, "green": 0, "blue": 0}

	for _, el := range strings.Split(strings.Split(input, ": ")[1], "; ") {
		for _, count := range strings.Split(el, ", ") {
			num_col := strings.Split(count, " ")
			number, _ := strconv.Atoi(num_col[0])
			color := num_col[1]

			if max[color] < int64(number) {
				max[color] = int64(number)
			}
		}
	}
	return max["red"] * max["green"] * max["blue"]
}
